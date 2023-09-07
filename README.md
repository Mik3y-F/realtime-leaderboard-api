# Realtime Leaderboard API

This is a realtime leaderboard API built using go.

## Features

* Realtime updates: The leaderboard API provides realtime updates for scores and rankings.
* Dockerized: The API can be easily run using Docker Compose.

## Endpoints

The API exposes the following endpoints:

* `POST /scores`: Creates a new player score entry.
* `GET /scores/top/{n}`: Retrieves the top n players.

* `POST /players`: Creates a new player.
* `GET /players`: Retrieves a list of all players.
* `GET /players/{id}`: Retrieves a specific player by ID.
* `DELETE /players/{id}`: Deletes a specific player by ID.
* `GET /players/{id}/score`: Retrieves the Total score of a specific player.
* `PUT /players/{id}`: Updates the details of a specific player.

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
