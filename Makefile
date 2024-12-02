
.PHONY: .swag
.swag:
	cd ./backend && swag init

setup: docker-compose.yaml .swag
	docker-compose up --build -d

submit:
	git add .
	message=$(git diff --name-only --cached | sed ':a;N;$!ba;s/\n/, /g')
	git commit -m "Изменены файлы: $(message)"

