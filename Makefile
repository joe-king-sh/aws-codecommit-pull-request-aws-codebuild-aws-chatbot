all: gobuild sambuild

sambuild:
	sam build	

gobuild:
	cd ./src/comment-build-status-to-pull-request && \
	GOOS=linux go build main.go