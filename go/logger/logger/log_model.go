package logger

type LoggerConfig struct {
	Level   LogLevel
	Masking ConfigMasking
	Caller  Caller
}

type Caller struct {
	Disable   bool
	FieldName string
}

type ConfigMasking struct {
	FieldMap map[string]any
	Enabled  bool
}
