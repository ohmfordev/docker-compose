docker build -t ohmfordevdocker/server1:0.1 -f Dockerfile1 .
docker build -t ohmfordevdocker/server2:0.1 -f Dockerfile2 .

docker push ohmfordevdocker/server1:0.1
docker push ohmfordevdocker/server2:0.1
