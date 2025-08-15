NAME=neighborhood_watch

.PHONY: up down re clean fclean

up:
	@printf "🚀 Launching ${NAME}...\n"
	@docker compose up -d

down:
	@printf "🛑 Stopping ${NAME}...\n"
	@docker compose down -v || true

re:
	@printf "♻️ Rebuilding ${NAME}...\n"
	@$(MAKE) down
	@docker compose build --no-cache
	@$(MAKE) up

clean: down
	@printf "🧹 Cleaning up docker system...\n"
	@docker system prune -a -f

fclean:
	@printf "☢️  Nuking all docker configurations...\n"
	@if [ "$$(docker ps -qa)" ]; then docker stop $$(docker ps -qa); fi || true
	@docker system prune --all --volumes --force
	@docker compose down --remove-orphans -v || true