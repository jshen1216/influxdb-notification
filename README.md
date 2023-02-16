# Influxdb notification

## APIs:
- [POST] http://your_ip:5055/api/Alert/Email
- [POST] http://your_ip:5055/api/Alert/Line

## Introduce
InfluxDB (2.0 and above version) can set alert with notification methods : Slack, Pagerduty and http

This project is simply to received alert message from influx by http and send email or line to notify the alert message.

## A Simple Way to Set up Alert in Influxdb and use this project to send notification:
1. Set Up Rules in InfluxDB
2. Set Up Notification Endpoints in InfluxDB
    - Line using POST url: http://localhost:5055/api/Alert/Line
    - Email using POST url: http://localhost:5055/api/Alert/Email
3. Set Up Notification Rules
4. amend the file name of config.example.yml to config.yml
5. edit the config.yml (line token and email smtp)
6. `go run main.go` in your server to run the api

## Others
- Suggest to build the code to and execute file by `go build main.go` and set up it as a service in linux server so that you can enable it and run in background
- Swagger can be find through: http://localhost:5055/swagger/index.html

Thanks!