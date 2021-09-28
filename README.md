# golang-microservices

Gathering with a friend who asked me to introduce him to microservices, we created this repo to upload what comes out of our sessions. 


## Notes

### To create local certs for TLS

```
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 365 
```