# Realtime Leaderboard API

This is a realtime leaderboard API built using go.

## Features

* Realtime updates: The leaderboard API provides realtime updates for scores and rankings.
* Dockerized: The API can be easily run using Docker Compose.

## Endpoints

The API exposes the following endpoints:

* `/leaderboard`: Retrieves the leaderboard with scores and rankings.
* `/scores`: Retrieves the scores of all players.
* `/player/{id}`: Retrieves the score and ranking of a specific player.

## Getting Started

To start the API, run the following command:

```bash
docker compose -f local.yml up
```

the above command will have ther server running on port 8080.

## Generating Queries

The API uses SQLC for generating go from sql queries. To generate queries, run the following command:

```bash
sqlc generate -f configs/db/sqlc.yaml
```
