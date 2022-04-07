# keccak256-circom [![Test](https://github.com/vocdoni/keccak256-circom/workflows/Test/badge.svg)](https://github.com/vocdoni/keccak256-circom/actions?query=workflow%3ATest)

Keccak256 hash function (ethereum version) implemented in [circom](https://github.com/iden3/circom). Spec: https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.202.pdf

**Warning**: WIP, this is an experimental repo.

## Status
Initial version works, compatible with Ethereum version of Keccak256.

## Usage
```
// make sure to include from your copy of circomlib
include "circomlib/circuits/gates.circom";
include "circomlib/circuits/sha256/xor3.circom";
include "circomlib/circuits/sha256/shift.circom";

var INPUT_BITS = 1024; // number of bits of the input message as a multiple of 8 (one byte)
component keccak = Keccak(INPUT_BITS, 256);
for (var i = 0; i < INPUT_BITS; i++) {
    keccak.in[i] <== msg[i];
}
for (var i = 0; i < 512; i++) {
    out[i] <== keccak.out[i];
}
```

It needs around `150848` (`151k`) constraints. 
> For context: [Rapidsnark](https://github.com/iden3/rapidsnark) proof generation time:
> - 1.1M constraints -> 7 seconds (8 CPU)
> - 128M constraints -> <2min (64 CPU)
