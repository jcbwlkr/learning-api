# learning-api

This is a tiny server used to learn the basics of how to make requests to an
API.

Download one of the pre built binaries from the
[releases](https://github.com/jcbwlkr/learning-api/releases) page or build from
source. If you are using a Mac then you want one of the Darwin binaries.

To get around a Cross Origin Request issue that you may encounter in your
browser you can also use this application to serve files in a directory on
localhost. Change to the directory of your web application and run the
executable. Leave it running. This will create an HTTP server that
you can hit by browsing to `http://localhost:8080/articles/` and you can view your
site by browsing to `http://localhost:8080/site/`

For this API we are fetching, creating, updating, and deleting an Article
resource for a simple news application. A article entity looks like

```json
{
  "id": 1,
  "user": "jane",
  "body": "Hello, world!"
}
```

The following routes are available:

## GET /articles
Fetch all articles from the database.

* Response Body: Array of articles like `[{"id": 1 ...}, {"id": 2 ...}]`
* Response Status: `200 OK`

## POST /articles
Create a new article resource.

* Request Body: A article object (without id) like `{"user": "jane", "body": "hello"}`

* Response Body: The created article `{"id": 3, "user": "jane", "body": "hello"}`
* Response Status: `201 Created` on success, `403 Bad Request` if the request
  body has invalid JSON

## GET /articles/:id
Fetch a single article from the database. Replace `:id` with the id of the article you want.

* Response Body: A single article like `{"id": 1, "user": "jane", "body": "Hello, World!"}`
* Response Status: `200 OK` on success, `403 Bad Request` if id is not an
  integer, `404 Not Found` if there is no article for that id.

## PUT /articles/:id
Update an article. Replace `:id` with the id of the article you want to update.

* Request Body: The updated article without id like `{"user": "JANE", "body": "Howdy"}`

* Response Body: The updated article like `{"id": 1, "user": "JANE", "body": "Howdy"}`
* Response Status: `200 OK` on success, `403 Bad Request` if id is not an
  integer or if the provided JSON is invalid.

## DELETE /articles/:id
Delete an article. Replace `:id` with the id of the article you want to delete.

* Response Body: None.
* Response Status: `204 No Content` on success, `403 Bad Request` if id is not
  an integer.
