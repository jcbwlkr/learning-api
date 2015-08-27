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
you can hit by browsing to `http://localhost:8080/posts/` and you can view your
site by browsing to `http://localhost:8080/site/`

For this API we are fetching, creating, updating, and deleting a Blog Post
resource for a simple blog application. A post entity looks like

```json
{
  "id": 1,
  "user": "jane",
  "message": "Hello, world!"
}
```

The following routes are available:

## GET /posts
Fetch all posts from the database.

* Response Body: Array of posts like `[{"id": 1 ...}, {"id": 2 ...}]`
* Response Status: `200 OK`

## POST /posts
Create a new post resource. `POST` the HTTP method should not be confused with
"post" the name of our resource. That is just a coincidence. If our resource
was named "widget" we would issue a `POST` to `/widgets`.

* Request Body: A post object (without id) like `{"user": "jane", "message": "hello"}`

* Response Body: The created post `{"id": 3, "user": "jane", "message": "hello"}`
* Response Status: `201 Created` on success, `403 Bad Request` if the request
  body has invalid JSON

## GET /posts/:id
Fetch a single post from the database. Replace `:id` with the id of the post you want.

* Response Body: A single post like `{"id": 1, "user": "jane", "message": "Hello, World!"}`
* Response Status: `200 OK` on success, `403 Bad Request` if id is not an
  integer, `404 Not Found` if there is no post for that id.

## PUT /posts/:id
Update a post. Replace `:id` with the id of the post you want to update.

* Request Body: The updated post, id optional. `{"user": "JANE", "message": "Howdy"}`

* Response Body: The updated post like `{"id": 1, "user": "JANE", "message": "Howdy"}`
* Response Status: `200 OK` on success, `403 Bad Request` if id is not an
  integer or if the provided JSON is invalid.

## DELETE /posts/:id
Delete a post. Replace `:id` with the id of the post you want to delete.

* Response Body: None.
* Response Status: `204 No Content` on success, `403 Bad Request` if id is not
  an integer.
