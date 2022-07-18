// Copyright (C) liasica. 2022-present.
//
// Created at 2022-06-19
// Based on cryptotest by liasica, magicrolan@qq.com.

package dh

import (
    "bytes"
    "crypto"
    "crypto/rand"
    "github.com/wsddn/go-ecdh"
    "log"
)

func GenerateKey() []byte {
    e := ecdh.NewCurve25519ECDH()

    var privKey1, privKey2 crypto.PrivateKey
    var pubKey1, pubKey2 crypto.PublicKey
    var pubKey1Buf, pubKey2Buf []byte
    var err error
    var ok bool
    var secret1, secret2 []byte

    privKey1, pubKey1, err = e.GenerateKey(rand.Reader)
    if err != nil {
        log.Fatalln(err)
    }
    privKey2, pubKey2, err = e.GenerateKey(rand.Reader)
    if err != nil {
        log.Fatalln(err)
    }

    pubKey1Buf = e.Marshal(pubKey1)
    pubKey2Buf = e.Marshal(pubKey2)

    pubKey1, ok = e.Unmarshal(pubKey1Buf)
    if !ok {
        log.Fatalln("Unmarshal does not work")
    }

    pubKey2, ok = e.Unmarshal(pubKey2Buf)
    if !ok {
        log.Fatalln("Unmarshal does not work")
    }

    secret1, err = e.GenerateSharedSecret(privKey1, pubKey2)
    if err != nil {
        log.Fatalln(err)
    }
    secret2, err = e.GenerateSharedSecret(privKey2, pubKey1)
    if err != nil {
        log.Fatalln(err)
    }

    if !bytes.Equal(secret1, secret2) {
        log.Fatalf("The two shared keys: %d, %d do not match", secret1, secret2)
    }

    return secret1
}
