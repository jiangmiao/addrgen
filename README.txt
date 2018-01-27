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
  addrgen -seed [LENGTH=32]

Options:
  -sequence       the private key is SHA256($PREFIX$ID)
  -random         generate in random
  -format         available fields ID PRIVKEY PUBKEY BASE
  -seed           generate random string
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
  KwiEo2kXaEEZwr1BErrRURx4n1LMHJwfsbK4pjc5NzKw6zYRVWKm 19ta7wGgYzGygA4xA69ooVsPWzqmX41ysQ
  L4kNSVmv1FZ41DpSGQ7M7mNcrP4EQ5E3eWKxfYfoKZZoDnvkf89o 1D4TtqEiQSAE3ehkEBkUHEF94bAmZGPoK6
  KzYZh73XV4nndjemXDeW4j4o96RmfeAgs5gMmTqYJ34LLdEAinBy 17p6w6oBWhAf4gNAqQThkqtpuFDLoCPbKN
  L2wLu11y5gxv1fqUEEuEjGsh45MDB1hP4UtrDtYMiSFQnHKfowzg 1AKMjZCryr5gtSxQzMrkDDiAQ2UwwV2Qfv
  L2KQuWdnHxzUTpRSMJXtnkgtjqjeJCrXPXvgPh9EYXujcdPVB79x 171mWwADqjg4JThA9CHcykGXNih3eGQ1mu

Security Tips:
  The source and algorithm of addrgen is public. So the privkey could be enumerated.
  A long random prefix is highly recommended for security reason.

  # generate 32 bytes random string
  $ addrgen -seed
  azGJNtWp1QSPBmxC5uz3kf273jCPhmy7

License:
  MIT
