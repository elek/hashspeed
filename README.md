Test code to compare hash speed. 

Copied from https://github.com/minio/sha256-simd/

## Run

Run and display the results in pivot table (you can run the test multiple times):

```
go test -run=NONE -bench=. >> out

go install golang.org/x/perf/cmd/...@latest

cat out | grep Bench | awk '{print $1}' | awk -F '/' '{print $2}' | sort | uniq | xargs -IFILE bash -c 'cat out | grep FILE | sed 's/FILE//g' > FILE.txt'

benchstat *.txt
```