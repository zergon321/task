WD := $(shell pwd)

all: build run

build:
	docker build -t task:0.0.1 .

run-json:
	docker run --mount type=bind,source=$(WD)/test.json,target=/bin/test.json task:0.0.1 /bin/test.json

run-csv:
	docker run --mount type=bind,source=$(WD)/test.csv,target=/bin/test.csv task:0.0.1 /bin/test.csv

run:
	docker run --mount type=bind,source=$(WD)/$(FILE),target=/bin/$(FILE) task:0.0.1 /bin/$(FILE)