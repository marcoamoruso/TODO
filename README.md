# Simple TODO list

Simple TODO list using Rest and http with Go, Echo, Gorm, MySQL and Docker

## Requirements

Docker

## Installation


```bash
git clone https://github.com/marcoamoruso/TODO.git
cd TODO
```

## Running Containers

To build containers:
```bash
docker-compose build
```
To start containers in background:
```bash
docker-compose up -d
```
To run client and test TODO list with CLI:
```bash
docker-compose exec client go run client.go
```
