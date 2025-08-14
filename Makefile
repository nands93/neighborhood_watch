NAME=neighborhood_watch

.PHONY: up down re clean fclean

up:
	@printf "Launch development configuration ${name}...\n";
	@docker compose up -d

down:
	@printf "🛑 Stopping ${NAME}...\n"
	@docker compose down -v || true

re:
	@printf "♻️ Rebuilding ${NAME}...\n"
	@$(MAKE) down
	@$(MAKE) up

clean: down
	@printf "🧹 Cleaning up docker system...\n"
	@docker system prune -a -f

fclean:
	@printf "☢️  Nuking all docker configurations...\n"
	@docker stop $$(docker ps -qa) || true
	@docker system prune --all --volumes --force
	@$(MAKE) down --remove-orphans