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

### Features

- User registration: users can register for a new account
- User authentication: users can log in to the application using their email and password
- Image upload: users can upload images to the application
- Image viewing: users can view all images uploaded to the application
- Image deletion: users can delete only their images from the application
- Image editing: users can edit only their images in the application
- Album creation: users can create albums for their images
- Album viewing: users can view all albums created by any user
- Album deletion: users can delete only their albums from the application
