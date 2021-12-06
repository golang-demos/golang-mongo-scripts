## Go + Mux + MongoDB Demo

### Dependencies
```
go get github.com/gorilla/mux
go get github.com/mongodb/mongo-go-driver/mongo
```

### Run MongoDB in docker
```
docker pull mongo
docker run --name some-mongo \
    -e MONGO_INITDB_ROOT_USERNAME=mongoadmin \
    -e MONGO_INITDB_ROOT_PASSWORD=secret \
    -p "27017:27017" \
    -d mongo
```
