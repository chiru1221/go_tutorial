# go tutorial
- module  
https://golang.org/doc/tutorial/create-module
- wiki  
https://golang.org/doc/articles/wiki/
- mysql  
http://go-database-sql.org/
- memo  



## Docker run
- Golang  
```
docker run -it -v $PWD:$PWD --network <NETWORK_NAME> --name <CONTAINER_NAME> <IMAGE_NAME>
```
- MySQL  
```
docker run -it -v $PWD:$PWD --network <NETWORK_NAME> --name <CONTAINER_NAME> -e MYSQL_ROOT_PASSWORD=mysql -e BIND-ADDRESS=0.0.0.0 -d -p 3306:3306 <IMAGE_NAME>
```

### MySQL after run  
```
docker exec -it <CONTAINER_NAME> bash
mysql -u root -p
```
