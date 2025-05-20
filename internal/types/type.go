package types

type Language struct {
	Name     string   `yaml:"name"`
	Database string   `yaml:"database"`
	Alias    []string `yaml:"alias"`
}

type MainConfig struct {
	Languages []Language `yaml:"languages"`
}

type Section struct {
	Name        string   `yaml:"name"`
	Alias       []string `yaml:"alias,omitempty"`
	Syntax      string   `yaml:"syntax"`
	Description string   `yaml:"description"`
}
