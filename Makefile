build:
		cd docker && docker build -t api-starwars-image -f Dockerfile ..
run:
docker run -d --name api-starwars-container -p 8090:8090 --network host --restart always api-starwars-image

remove:
		docker rm -f api-starwars-container

purge:
		docker rmi api-starwars-image
