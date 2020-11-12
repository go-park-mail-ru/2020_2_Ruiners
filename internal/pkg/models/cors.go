package models

type CORSConfig struct {
	AllowedOrigins []string
	AllowedHeaders []string
	AllowedMethods []string
	ExposedHeaders []string
}
