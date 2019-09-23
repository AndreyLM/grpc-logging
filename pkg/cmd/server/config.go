package cmd

// Config - configuration for Server
type Config struct {
	GRPCPort   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBSchema   string
	LogPath    string
}
