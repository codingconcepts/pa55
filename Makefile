build:
	GOOS=windows go build -o pa55_windows .
	GOOS=linux go build -o pa55_linux .
	GOOS=darwin go build -o pa55_darwin .