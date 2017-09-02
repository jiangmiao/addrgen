Bitcoin Address Generator

Installation:
  go get github.com/jiangmiao/addrgen

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
  L5cAZMXZtUnUTgPxQnmKWdSJuBR3DQP6AXnexb9VbX8ckPEhhQkU 13zPGTa9shGfi6FkH24xmkEj2MAJKBwSZd
  KzTrwvrSzJ7mVXgR7VzQ7GefDj58bj2RbcJsW1hbN5MKUeW78EqQ 1PimV8TXJ5Jt6R8diSTazBEyhiA9VqVFwE
  L29iQGxCTRtt8DZWNDTTBxmC3zWiw52hMNcJWTdp7GqmAhFGxTmj 18cEXd2hEChqWejkAqLYw85AWEW7hsJZEz
  KxKHXG7sDMkY5saeGB4L9Yx7rqMmofTmCCCXy8726ALRDxaPjz6m 1Hnrcr1GTsp3qDDkPDR8oNHcXGMusuNJ8L
  KwPHu7vUgyceNYnk8kcHVepwEKyE2MBZBLHqop2M6xZfxEVnceNV 1Eea7DuGVbStDPCHkUXi2knazuwVdQqqVL

License:
  MIT
