# Docker-workshop 14.11.2016

## Excercise 1:

 - Create a Dockerfile for a statically linked binary
 - Use "scratch" as the base image
 - The service provides an HTTP server on port 8080, be sure that it will be accessible
 - Processes in docker-containers should not be started as root for security reasons

```# docker build . -t provider:latest
# docker run -p 8080:8080 provider:latest
```

