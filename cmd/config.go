package cmd

type Config struct {
	Verbose bool   `json:"verbose"`
	Type    string `json:"type"`   //Type: table/json/csv/yaml
	Output  string `json:"output"` //Output file
}
