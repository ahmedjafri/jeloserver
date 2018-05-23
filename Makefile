
build:
	docker build -t jelo .

run:
	docker run -it -p 3000:3000 jelo