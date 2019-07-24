docker built .

# docker run -d -p 5000:80 --name Nginx -v /Users/lvtao/code/Docker/Nginx/Files/Vue_t:/usr/share/nginx/html python_docker_app_test:latest
docker run -d -p 9999:5000 python_docker_app_test:latest

docker run -d -it -p 5000:5000 --name python_docker_test python_docker_app_test:latest

docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' python_docker_test

