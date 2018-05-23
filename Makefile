
build:
	docker build -t jelo .

run:
	docker run -it -p 8080:3000 jelo