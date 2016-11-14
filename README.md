# Docker-workshop 14.11.2016

## Excercise 2:

 - Create a Dockerfile for an nginx webserver
 - Use "nginx" as the base image
 - Nginx provides an HTTP server on port 80
 - The Configuration for the nginx-server is available in "consumer/nginx/conf.d/default.conf"
 - This file should be copied/added to the folder "/etc/nginx/conf.d/default.conf" of the docker image
 - There is also a static website under "consumer/www" which should be copied to the folder "/www" of the docker image
 - Start and test the created image
