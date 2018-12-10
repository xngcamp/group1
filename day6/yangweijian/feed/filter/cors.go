package filter

import "net/http"

// 前置过滤器，用于跨域
func Cors(w http.ResponseWriter, r *http.Request, m map[string]interface{}) bool {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Request-Method", "GET, POST, DELETE, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Range")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	if r.Method == "OPTIONS" {
		requestMethod := r.Header.Get("Access-Control-Request-Method")
		if requestMethod == http.MethodDelete ||
			requestMethod == http.MethodGet ||
			requestMethod == http.MethodPut ||
			requestMethod == http.MethodPost {
			w.WriteHeader(200)
			return false
		} else {
			w.WriteHeader(405)
			return false
		}
	}
	return true
}