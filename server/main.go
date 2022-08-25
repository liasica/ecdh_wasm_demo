// Copyright (C) liasica. 2022-present.
//
// Created at 2022-07-19
// Based on crypto_wasm by liasica, magicrolan@qq.com.

package main

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "fmt"
)

func main() {
    // fs := http.FileServer(http.Dir("."))
    // http.Handle("/", fs)
    //
    // log.Println("Listening on :3000...")
    // err := http.ListenAndServe(":3000", nil)
    // if err != nil {
    //     log.Fatal(err)
    // }

    fmt.Printf("--ECC Parameters--\n")
    fmt.Printf(" Name: %s\n", elliptic.P256().Params().Name)
    fmt.Printf(" N: %x\n", elliptic.P256().Params().N)
    fmt.Printf(" P: %x\n", elliptic.P256().Params().P)
    fmt.Printf(" Gx: %x\n", elliptic.P256().Params().Gx)
    fmt.Printf(" Gy: %x\n", elliptic.P256().Params().Gy)
    fmt.Printf(" Bitsize: %x\n\n", elliptic.P256().Params().BitSize)

    priva, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    privb, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

    puba := priva.PublicKey
    pubb := privb.PublicKey

    fmt.Printf("\nPrivate key (Alice) %x", priva.D)
    fmt.Printf("\nPrivate key (Bob) %x\n", privb.D)

    fmt.Printf("\nPublic key (Alice) (%x %x)", puba.X, puba.Y)
    fmt.Printf("\nPublic key (Bob) (%x %x)\n", pubb.X, pubb.Y)

    a, _ := puba.Curve.ScalarMult(puba.X, puba.Y, privb.D.Bytes())
    b, _ := pubb.Curve.ScalarMult(pubb.X, pubb.Y, priva.D.Bytes())

    shared1 := sha256.Sum256(a.Bytes())
    shared2 := sha256.Sum256(b.Bytes())

    fmt.Printf("\nShared key (Alice) %x", shared1)
    fmt.Printf("\nShared key (Bob)  %x", shared2)
}
