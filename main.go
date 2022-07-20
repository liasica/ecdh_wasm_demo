//go:build js && wasm

// Copyright (C) liasica. 2022-present.
//
// Created at 2022-07-18
// Based on crypto_wasm by liasica, magicrolan@qq.com.

// 瘦身 https://yryz.net/post/go-wasm/
// GOOS=js GOARCH=wasm go build -o main.wasm
// tinygo build -o tiny/crypto.wasm -target wasm ./main.go

package main

import (
    "crypto_wasm/crypto"
    "crypto_wasm/dh"
    "fmt"
    "log"
    "sync"
    "syscall/js"
    "time"
)

var done chan struct{}

func main() {
    done = make(chan struct{})
    fmt.Println("wasm worked!")

    js.Global().Set("runner", js.FuncOf(start))

    <-done
}

func start(_ js.Value, args []js.Value) interface{} {
    // fmt.Println(args)
    if len(args) < 2 {
        log.Fatalln("参数错误")
    }
    b := []byte(args[0].String())
    num := args[1].Int()

    log.Printf("开始生成%d个密钥\n", num)
    var keys sync.Map
    var kwg sync.WaitGroup
    for i := 0; i < num; i++ {
        kwg.Add(1)
        go func(i int) {
            keys.Store(i, dh.GenerateKey())
            kwg.Done()
        }(i)
    }
    kwg.Wait()
    log.Printf("已生成%d个密钥, 开始测试%d次加密\n", num, num)

    xs := make(map[int][]byte)

    now := time.Now()
    var wg sync.WaitGroup
    for i := 0; i < num; i++ {
        wg.Add(1)
        go func(i int) {
            key, _ := keys.Load(i)
            x, err := crypto.AesEncrypt(b, key.([]byte))
            if err != nil {
                log.Fatalln(err)
            }
            xs[i] = x
            wg.Done()
        }(i)
    }
    wg.Wait()
    t := float64(time.Now().Sub(now).Microseconds()) / 1000
    log.Printf("测试%d次加密完成, 耗时%.2fms\n", num, t)

    log.Printf("开始测试%d次解密\n", num)
    now = time.Now()
    var dwg sync.WaitGroup
    for i := 0; i < num; i++ {
        dwg.Add(1)
        go func(i int) {
            key, _ := keys.Load(i)
            _, err := crypto.AesDecrypt(xs[i], key.([]byte))
            if err != nil {
                log.Fatalln(err)
            }
            dwg.Done()
        }(i)
    }

    dwg.Wait()
    t = float64(time.Now().Sub(now).Microseconds()) / 1000
    log.Printf("测试%d次解密完成, 耗时%.2fms\n", num, t)

    return t
}
