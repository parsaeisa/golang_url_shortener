# technical_test_golang

URL shortener , written in golang 
it has two api's : 


/get_shortened_url : 
in this request , you can post a json wich has url in it . 
then in the response it will return the shortened version of that url . 

shortened_url :
for example in 
"http://localhost:8080/cKHeYx6rabt"
wich the last part is shortened url , it will automatically redirect to original url

I used redis for caching data and gin-gonic framework for http requests.
also for encoding I used base85 and MG5 hash system .

with tests . 

## Run Server
- Clone the repository
```shell script 
$ git clone https://github.com/parsaeisa/technical_test_golang
```

- Install dependencies 
```shell script
$ go get .
```

- Run the server
```shell script
$ go run .
```