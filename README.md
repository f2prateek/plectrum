# Plectrum

 Plectrum lets you quickly pick specific keys in map.

## Example

```go
event := map[string]interface{}{
  "name": map[string]interface{}{
    "first": "prateek",
    "last":  "srivastava",
  },
  "username": "f2prateek",
  "age": 23
}

Pick(event, []string{"username", "age"})
// {"username": "f2prateek", "age": 23}

Pick(event, []string{"username", "foo"})
// {"username": "f2prateek"}
```