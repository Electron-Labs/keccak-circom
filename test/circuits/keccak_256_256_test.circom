pragma circom 2.0.0;

include "../../circuits/keccak.circom";

component main = Keccak(32*8, 32*8);
