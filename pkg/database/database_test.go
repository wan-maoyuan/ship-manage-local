package database

import (
	"ship-manage-local/pkg/config"
	"testing"

	"github.com/sirupsen/logrus"
)

func init() {
	c, err := config.New()
	if err != nil {
		logrus.Fatalf("配置信息加载失败: %v", err)
	}

	c.Show()
}

func TestInitDatabase(t *testing.T) {
	dsn := "sqlserver://sa:smart_ship123..@127.0.0.1:1433?database=ISM_Vsl_9977551&encrypt=disable"

	if err := InitDatabase(dsn); err != nil {
		t.Errorf("InitDatabase: %v", err)
		return
	}

	t.Logf("sqlserver InitDatabase success")
}
