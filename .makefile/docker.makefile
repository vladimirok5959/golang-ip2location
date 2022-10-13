docker-build: all
	@-docker rmi ${DOCKER_IMG_NAME}:latest
	docker build -t ${DOCKER_IMG_NAME}:latest ./

docker-export:
	@-rm ./build/${DOCKER_IMG_NAME}.tar
	docker save ${DOCKER_IMG_NAME}:latest > ./build/${DOCKER_IMG_NAME}.tar

docker-import:
	@-docker rmi ${DOCKER_IMG_NAME}:latest
	docker load < ./build/${DOCKER_IMG_NAME}.tar

docker-push: docker-build
	docker tag ${DOCKER_IMG_NAME}:latest vladimirok5959/${DOCKER_IMG_NAME}:${VERSION}
	docker tag ${DOCKER_IMG_NAME}:latest vladimirok5959/${DOCKER_IMG_NAME}:latest
	docker login
	docker push vladimirok5959/${DOCKER_IMG_NAME}:${VERSION}
	docker push vladimirok5959/${DOCKER_IMG_NAME}:latest
	docker rmi vladimirok5959/${DOCKER_IMG_NAME}:${VERSION}
	docker rmi vladimirok5959/${DOCKER_IMG_NAME}:latest
