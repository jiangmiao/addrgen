trap 'rm actual.txt' EXIT
go run addrgen.go test 1 50 > actual.txt
diff actual.txt expected.txt || {
	echo "test failed"
	exit 1
}
