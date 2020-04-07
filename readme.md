Для запуска:

1. docker/docker-compose.override.yml.dist -> docker/docker-compose.override.yml (настроить как нужно)
2. middleware/clients.json.dist ->  middleware/clients.json (настроить как нужно)
3. docker/resources/nginx/lua/white_list.dist -> docker/resources/nginx/lua/white_list (настроить как нужно)

cd docker
docker-compose up

По указанному порту будет доступны 2 блоа url:
1. /public/{поисковый url meilisearch} - публичная часть, закрытая токеном
2. /{поисковый url meilisearch} - приватная часть, закрытая white_list

Документация на meilisearch:
https://docs.meilisearch.com