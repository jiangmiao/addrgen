Bitcoin Address Generator

Installation:
  go get github.com/jiangmiao/addrgen

  # or if meets any compile compatible error, try

  go get github.com/golang/dep/cmd/dep
  go get -d github.com/jiangmiao/addrgen
  cd $GOPATH/src/github.com/jiangmiao/addrgen
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
  KwjRkVWPYvCQ78EPs72dcWDDtTvnYWMDx25fB54xnbf81qbcMy19 1BqQkDXVhEXT6DebY7ksmipaHcYcSyzVX3
  L5jt7nMfVK8iRYeNyTTmbXcUvZBKYnMokbwzVZENvg3YYtwyEcM4 18ZanY5EN87Nopp4kxVXkwNMmiXgt9Qkgs
  Ky3uFueeBLyzUrJUuEyN1cqnNM4KqCx18m35sWS4R5p4JnBVzC1J 1NHJJ6EUmbw3QipXC7s1MP1A7UnxRLs5Vq
  KwFJ1bMhbgAMkw6sGGLFecoiZA8PD7JgSWRbpFKMLuDsPgPDGd94 1L2v9C6CyE5dWRYoq7UEYKpJrRUBLC7KnR
  L1VymFzUG7uwopZj53oxPbx4hfNszZ7snEfH8LDSqwKxm8FkQ4U5 1KxgYcG9NxBFfPi6LugcgZfaE4YdrLPHye

License:
  MIT
