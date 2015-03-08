package plectrum

// Returns a copy of the object
func Pick(m map[string]interface{}, keys []string) map[string]interface{} {
  ret := make(map[string]interface{})
  for _, key := range keys {
    if val, ok := m[key]; ok {
      ret[key] = val
    }
  }
  return ret
}
