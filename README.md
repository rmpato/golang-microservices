# golang-microservices

Gathering with a friend who asked me to introduce him to microservices, we created this repo to upload what comes out of our sessions. 


## Notes

### To create local certs for TLS

```
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 365 
```

### To run a RabbitMQ instance from docker 

```
docker run --detach --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management  
```

To open an AMQP connection, use: `amqp://localhost:5672` 

To browse the management UI `http://localhost:15672`
By default, user is `guest` as so is the password.