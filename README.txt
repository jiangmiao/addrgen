Bitcoin Address Generator

Installation:
  go get github.com/jiangmiao/addrgen

  # or if meets any compile compatible error, try

  go get github.com/golang/dep/cmd/dep
  go get -d github.com/jiangmiao/addrgen
  cd /home/miao/go/src/github.com/jiangmiao/addrgen
  dep ensure
  go install

Features:
  1. Generate bitcoin address in sequence.
  2. Generate bitcoin address in random.
  3. High performance by using multicore.

Usage:
  addrgen [-sequence] [-format "ID PUBKEY"] PREFIX [START_ID=1] [COUNT=1000]
  addrgen -random [-format "PRIVKEY PUBKEY"] [COUNT=1000]

Options:
  -sequence       the private key is SHA256($PREFIX$ID)
  -random         generate in random
  -format         available fields ID PRIVKEY PUBKEY BASE
  -version        show version
  -help           show help

Examples:
  $ addrgen test 10 5
  10 1NoDR9vdm6Rz4mnoVYbFxQ35effYXRDxSZ
  11 1EuDwDRuLfAdY6UVBXDms4BdjwoVkBq7dm
  12 1NPCP7RpG2uxKZh9rPMDdCj7FRLMLPcQm5
  13 1HYZwL6r4Ky2ZUbBYGx3NPdGLPcDb4xVve
  14 1PfNM7DFKNCs9aSUxTwhUqJV6hLYf5Gkho

  $ addrgen -format "BASE PRIVKEY" test 10 5
  test10 L58m9BShPM6H5H53VDq3jaZraajYZbxNg6UNAPKJ9S5Ks8FnZBFk
  test11 L17oBUvAg4dovbLtcFk9Fba1DqDCiqVWdkaWj8RsdJuccsENLoom
  test12 L2uJsQR6rcxdrATnaARB37g58338reHFChg6r4CaoTxmmEmro7tm
  test13 KwyKmCbKtYYH1Ktd1stftghs1BfdYKu1DGUJaamgbuSafHTwjnRY
  test14 KyPfxQ2koXBfpmyknxDCbpkkztjFRx9JmLQmhnWxEAv2YGDoRash

  $ addrgen -random 5
  L1QWEUaAGW2fqL7Xv72gdzuvXvWiCGTsx6Zyjp4QCEMWyTxQBkSs 1MH4tAB2RwCVWeeWiC5JLCWWjeD6SpctmF
  KwtjvSRmpGqBJYcKmWvVdnSqFdgcSdBnF1DfEFEry6kKVxia4ECG 16K8BCWMTDZvQUijBfYrB83v1uraKJqKkV
  L4WNciS28Himvz8nkZNnuYsh7LaS34SSfcvTpMxEzwwdoNcCoEzn 17P9Mm1QW94NpPvpmUr8cAuoLREWdcebY1
  Kwm2KmrZPbLiuDYiWm1HJ5K8nj5X8V1mWkMCLFrp3qJawci3U7tE 1P21Z5M8zCThxneHFNLFwNXjiwWF74PwX7
  L3spoNwkVp5znAPBfECLvhnXxZTjfWjG86X8wP81r68BjnjWzrtM 1PvhAp91jboUM6BerLTymqMABeXEoohYe4

License:
  MIT
