# Golang API (Dockerize) with Car Sale project

## System Design Diagram

<p align="center"><img src='/docs/files/system_diagram.png' alt='Golang Web API System Design Diagram' /></p>

## Database Design Diagram

<p align="center"><img src='/docs/files/db_diagram.png' alt='Golang Web API System Design Diagram' /></p>

### How to run

### Docker start

```
docker compose -f "docker/docker-compose.yml" up -d --build
```

#### Web API

##### Run local manually [http://localhost:5005](http://localhost:5005)

##### Run in docker [http://localhost:9001](http://localhost:9001)

```
Token Url: http://localhost:5005/api/v1/users/login-by-username
Username: admin
Password: 12345678
```

#### Kibana

##### [http://localhost:5601](http://localhost:5601)

```
Username: elastic
Password: @aA123456
```

#### Grafana

##### [http://localhost:3000](http://localhost:3000)

```
Username: admin
Password: foobar
```

#### PgAdmin

##### [http://localhost:8090](http://localhost:8090)

```
Username: h.naimaei@gmail.com
Password: 123456
```

Postgres Server info:

```
Host: postgres_container
Port: 5432
Username: postgres
Password: admin
```

#### Prometheus

##### [http://localhost:9090](http://localhost:9090)

### Docker Stop

```
docker compose --file 'docker/docker-compose.yml' --project-name 'docker' down
```

### Linux

0. build Project and copy configuration

```
/src > go build -o ../prod/server ./cmd/main.go
/src > mkdir ../prod/config/ && cp config/config-production.yml ../prod/config/config-production.yml
```

1. Create systemd unit

```
sudo vi /lib/systemd/system/go-api.service
```

2. Service config

```
[Unit]
Description=go-api

[Service]
Type=simple
Restart=always
RestartSec=20s
ExecStart=/home/hamed/github/golang-clean-web-api/prod/server
Environment="APP_ENV=production"
WorkingDirectory=/home/hamed/github/golang-clean-web-api/prod
[Install]
WantedBy=multi-user.target
```

3. Start service

```
sudo systemctl start go-api
```

4. Stop service

```
sudo systemctl stop go-api
```

5. Show service logs

```
sudo journalctl -u go-api -e
```

## Project preview

## Swagger

<p align="center"><img src='/docs/files/swagger.png' alt='Golang Web API preview' /></p>

## Grafana

<p align="center"><img src='/docs/files/grafana.png' alt='Golang Web API grafana dashboard' /></p>

## Kibana

<p align="center"><img src='/docs/files/kibana.png' alt='Golang Web API grafana dashboard' /></p>
