# PacksMath API

PacksMath is a lightweight HTTP API that calculates the optimal number of item packs required to fulfill customer orders, using predefined pack sizes. It ensures the least number of items are shipped, and among those, the fewest number of packs are used.

## 🧩 Endpoints Overview

### ✅ Register a new pack size

**POST** `/packs`

Registers a new pack size.

```bash
curl -X POST http://localhost:3000/packs \
  -H "Content-Type: application/json" \
  -d '{"size": 500}'
```

**Success response:**

```
HTTP 204 No Content
```

**Conflict response (already registered):**

```json
HTTP 409 Conflict
{
  "message": "pack size already registered"
}
```

---

### 📦 List all available pack sizes

**GET** `/packs`

Returns all registered pack sizes.

```bash
curl http://localhost:3000/packs
```

**Success response:**

```json
[250, 500, 1000]
```

**Error response:**

```json
HTTP 500 Internal Server Error
{
  "message": "unknown error"
}
```

---

### ❌ Delete a specific pack size

**DELETE** `/packs/{size}`

Deletes the pack size passed in the path.

```bash
curl -X DELETE http://localhost:3000/packs/500
```

**Success response:**

```
HTTP 204 No Content
```

**Not found response:**

```json
HTTP 404 Not Found
{
  "message": "pack size not found"
}
```

**Error response:**

```json
HTTP 500 Internal Server Error
{
  "message": "unknown error"
}
```

---

### 📐 Calculate packs for an order

**POST** `/orders`

Returns a mapping of pack sizes to quantity, optimized according to the business rules.

```bash
curl -X POST http://localhost:3000/orders \
  -H "Content-Type: application/json" \
  -d '{"order": 1250}'
```

**Success response:**

```json
HTTP 200 OK
{
  "1000": 1,
  "250": 1
}
```

**Error response:**

```json
HTTP 500 Internal Server Error
{
  "message": "unknown error"
}
```

---

## 🚀 Running the Project

You only need [Docker](https://www.docker.com/) installed.

To run the API:

```bash
make docker-run
```

The API will be available at: [http://localhost:3000](http://localhost:3000)
Swagger UI documentation: [http://localhost:3000/docs/index.html](http://localhost:3000/docs/index.html)

---

## 📬 Postman Collection

You can also test the API using a preconfigured Postman collection located at the root of this project. It includes all requests ready to use and allows for easy experimentation with different inputs and responses.
