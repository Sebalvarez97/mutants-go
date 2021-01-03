# MUTANTS
API RESTful to analize a dna chain to find mutant sequences.

## Hosting

This API is hosted in Google Kubernetes Engine, using the structure detailed in the file Cloud_API_Architecture.png. The yaml sources needed to deploy it are in /kubernetes. For the mongodb statefulset the yaml is in /mongodb with the instructions to setup it.

As this is a kubernetes solution, it could be deployed in any other kubernetes engine (AWS, Openshift, etc).

It is a scalable solution, this means that can perform well in low or big load.

The public IP tu consume the API is http://104.197.202.46/

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

***Important*** The token expires every 15 minutes and you can refresh it if it was created less than 20 minutes ago. If you need to refresh it use this request:

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

***Important*** Only NxN matrix are allowed. The values allowed for the nitrogen bases are only A,T,C or G.

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

## Get dna stats
This endpopint allows you to get some statistics of the dna chains previously analized.

The format of the response is the following:
```josn
{
  "count_mutant_dna": 40,
  "count_human_dna": 100,
  "ratio": 0.4
}
```
Here we have:
- **count_mutant_dna**: the number of mutant dna chains analized.
- **count_human_dna**: the number of human (not mutant) dna chains analized.
- **ratio**: the ratio between this two values.

***Important*** In case we have 0 humans the ratio will be 1 (one). Otherwise this value is 

count_mutant_dna / count_human_dna

##### REQUEST:
GET &rarr; /mutant/stats

Header: "Authentication: Bearer ${token}"

##### RESPONSE:
200 OK

Body:
```josn
{
  "count_mutant_dna": {quantity_of_mutants},
  "count_human_dna": {quantity_of_humans},
  "ratio": {quantity_of_mutants/quantity_of_humans}
}
```

## Error responses

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

400-Bad Request

Body:

```json
{
    "code": 400,
    "message": "Value entered is not valid: ....."
}
```
&rarr; if the request is not the expected

500-Internal Server Error
Body:

```json
{
    "code": 500,
    "message": "VInternal error..."
}
```
&rarr; if the server failed to perform the request
