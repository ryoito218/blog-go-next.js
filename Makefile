.PHONY: help up down restart build logs ps clean \
        db-shell db-tables \
        be-shell fe-shell

help:
	@echo "Targets:"
	@echo "  make up        : 起動（ビルド込み）"
	@echo "  make down      : 停止"
	@echo "  make restart   : 再起動（停止 -> 起動）"
	@echo "  make build     : ビルド"
	@echo "  make logs      : ログ表示"
	@echo "  make ps        : 状態表示"
	@echo "  make clean     : 停止（ボリューム削除）"
	@echo "  make db-shell  : MySQLに入る"
	@echo "  make db-tables : テーブル一覧表示"
	@echo "  make be-shell  : backendコンテナに入る"
	@echo "  make fe-shell  : frontendコンテナに入る"

up:
	docker compose up -d --build

down:
	docker compose down

restart:
	$(MAKE) down
	$(MAKE) up

build:
	docker compose build

logs:
	docker compose logs -f --tail=200

ps:
	docker compose ps

clean:
	docker compose down -v

db-shell:
	docker compose exec db mysql -u blog -pblog -D blog

db-tables:
	docker compose exec db mysql -u blog -pblog -D blog -e "SHOW TABLES;"

be-shell:
	docker compose exec backend sh

fe-shell:
	docker compose exec frontend sh
