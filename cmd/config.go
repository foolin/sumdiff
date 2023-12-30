package cmd

type Config struct {
	Verbose bool   `json:"verbose"`
	Format  string `json:"format"` //Format: table/json/csv/yaml
	Output  string `json:"output"` //Output file
}
