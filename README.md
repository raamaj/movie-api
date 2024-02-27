# Movie APP
MOVIE APP is a web application for film management. This application provides basic CRUD functionality.

```
Contact: Rama Jayapermana
Email: jayapermanarama@gmail.com
```

## License
>This project is licensed under the Apache 2.0 License. For more information, please refer to the LICENSE file.


#### Run App
```
go run cmd/main.go
```
## List
>Retrieve the results for each movie.

```http
GET /Movies
```

## View
>Retrieve the detail of movie based on ID.

```http
GET /Movies/:id
```
| Parameter | Type      | Description   |
|:----------|:----------|:--------------|
| `id`      | `integer` | `ID of movie` |

## Insert
>Insert Movie.

```http
POST /Movies
```

#### Body
```javascript
{
    "description": "string",
    "image": "string",
    "rating": 10,
    "title": "string"
}
```

| Field | Type      | Description            |
|:------|:----------|:-----------------------|
| `title`  | `string`  | `Title of movie`       |
| `description`  | `string`  | `Description of movie` |
| `rating`  | `float`   | `Rating of movie`      |
| `image`  | `string` | `Image (URL) of movie` |

### Success Response

```javascript
{
    "data": {}, 
    "message": "string",
    "status_code": 200
}
```

## Update
>Update movie based in ID.

```http
PATCH /Movies/:id
```

| Parameter | Type      | Description   |
|:----------|:----------|:--------------|
| `id`      | `integer` | `ID of movie` |

#### Body
```javascript
{
    "description": "string",
    "image": "string",
    "rating": 10,
    "title": "string"
}
```

| Field | Type      | Description            |
|:------|:----------|:-----------------------|
| `title`  | `string`  | `Title of movie`       |
| `description`  | `string`  | `Description of movie` |
| `rating`  | `float`   | `Rating of movie`      |
| `image`  | `string` | `Image (URL) of movie` |

## Delete
>Delete movie based on ID.

```http
DELETE /Movies/:id
```
| Parameter | Type      | Description   |
|:----------|:----------|:--------------|
| `id`      | `integer` | `ID of movie` |

### Success Response

```javascript
{
    "data": {}, 
    "message": "string",
    "status_code": 200
}
```

### Error Response

```javascript
{
    "error": true,
    "message": "string",
    "status_code": 400
}
```

# Tech Stack
- Golang
- PostgresSQL
- Go Fiber
