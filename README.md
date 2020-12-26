# mutants
Cerebro API to find mutants
## Analize a dna chain
To analize a dna chain we need to arrange a group of nitrogen bases in an array like the one in the example.

This array represents a matrix of the dna chain arrangement.

Only NxN matrix are allowed. The values allowed for the nitrogen bases are only A,T,C or G.

The method analize sequences of 4 nitrogen bases. In the example we can see this sequences like [A,A,A,A] or [G,G,G,G].

If the API found 2 or more sequences like that in the matrix input, the dna is mutant, else the dna is not from a mutant.

Here we have two examples from a No-Mutant and a Mutant.

No-Mutant &rarr; [
        "ATGCGA",
        "CAGTGC",
        "TTATGT",
        "AGAGTG",
        "CCCTTA",
        "TCACTG"
    ]
Mutant &rarr; [
        "ATGCGA",
        "CAGTGC",
        "TTATGT",
        "AGAAGG",
        "CCCCTA",
        "TCACTG"
    ]
##### REQUEST:
POST &rarr; /mutant

Body: 
```json
{
   "dna":[
      "ATGCGA",
      "CAGTGC",
      "TTATGT",
      "AGAAGG",
      "CCCCTA",
      "TCACTG"
   ]
}
```
##### RESPONSE:
200 OK &rarr; if the analized dna is mutant

403-Forbidden &rarr;  if the analized dna is not mutant
