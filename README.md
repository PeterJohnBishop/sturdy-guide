# sturdy-guide - Docker Deployment

docker build -t sturdy-guide-app:latest .
docker tag sturdy-guide-app {docker_username}/sturdy-guide-app:latest
docker push {docker_username}/sturdy-guide-app:latest
docker pull {docker_username}/sturdy-guide-app:latest
docker run -p 8080:8080 {docker_username}/sturdy-guide-app:latest

