package Models

type CORSConfig struct {
	AllowedOrigins []string
	AllowedHeaders []string
	AllowedMethods []string
	ExposedHeaders []string
}

var GlobalCORSConfig = CORSConfig{
	AllowedOrigins: []string{"http://localhost", "http://95.163.208.72:3000", "http://localhost:3000"},
	AllowedHeaders: []string{"If-Modified-Since", "Cache-Control", "Content-Type", "Range"},
	AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
	ExposedHeaders: []string{"Content-Length", "Content-Range"},
}
