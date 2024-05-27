# Photo Album

## Links

- [Frontend Repository](https://github.com/apella1/photo_album_ui)
- [Deployed Website](https://pt-album.vercel.app/)

## Technologies

- Go
- Go-Chi
- Postgres
- Goose
- SQLC
- Go JWT

### Hosting

- Heroku

## Tradeoffs (Ignored Limitations)

- For the sake of the project's simplicity, the image files are stored directly in the database.

### Setting up Locally

- Clone the repository
- Create a `.env` file at the root of the application and add values for the keys provided in the sample env file
- Run `make` at the root of the project to build and run the application
