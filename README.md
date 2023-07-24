# Nidus MQTT

Home Monitoring Messaging Service part of the
[Nidus](https://github.com/alexandrelamberty/nidus) project.

## Features

## Technologie and Frameworks

- [Docker](https://www.docker.com/)
- [Go](https://go.dev/)

## Usage

This application is part of a Docker stack and rely on a EMQX MQTT broker service. see:
[Nidus](https://github.com/alexandrelamberty/nidus) project to launch the
complete stack or only the `messaging` service.

### Run with Go

If the messaging service is up and running, create an .env file and fill it
accordingly with the `database` service configuration.

```properties
MQTT_SERVER=9fca54477c8ad4e70dc5e1084f884aad
MQTT_CLIENT_ID=d7a481461577ba4c3c4c6946cca7204b
MQTT_USERNAME=
MQTT_PASSWORD=
```
