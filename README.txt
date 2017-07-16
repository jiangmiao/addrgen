Bitcoin Address Generator

Usage:
  addrgen [-sequence] [-format "ID PUBKEY"] PREFIX [START_ID=1] [COUNT=1000]
  addrgen -random [-format "PRIVKEY PUBKEY"] [COUNT=1000]

Options:
  -sequence       the private key is SHA256($PREFIX$ID)
  -random         generate in random
  -format         available fields ID PRIVKEY PUBKEY BASE
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
  L2G6iajrFr2w54SATZXUgc9C7A3Hb5JPAKseS8iQqLPzNkKPa8rq 1BVivwzPJf8dPZS7j6ap7tPPwKtQ5tAzo8
  KwrBDFgfQkB7ivuDHuQpnGHFKna54dPzk8kaLQva2GVPCxXhyVyt 1KuzAEQoRDvRk4czNSSx8SRgjnuzZGi8tp
  Kwuw5AFUtdbpE1cAjdx9JhdYsY29noNR4o1yE2V1v3gz637zDmac 1C3dfrUeKWWmo9skBn9yii5DSrcjaB41y
  L5AmB56e4Gx7B5oXDcHg3MjDuD3wiFKnSmccHXT43wN4gmvwUdU5 1g12HNzvNvqdFuCdf8fqfsVWreibbfFAD
  Kyrm69Uw3mGMhnqFbv6ihkhWoH3Ph7TRhfLVMcqLgd7Wfd3KK3Xk 1AFXToNybKPW4PuKMY97iWuVpn7iExEbVq

License:
  MIT
