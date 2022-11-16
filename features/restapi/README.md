# [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)

from https://go.dev/doc/tutorial/web-service-gin

## Endpoints

- `/albums`
  - GET – Get a list of all albums, returned as JSON.
  - POST – Add a new album from request data sent as JSON.
- `/albums/{id}`
  - GET – Get an album by its ID, returning the album data as JSON.

## Notes

in this tutorial, we use in-memory data store instead of a database for simplicity.

## Usage

Execute the following command to run the tutorial:

```bash
make build
bin/tutorial -feature restapi

# or

make run
```

Then, the restapi server will be running on port 8080.

Get to request the list of all albums:


```bash
curl -X GET localhost:8080/albums
```

Add a new album:

```bash
curl -X POST localhost:8080/albums -d '{"title":"The Dark Side of the Moon","artist":"Pink Floyd","price":10.99}' -H "Content-Type: application/json"
```

Get an album by its ID:

```bash
curl -X GET localhost:8080/albums/1
```