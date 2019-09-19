## Ani Lobby API

CRUD API backend built using [Golang](https://golang.org/) for use in my capstone project [Ani Lobby](https://github.com/rj-ortega/ani-lobby)

It also features two endpoints that will query a [Jikan's unofficial MyAnimeList API](https://jikan.docs.apiary.io/#) for anime by season or search query

Currently deployed at https://ani-lobby.herokuapp.com/api/v1/

#### Endpoints:

* for anime based on season - `https://ani-lobby.herokuapp.com/api/v1/seasons?year=<your_year_here>&season=<your_season_here>`

* for searching for an anime - `https://ani-lobby.herokuapp.com/api/v1/search?search=<your_query_here>`

* for anime database endpoints - `https://ani-lobby.herokuapp.com/api/v1/anime`

* for user database endpoints - `https://ani-lobby.herokuapp.com/api/v1/users`
