Bitcoin Address Generator

Usage:
  addrgen [-sequence] [-format "ID PUBKEY"] PREFIX [START_ID=1] [COUNT=1000]
  addrgen -random [-format "PRIVKEY PUBKEY"] [COUNT=1000]

Options:
  -sequence       the private key is SHA256($PREFIX$ID)
  -random         generate in random
  -format         available fields ID PRIVKEY PUBKEY BASE
  -version        show version
  -help           show help

Example:
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
  KxDxPJZyBdidNhiqGGJERmwFFKwyEVfmybBSjqmSTsxryLpEFnJS 19Q7n2bTgUAdQKpDxdxaSw5Jgnqh3iFRS
  L1HTfa1XkkWydpM75Dx2LP3PbSiECPNA5kyxtAhSdLR6b3kc5cwq 14uYxoVXN32HvwViZPEBkzocQZJv7DLgoZ
  L4dhtxq8RSGH79JAdGNhpxDivys6cuw295jviwTRoGtoZoYbWwmY 1875UPaXiUuHygXcjoxkPCGtqf9BvqcVnH
  L2htKiF4JD8ovytTWtgTySkn8fpU6yfLDnzhkA3CwXutxRdoj5Ef 1JSW3GEieJgCDxfD3fJ6GbaFvmDRPXzg1D
  L21LvyVxSptCW79jXnH7PVsbbYgxrQNzYhVZ36raanD8uk6Q1zLw 13vXHgpd3QGLKu7Wciq9uECdQuZfRmSW6d

License:
  MIT
