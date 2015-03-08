package plectrum

// Returns a map containing only the values mapped by the given keys
func Pick(m map[string]interface{}, keys []string) map[string]interface{} {
  ret := make(map[string]interface{})
  for _, key := range keys {
    if val, ok := m[key]; ok {
      ret[key] = val
    }
  }
  return ret
}
