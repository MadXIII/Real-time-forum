local-run: 
	docker build -t forum .
	docker run -p 8383:8383 forum

all:
	(cd cmd && go build && ./cmd)