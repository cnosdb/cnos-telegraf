# OpenTSDB JSON Style Parser Plugin

## Configuration

```toml
[[inputs.file]]
  files = ["example"]
  data_format = "opentsdb_json"
```

## Example

```json
[
  {
    "metric": "sys.cpu.nice",
    "timestamp": 1346846400,
    "value": 18,
    "tags": {
      "host": "web01",
      "dc": "lga"
    }
    },
  {
    "metric": "sys.cpu.nice",
    "timestamp": 1346846400,
    "value": 9,
    "tags": {
      "host": "web02",
      "dc": "lga"
    }
  }
]
```
