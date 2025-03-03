package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
	"ship-manage-local/pkg/global"
	"ship-manage-local/pkg/model"
	"time"
)

type PMS struct {
	*gorm.DB
}

func newPMS(dir string) (*PMS, error) {
	db, err := gorm.Open(
		sqlite.Open(filepath.Join(dir, "pms.db")),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold:             time.Second * 5,
					LogLevel:                  logger.Silent,
					IgnoreRecordNotFoundError: false,
					ParameterizedQueries:      true,
					Colorful:                  false,
				},
			),
			CreateBatchSize: global.CreateBatchSize,
		})

	if err != nil {
		return nil, fmt.Errorf("open pms database file: %w", err)
	}

	pmsList := []any{&model.PMSCatLog{}, &model.PMSGroup{}, &model.PMSEquipment{}, &model.PMSComponent{}}

	if err := db.AutoMigrate(pmsList...); err != nil {
		return nil, fmt.Errorf("auto migrate pms catlog: %w", err)
	}

	return &PMS{db}, nil
}

func BatchInsertPMSCatLog(catLogList []*model.PMSCatLog) error {
	if catLogList == nil {
		return fmt.Errorf("pms_cat_log list is nil")
	}

	database := db.pms.
		Model(&model.PMSCatLog{}).
		CreateInBatches(catLogList, global.CreateBatchSize)

	if err := database.Error; err != nil {
		return err
	}

	return nil
}

func QueryPMSCatLog(catLogID string) ([]*model.PMSCatLog, error) {
	var list []*model.PMSCatLog

	var condition = make(map[string]any)
	if catLogID != "" {
		condition["cat_log_id"] = catLogID
	}

	database := db.pms.
		Model(&model.PMSCatLog{}).
		Where(condition).
		Find(&list)

	if err := database.Error; err != nil {
		return nil, err
	}

	return list, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func BatchInsertPMSGroup(groupList []*model.PMSGroup) error {
	if groupList == nil {
		return fmt.Errorf("pms_group list is nil")
	}

	database := db.pms.
		Model(&model.PMSGroup{}).
		CreateInBatches(groupList, global.CreateBatchSize)

	if err := database.Error; err != nil {
		return err
	}

	return nil
}

func QueryPMSGroup(catLogID, groupID string) ([]*model.PMSGroup, error) {
	var list []*model.PMSGroup

	var condition = make(map[string]any)
	if catLogID != "" {
		condition["cat_log_id"] = catLogID
	}
	if groupID != "" {
		condition["group_id"] = groupID
	}

	database := db.pms.
		Model(&model.PMSGroup{}).
		Where(condition).
		Find(&list)

	if err := database.Error; err != nil {
		return nil, err
	}

	return list, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func BatchInsertPMSEquipment(equipmentList []*model.PMSEquipment) error {
	if equipmentList == nil {
		return fmt.Errorf("pms_equipment list is nil")
	}

	database := db.pms.
		Model(&model.PMSEquipment{}).
		CreateInBatches(equipmentList, global.CreateBatchSize)

	if err := database.Error; err != nil {
		return err
	}

	return nil
}

func QueryPMSEquipment(catLogID, groupID, equipmentID string) ([]*model.PMSEquipment, error) {
	var list []*model.PMSEquipment

	var condition = make(map[string]any)
	if catLogID != "" {
		condition["cat_log_id"] = catLogID
	}
	if groupID != "" {
		condition["group_id"] = groupID
	}
	if equipmentID != "" {
		condition["equipment_id"] = equipmentID
	}

	database := db.pms.
		Model(&model.PMSEquipment{}).
		Where(condition).
		Find(&list)

	if err := database.Error; err != nil {
		return nil, err
	}

	return list, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func BatchInsertPMSComponent(componentList []*model.PMSComponent) error {
	if componentList == nil {
		return fmt.Errorf("pms_component list is nil")
	}

	database := db.pms.
		Model(&model.PMSComponent{}).
		CreateInBatches(componentList, global.CreateBatchSize)

	if err := database.Error; err != nil {
		return err
	}

	return nil
}

func QueryPMSComponent(catLogID, groupID, equipmentID, componentID string) ([]*model.PMSComponent, error) {
	var list []*model.PMSComponent

	var condition = make(map[string]any)
	if catLogID != "" {
		condition["cat_log_id"] = catLogID
	}
	if groupID != "" {
		condition["group_id"] = groupID
	}
	if equipmentID != "" {
		condition["equipment_id"] = equipmentID
	}
	if componentID != "" {
		condition["component_id"] = componentID
	}

	database := db.pms.
		Model(&model.PMSComponent{}).
		Where(condition).
		Find(&list)

	if err := database.Error; err != nil {
		return nil, err
	}

	return list, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//func InsertPMSJobOrder(order *model.PMSJobOrder) error {
//	if order == nil {
//		return fmt.Errorf("job order is nil")
//	}
//
//	db := msDB.
//		Model(&model.PMSJobOrder{}).
//		Create(order)
//
//	if err := db.Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func QueryPMSJobOrder(page, size uint) ([]*model.PMSJobOrder, error) {
//	var list []*model.PMSJobOrder
//
//	if page == 0 {
//		page = 1
//	}
//
//	db := msDB.
//		Model(&model.PMSJobOrder{}).
//		Limit(int(size)).
//		Offset(int(page-1) * int(size)).
//		Find(&list)
//
//	if err := db.Error; err != nil {
//		return nil, err
//	}
//
//	return list, nil
//}
//
//func QueryPMSJobOrderCount() (int64, error) {
//	var count int64
//
//	db := msDB.
//		Model(&model.PMSJobOrder{}).
//		Count(&count)
//
//	if err := db.Error; err != nil {
//		return 0, err
//	}
//
//	return count, nil
//}
//
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//func InsertPMSWorkDone(work *model.PMSWorkDone) error {
//	if work == nil {
//		return fmt.Errorf("pms work done is nil")
//	}
//
//	db := msDB.
//		Model(&model.PMSWorkDone{}).
//		Create(work)
//
//	if err := db.Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func QueryPMSWorkDone(page, size uint) ([]*model.PMSWorkDone, error) {
//	var list []*model.PMSWorkDone
//
//	if page == 0 {
//		page = 1
//	}
//
//	db := msDB.
//		Model(&model.PMSWorkDone{}).
//		Limit(int(size)).
//		Offset(int(page-1) * int(size)).
//		Find(&list)
//
//	if err := db.Error; err != nil {
//		return nil, err
//	}
//
//	return list, nil
//}
//
//func QueryPMSWorkDoneCount() (int64, error) {
//	var count int64
//
//	db := msDB.
//		Model(&model.PMSWorkDone{}).
//		Count(&count)
//
//	if err := db.Error; err != nil {
//		return 0, err
//	}
//
//	return count, nil
//}
//
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
//
//func InsertPMSOverDueOrder(order *model.PMSOverDueOrder) error {
//	if order == nil {
//		return fmt.Errorf("pms work done is nil")
//	}
//
//	db := msDB.
//		Model(&model.PMSOverDueOrder{}).
//		Create(order)
//
//	if err := db.Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func QueryPMSOverDueOrder(page, size uint) ([]*model.PMSOverDueOrder, error) {
//	var list []*model.PMSOverDueOrder
//
//	if page == 0 {
//		page = 1
//	}
//
//	db := msDB.
//		Model(&model.PMSOverDueOrder{}).
//		Limit(int(size)).
//		Offset(int(page-1) * int(size)).
//		Find(&list)
//
//	if err := db.Error; err != nil {
//		return nil, err
//	}
//
//	return list, nil
//}
//
//func QueryPMSOverDueOrderCount() (int64, error) {
//	var count int64
//
//	db := msDB.
//		Model(&model.PMSOverDueOrder{}).
//		Count(&count)
//
//	if err := db.Error; err != nil {
//		return 0, err
//	}
//
//	return count, nil
//}
