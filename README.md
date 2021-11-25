# keccak256-circom [![Test](https://github.com/vocdoni/keccak256-circom/workflows/Test/badge.svg)](https://github.com/vocdoni/keccak256-circom/actions?query=workflow%3ATest)

Keccak256 hash function (ethereum version) implemented in [circom](https://github.com/iden3/circom). Spec: https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.202.pdf

**Warning**: WIP, this is an experimental repo.

## Status
Initial version works, compatible with Ethereum version of Keccak256.

It needs around `150848` (`151k`) constraints. 
> For context: [Rapidsnark](https://github.com/iden3/rapidsnark) proof generation time:
> - 1.1M constraints -> 7 seconds (8 CPU)
> - 128M constraints -> <2min (64 CPU)
