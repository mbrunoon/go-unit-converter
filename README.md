# 🧮 Go Unit Converter

A simple and extensible unit conversion API written in Go.  
Supports unit conversions like meters to kilometers, and can be easily extended to support other categories such as temperature, weight, etc.

## 🚀 Features

- Convert units from one to another (e.g., `kilometers → meters`)
- Validate unit categories (e.g., distance, temperature, weight)
- Easy to add new conversion formulas
- RESTful API with JSON support
- Includes tests

---

## 📦 Project Structure

```
go-unit-converter/
├── converter/           # Core conversion logic
├── routes/              # HTTP handlers and router
├── tests/               # Unit and integration tests
├── main.go              # Application entrypoint
└── go.mod
```

---

## 🔧 Installation

```bash
git clone https://github.com/your-username/go-unit-converter.git
cd go-unit-converter
go mod tidy
```

---

## 🧪 Running the Tests

From the root of the project:

```bash
go test ./...
```

---

## ▶️ Running the API

```bash
go run main.go
```

The server will start on `http://localhost:8080`

---

## 📡 API Endpoints

### `GET /`

Returns available unit conversion formulas.

**Example response:**

```json
{
  "meters": ["kilometers"],
  "kilometers": ["meters"]
}
```

---

### `POST /converter`

Converts a value from one unit to another.

**Request body:**

```json
{
  "value": 10,
  "from": "kilometers",
  "to": "meters"
}
```

**Successful response:**

```json
{
  "value": 10000
}
```

**Error response (invalid unit):**

```json
{
  "error": "kilometers can only be converted to [meters]"
}
```

---

## ➕ Adding New Units

To add a new unit:

1. Add the formula in `converter/formulas` map.
2. Update the `unitCategories` map with the new unit and its category.
3. Optionally write tests in `/tests`.

Example (add miles):

```go
var formulas = map[string]map[string]func(float64) float64{
  "miles": {
    "kilometers": func(f float64) float64 { return f * 1.60934 },
  },
  "kilometers": {
    "miles": func(f float64) float64 { return f / 1.60934 },
  },
}

var unitCategories = map[string]string{
  "miles":      "distance",
  "kilometers": "distance",
}
```

---

## 🙌 Contributing

Pull requests are welcome.  
For major changes, please open an issue first to discuss what you would like to change.

---

*README generated with the help of AI.*