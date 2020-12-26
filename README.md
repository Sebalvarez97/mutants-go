# MUTANTS
Cerebro API to find mutants

## Authentication
The API is setted with JWT authentication consume endpoints.

Users allowed are:

        username:"admin"
        password:"admin"
        
        username:"magneto"
        password:"magneto"

### Login and Obtaining a JWT
We need a JWT to consume the API. To obtain it you need to login with one of the allowed users.

##### REQUEST:
POST &rarr; /auth/login

Body: 
```json
{
   "username":"magneto",
   "password":"magneto"
}
```
##### RESPONSE:
```json
{
    "code": 200,
    "expire": "2020-12-26T17:46:50-03:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDkwMTU2MTAsImlkIjoibWFnbmV0byIsIm9yaWdfaWF0IjoxNjA5MDE0NzEwfQ.TlX3Y5tRqSkH-A5vvLvbE1VcCH6K_OhP1LWkK13IOS0"
}
```
You can find the JWT in the response body.token. This is used for authentication to the API sending it in header "Authentication: Bearer ${token}".

***Important* The token expires every 15 minutes and you can refresh it if it was created less than 20 minutes ago. If you need to refresh it use this request:

##### REQUEST:
GET &rarr; /auth/refresh_token 

Header: "Authentication: Bearer ${token_you_want_to_refresh}"
##### RESPONSE:
```json
{
    "code": 200,
    "expire": "2020-12-26T17:46:50-03:00",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDkwMTU2MTAsImlkIjoibWFnbmV0byIsIm9yaWdfaWF0IjoxNjA5MDE0NzEwfQ.TlX3Y5tRqSkH-A5vvLvbE1VcCH6K_OhP1LWkK13IOS0"
}
```

## Analize a dna chain
To analize a dna chain we need to arrange a group of nitrogen bases in an array like the one in the example.

This array represents a matrix of the dna chain arrangement.

***Important* Only NxN matrix are allowed. The values allowed for the nitrogen bases are only A,T,C or G.

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

Header: "Authentication: Bearer ${token}"

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

401-Unauthorized 

Body:
```json
{
    "code": 401,
    "message": "auth header is empty"
}
```
&rarr; if the JWT token for authentication is missing

401-Unauthorized 

Body:
```json
{
    "code": 401,
    "message": "Token is expired"
}
```
&rarr; if the JWT token for auth is expired

401-Unauthorized

Body:

```json
{
    "code": 403,
    "message": "you don't have permission to access this resource"
}
```
&rarr; if the resource is not allowed for the user
