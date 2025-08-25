NAME=neighborhood_watch

.PHONY: up down re clean fclean

up:
	@printf "🚀 Launching ${NAME}...\n"
	@docker compose up -d

down:
	@printf "🛑 Stopping ${NAME}...\n"
	@docker compose down

re:
	@printf "♻️ Rebuilding ${NAME}...\n"
	@$(MAKE) down
	@docker compose up -d --build

prune:
	@printf "🧹 Pruning unused docker resources...\n"
	@docker system prune -f
	@docker volume prune -f
	@docker network prune -f

clean: down
	@printf "🧹 Cleaning up docker system...\n"
	@docker system prune -a -f

fclean:
	@docker ps -qa | xargs -r docker stop
	@docker system prune --all --force
	@docker network prune --force
	@docker compose -f docker-compose.yml down --remove-orphans || true