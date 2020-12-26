# mutants
Cerebro API to find mutants
## Analize a dna chain
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
