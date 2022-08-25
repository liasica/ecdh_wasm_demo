// Copyright (C) liasica. 2022-present.
//
// Created at 2022-07-19
// Based on crypto_wasm by liasica, magicrolan@qq.com.

package main

import (
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("."))
    http.Handle("/", fs)

    log.Println("Listening on :3000...")
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatal(err)
    }
}
