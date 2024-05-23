package util

import (
  "encoding/json"
  "net/http"
)

func EncodeJSON(w http.ResponseWriter, data interface{}) error {
  return json.NewEncoder(w).Encode(data)
}

func DecodeJSON(r *http.Request, data interface{}) error {
  return json.NewDecoder(r.Body).Decode(data)
}