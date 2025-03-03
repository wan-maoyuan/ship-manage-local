package global

import "path/filepath"

const (
	DefaultLevel      = "info"
	DefaultMaxLogSize = 20
	DefaultMaxLogAge  = 10
	DefaultMaxBackups = 5

	CreateBatchSize = 300
)

var (
	ResourceDir = "resources"
	DBFilesDir  = filepath.Join(ResourceDir, "db_files")

	DefaultAddr = ":8888"
)
