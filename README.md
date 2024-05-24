
# Description

Gomonitor is simple realtime performance monitoring system like RAM, CPU usage. Built in microservices architecture for easy scaling. Designed with local setup in mind. Depending only on Docker for containerization for Kafka runs.

# Installation

### Kafka setup
```bash
  docker-compose up -d  
```
Command builds, starts, attaches to containers for Kafka and Zookeeper

### Run services
Before running project, keeping in mind you installed make package on your machine.

Using this command for running each service
```bash
  make run  
```

### Client connection
Clients use Websocket protocol to connect with Consumer service on port `50071` and this endpoint

```bash
  ws://localhost:50071/ws
```
