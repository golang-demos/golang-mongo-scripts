## Golang + Mongo - Demo Scripts
The main function triggers seeder logic which inserts dummy data into the database.

## Test Run
```
docker-compose -f docker-compose.dev.yml up -d
go run api/main.go
```

### Mongodb CLI access
```
mongo --username mongoadmin --password secret
```


### Watch for changes to rerun
```
nodemon --exec go run main.go --signal SIGTERM
```


## Author
### Vinay Jeurkar

<p>
  <a href="https://www.linkedin.com/in/vinay-jeurkar/" rel="nofollow noreferrer" style="text-decoration:none;"><img src="https://img.shields.io/badge/LinkedIn-0077B5?style=flat&logo=linkedin&logoColor=white" /></a> 
	&nbsp; 
  <a href="https://github.com/vinay03" rel="nofollow noreferrer" style="text-decoration:none;"><img src="https://img.shields.io/badge/GitHub-100000?style=flat&logo=github&logoColor=white" /></a> 
	&nbsp; 
  <a href="https://twitter.com/Vinay_Jeurkar" rel="nofollow noreferrer" style="text-decoration:none;"><img src="https://img.shields.io/badge/Twitter-1DA1F2?style=flat&logo=twitter&logoColor=white" /></a>
</p>