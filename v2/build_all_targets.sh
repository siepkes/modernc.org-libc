set -e
for tag in none # dmesg libc.membrk libc.memgrind
do
	echo "-tags=$tag"
	echo "GOOS=linux GOARCH=386"
	GOOS=linux GOARCH=386 go build -tags=$tag -v ./...
	GOOS=linux GOARCH=386 go test -tags=$tag -c -o /dev/null
	echo "GOOS=linux GOARCH=amd64"
	GOOS=linux GOARCH=amd64 go build -tags=$tag -v ./...
	GOOS=linux GOARCH=amd64 go test -tags=$tag -c -o /dev/null
done
