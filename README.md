# technical_test_golang

## URL shortener , written in golang 
it has two api's : 


- get_shortened_url : 
in this request , you can post a json wich has url in it . 
then in the response it will return the shortened version of that url . 

- shortened_url :
for example in 
"http://localhost:8080/cKHeYx6rabt"
wich the last part is shortened url , it will automatically redirect to original url

--------
stacks :
- gin-gonic framework
- redis 
- docker 

encodings :
- base58 bitcoin encoding

tests :
- testify package
- miniredis package
--------

### with CI pipeline
in github actions 

--------

## Run Server
- Clone the repository
```shell script 
$ git clone https://github.com/parsaeisa/technical_test_golang
```

- Run docker composer 
```shell script
$ docker-compose -f docker-compose.yml up --build
```
