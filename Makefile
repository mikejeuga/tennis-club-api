repo=$(shell basename "`pwd`")
gopher:
	@git init
	@touch .gitignore
	@touch Dockerfile
	@touch docker-compose.yml
	@go mod init github.com/mikejeuga/$(repo)
	@go get github.com/adamluzsi/testcase
	@go get -u github.com/gorilla/mux
	@go mod tidy


run:
	@go run ./cmd/main.go

t: test
test:
	@make ut at

ut: unit-test
unit-test:
	@go test -tags=unit ./... -v


at: acceptance-test
acceptance-test:
	@docker-compose -f docker-compose.yml up -d
	@go test -tags=acceptance ./... -v
	@docker-compose down

ic: init
init:
	@git add .
	@git commit -m "Initial commit"
	@git remote add origin git@github.com:mikejeuga/${repo}.git
	@git branch -M main
	@git push -u origin main

c: commit
commit:
	@git add .
	@git commit -m "$m"
	@git pull --rebase
	git push
