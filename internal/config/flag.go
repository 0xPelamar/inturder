package config

import "strings"

// For those flags that repeat more than once
type multiStringFlag []string
type wordListFlag []string

func (m *multiStringFlag) String() string { return "" }
func (m *multiStringFlag) Set(value string) error {
	*m = append(*m, value)
	return nil
}

// For those flags separated by comma

func (w *wordListFlag) String() string { return "" }
func (w *wordListFlag) Set(value string) error {
	words := strings.Split(value, ",")
	if len(words) > 1 {
		*w = append(*w, words...)
	} else {
		*w = append(*w, value)
	}
	return nil
}

func ParseFlages(config *Config) *Config {

}
