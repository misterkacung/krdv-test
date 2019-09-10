### Instalation
Make sure you already have docker installed

## how to build this
docker build --no-cache -t ydhprs/krdv-test:test1 .

## Usage example

- http://localhost/ <-- index 
- http://localhost/404 <-- http not found
- http://localhost/400 <-- http bad request
- http://localhost/401 <-- http authentication

For viewing http response code, you can use curl for view the header
