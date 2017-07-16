output=`./addrgen`
example=`./addrgen test 10 5|sed 's/^/  /'`
example2=`./addrgen -format "BASE PRIVKEY" test 10 5|sed 's/^/  /'`
example3=`./addrgen -random 5|sed 's/^/  /'`
cat <<EOT
Bitcoin Address Generator

$output

Example:
  $ addrgen test 10 5
$example

  $ addrgen -format "BASE PRIVKEY" test 10 5
$example2

  $ addrgen -random 5
$example3

License:
  MIT
EOT
