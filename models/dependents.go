package models

// IConfig ... IConfig
type IConfig interface {
	SetConfiguration(key, value string)
	GetConfiguration(key string) (string, bool)
}

// ConfigDeps ... ConfigDeps
type ConfigDeps struct {
	dep map[string]string
}

// SetConfiguration ... SetConfiguration
func (config *ConfigDeps) SetConfiguration(key, value string) {
	if len(config.dep) == 0 {
		config.dep = make(map[string]string)
	}
	config.dep[key] = value
}

// GetConfiguration ... GetConfiguration
func (config *ConfigDeps) GetConfiguration(key string) (string, bool) {

	if len(config.dep) == 0 {
		return "", false
	}

	value, exists := config.dep[key]
	return value, exists

}
