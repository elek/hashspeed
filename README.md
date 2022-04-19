Test code to compare hash speed. 

Originally started from https://github.com/minio/sha256-simd/, but now it's different.

## Run

Run and display the results in pivot table (you can run the test multiple times):

```
env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

./all.sh

#or use:
#
#for algo in sha256 sha256simd blake3 xxh3; do ./hashspeed -algo $algo > $algo.txt; done
#

go install golang.org/x/perf/cmd/...@latest

benchstat *.txt
```
