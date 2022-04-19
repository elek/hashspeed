#!/usr/bin/env bash
rm *.txt
for algo in sha256 sha256simd blake3 xxh3 sha512; do
   ./hashspeed -algo $algo > $algo.txt
done

GODEBUG=cpu.avx2=off ./hashspeed -algo sha256 > sha256-noavx.txt

cat *.txt | head -n 3 | tail -n +2
./benchstat *.txt | head -n 3
