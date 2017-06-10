output=`./addrgen`
example=`./addrgen test 10 5|sed 's/^/  /'`
cat <<EOT
Bitcoin Address Generator

$output

Example:
  $ addrgen test 10 5
$example

License:
  MIT
EOT
