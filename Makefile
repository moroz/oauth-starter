install:
	go get
	which modd || go install github.com/cortesi/modd/cmd/modd@latest
