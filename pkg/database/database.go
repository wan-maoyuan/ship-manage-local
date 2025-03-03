package database

import (
	"fmt"
)

var db *Database

type Database struct {
	pms *PMS
}

func InitDatabase(dir string) error {
	pms, err := newPMS(dir)
	if err != nil {
		return fmt.Errorf("init pms database: %w", err)
	}

	db = &Database{
		pms: pms,
	}

	return nil
}

//func generateTimestampID() int64 {
//	return time.Now().UnixMicro() - time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC).UnixMicro()
//}
