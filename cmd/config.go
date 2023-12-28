package cmd

type Config struct {
	Verbose bool   `json:"verbose"`
	Format  string `json:"format"` //Format: table/json/csv/yaml2
	Output  string `json:"output"` //Output write file
}
