//go:build js && wasm

// Copyright (C) liasica. 2022-present.
//
// Created at 2022-07-18
// Based on crypto_wasm by liasica, magicrolan@qq.com.

// 瘦身 https://yryz.net/post/go-wasm/

package main

import (
    "crypto_wasm/crypto"
    "crypto_wasm/dh"
    "fmt"
    "log"
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
    keys := make(map[int][]byte)
    for i := 0; i < num; i++ {
        keys[i] = dh.GenerateKey()
    }
    log.Printf("已生成%d个密钥, 开始测试%d次加密\n", num, num)

    xs := make(map[int][]byte)

    now := time.Now()
    for i := 0; i < num; i++ {
        x, err := crypto.AesEncrypt(b, keys[i])
        if err != nil {
            log.Fatalln(err)
        }
        xs[i] = x
    }
    t := float64(time.Now().Sub(now).Microseconds()) / 1000
    log.Printf("测试%d次加密完成, 耗时%.2fms\n", num, t)

    log.Printf("开始测试%d次解密\n", num)
    now = time.Now()
    for i := 0; i < num; i++ {
        _, err := crypto.AesDecrypt(xs[i], keys[i])
        if err != nil {
            log.Fatalln(err)
        }
    }
    t = float64(time.Now().Sub(now).Microseconds()) / 1000
    log.Printf("测试%d次解密完成, 耗时%.2fms\n", num, t)

    return t
}
