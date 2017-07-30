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
  L4EwCUjgmp8iHuDdJVEzdKvhHetbcUNtNe8nzYQE7ri1ZxNpRPqE 18ZAeG1D1fTbqYNHzALBL1eKEqX8LV1rPn
  L1QiVPpgk3LwjjLSArREJ12ZZ3X5MGu6irz3MccnLFYdSehPZYzF 1Hvttz2DMwoEv8cJGcV4dAUrZViNgH3biQ
  KzCKQ8N73cBy5ArnAEGxGudC1oGjDkZBWdV7eD4V7C8EuMjqEwMD 1ADUbTKBnD1ED1MtFVZ1A4YXpbw1kK89hG
  L2FbQG8p9XECTr1JMTdQqbgT2NXJhuBPHgmh1jehnk4GjSMvU6do 1JbwbbWJccwV2JngGWmCFHNDQ2hketUJbh
  Ky7qVNaSEqkrDH96sjHm8NvyYn4LEVQoDfmBZwKrMohjmyZALkKU 16QvTUwRVTbihuvj94EYZX6xRxHKPyeDkH

License:
  MIT
