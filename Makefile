local-run: 
	docker build -t forum .
	docker run -p 8282:8282 forum

all:
	(cd cmd && go build && ./cmd)