trap 'rm actual.txt' EXIT
go run addrgen.go test 1 50 > actual.txt
diff actual.txt expected.txt || {
	echo -e "\033[31mtest failed\033[0m"
	exit 1
}
echo -e "\033[32mtest passed\033[0m"
