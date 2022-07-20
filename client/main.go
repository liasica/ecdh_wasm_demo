// Copyright (C) liasica. 2022-present.
//
// Created at 2022-07-20
// Based on ecdh_wasm_demo by liasica, magicrolan@qq.com.

package main

import (
    "crypto_wasm/crypto"
    "crypto_wasm/dh"
    "log"
    "sync"
    "time"
)

func main() {
    b := []byte("滚滚长江东逝水，浪花淘尽英雄。是非成败转头空，青山依旧在，几度夕阳红。白发渔樵江渚上，惯看秋月春风。一壶浊酒喜相逢，古今多少事， 都付笑谈中。滚滚长江东逝水，浪花淘尽英雄。是非成败转头空，青山依旧在，几度夕阳红。白发渔樵江渚上，惯看秋月春风。一壶浊酒喜相逢，古今多少事， 都付笑谈中。滚滚长江东逝水，浪花淘尽英雄。是非成败转头空，青山依旧在，几度夕阳红。白发渔樵江渚上，惯看秋月春风。一壶浊酒喜相逢，古今多少事， 都付笑谈中。")
    num := 100000

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

    xs := sync.Map{}

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
            xs.Store(i, x)
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
            x, _ := xs.Load(i)
            _, err := crypto.AesDecrypt(x.([]byte), key.([]byte))
            if err != nil {
                log.Fatalln(err)
            }
            dwg.Done()
        }(i)
    }

    dwg.Wait()
    t = float64(time.Now().Sub(now).Microseconds()) / 1000
    log.Printf("测试%d次解密完成, 耗时%.2fms\n", num, t)

}
