# fictional-pancake

A Go server for a Vue SPA.

Initally following [How I set up a real world project with Go and Vue](https://www.freecodecamp.org/news/how-i-set-up-a-real-world-project-with-go-and-vue/).

## Environment variables

Either a 'proper' environment variable can be set, or add to the `.env` file. Copy `.example.env` to get default values.

`FRONTEND` is the location the front end will be served from. Defaults to `./frontend/dist`.
`PORT` is the port the server will listen for connections on. Default `3030`.
