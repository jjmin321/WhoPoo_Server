# 서버 실행 
.PHONY: run
run:
	@go run main.go

# 깃에 저장
.PHONY: push
push:
	@git push