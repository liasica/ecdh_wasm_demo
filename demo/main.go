// Copyright (C) liasica. 2022-present.
//
// Created at 2022-08-25
// Based on crypto_wasm by liasica, magicrolan@qq.com.

package main

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "crypto/x509"
    "fmt"
)

func generateAliceAndBobKeys() string {
    fmt.Printf("--ECC Parameters--\n")
    fmt.Printf(" Name: %s\n", elliptic.P256().Params().Name)
    fmt.Printf(" N: %x\n", elliptic.P256().Params().N)
    fmt.Printf(" P: %x\n", elliptic.P256().Params().P)
    fmt.Printf(" Gx: %x\n", elliptic.P256().Params().Gx)
    fmt.Printf(" Gy: %x\n", elliptic.P256().Params().Gy)
    fmt.Printf(" Bitsize: %x\n\n", elliptic.P256().Params().BitSize)

    priva, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    // elliptic.MarshalCompressed(elliptic.P256(), priva.X, priva.Y)
    bpriva, _ := x509.MarshalECPrivateKey(priva)
    rpriva, _ := x509.ParseECPrivateKey(bpriva)
    fmt.Printf("%x == %x\n\n", bpriva, rpriva.D)

    privb, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

    puba := priva.PublicKey
    pubb := privb.PublicKey

    fmt.Printf("\nPrivate key (Alice) %x", priva.D)
    fmt.Printf("\nPrivate key (Bob) %x\n", privb.D)

    pubab := elliptic.MarshalCompressed(priva.Curve, priva.PublicKey.X, priva.PublicKey.Y)
    pubbb := elliptic.MarshalCompressed(privb.Curve, privb.PublicKey.X, privb.PublicKey.Y)
    fmt.Printf("\nPublic key (Alice) (%x %x) = %x", puba.X, puba.Y, pubab)
    fmt.Printf("\nPublic key (Bob) (%x %x) = %x\n", pubb.X, pubb.Y, pubbb)

    a, _ := puba.Curve.ScalarMult(puba.X, puba.Y, privb.D.Bytes())
    b, _ := pubb.Curve.ScalarMult(pubb.X, pubb.Y, priva.D.Bytes())

    shared1 := sha256.Sum256(a.Bytes())
    shared2 := sha256.Sum256(b.Bytes())

    fmt.Printf("\nShared key (Alice) %x", shared1)
    fmt.Printf("\nShared key (Bob)  %x", shared2)

    return fmt.Sprintf("%x", pubab)
}

func main() {
    generateAliceAndBobKeys()
}
