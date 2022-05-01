# Gin API to access MySQL 
## router: localhost:8080/mysql to fetch all iplogs
## router: localhost:8080/mysql/ip to fetch a single iplog
### example: http://localhost:8080/mysql/10.128.2.1
## data will be returned as JSON
### example:     
{
    "IP": "10.128.2.1",
    "時間": "[29/Nov/2017:06:58:55",
    "網址": "GET /login.php HTTP/1.1",
    "狀態": 200
}
