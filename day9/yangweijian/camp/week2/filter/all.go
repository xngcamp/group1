package filter

import "net/http"

func All(w http.ResponseWriter, r *http.Request, m map[string]interface{}) bool {
	if Cors(w, r, m) {
		if Auth(w, r, m) {
			return true
		}
	}
	return false
}