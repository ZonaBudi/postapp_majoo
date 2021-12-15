
## POST APP MAJO

### Requirment IMPORTANT
1. MUST USE MYSQL 8 because some query only support in version 8
2. IF YOU NEED unit test please use docker first, because some package using dockertest 
3. GOLANG 1.7 

### HOW TO RUN
1. Reconfig 
Change value file in config.yaml by your own example i using, dont change any key, because key used by app
```server:
  host: 0.0.0.0
  port: 8080
  secret_access: "majo_app_pos_secret"
mysql:
  host: localhost
  port: 3306
  password: password
  user: root
  database: "post" 
```

2. Get Dependecy

``` go get ./...```

3. Please Migrate Data First
Run SQL file in cmd/migration to your database 

4. RUN APP
``` go run cmd/api/main.go ```

### DOCUMENTATION API
please go to `localhost:${server.port}/swagger/index.html`

### TEST

```
go test ./...
```

### WE ARE USING
1. go-chi router framework
2. gorm orm
3. swaggo for documentation
4. za for logger
5. ozzo-validation for validation json request
6. viper for load configuration