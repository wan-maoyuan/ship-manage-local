package database

import (
	"encoding/csv"
	"github.com/sirupsen/logrus"
	"os"
	"ship-manage-local/pkg/model"
	"strconv"
	"strings"
	"testing"
	"time"
)

const (
	CatLogPath    = "../../resources/data/pms_cat_log.csv"
	GroupPath     = "../../resources/data/pms_group.csv"
	EquipmentPath = "../../resources/data/pms_equipment.csv"
	ComponentPath = "../../resources/data/pms_component.csv"
)

func TestNewPMS(t *testing.T) {
	dir := "../../resources/db_files"
	_, err := newPMS(dir)
	if err != nil {
		t.Errorf("newPMS: %v", err)
	}
}

func TestBatchInsertPMSCatLog(t *testing.T) {
	dir := "../../resources/db_files"
	if err := InitDatabase(dir); err != nil {
		t.Errorf("InitDatabase: %v", err)
		return
	}

	file, err := os.Open(CatLogPath)
	if err != nil {
		t.Errorf("open CatLogPath file error: %v", err)
		return
	}
	defer file.Close()

	var lines [][]string

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			logrus.Errorf("read file content error: %v", err)
			break
		}

		lines = append(lines, record)
	}

	var catLogList []*model.PMSCatLog
	for _, line := range lines[1:] {
		catLogID := line[0]

		catLogList = append(catLogList, &model.PMSCatLog{
			VesselIMO:         2371283782,
			VesselName:        "test",
			CatLogID:          catLogID,
			CatLogDescription: line[1],
			Editor:            "system",
			Source:            "system",
			ModifyTime:        time.Now().UTC(),
			CreateTime:        time.Now().UTC(),
		})
	}

	if err := BatchInsertPMSCatLog(catLogList); err != nil {
		t.Errorf("BatchInsertPMSCatLog error: %v", err)
	}
}

func TestQueryPMSCatLog(t *testing.T) {
	dir := "../../resources/db_files"
	if err := InitDatabase(dir); err != nil {
		t.Errorf("InitDatabase: %v", err)
		return
	}

	catLogList, err := QueryPMSCatLog("")
	if err != nil {
		t.Errorf("QueryPMSCatLog error: %v", err)
		return
	}

	for _, catLog := range catLogList {
		t.Logf("%+v", catLog)
	}
}

func TestBatchInsertPMSGroup(t *testing.T) {
	dir := "../../resources/db_files"
	if err := InitDatabase(dir); err != nil {
		t.Errorf("InitDatabase: %v", err)
		return
	}

	file, err := os.Open(GroupPath)
	if err != nil {
		t.Errorf("open GroupPath file error: %v", err)
		return
	}
	defer file.Close()

	var lines [][]string

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			logrus.Errorf("read file content error: %v", err)
			break
		}

		lines = append(lines, record)
	}

	var groupList []*model.PMSGroup
	for _, line := range lines[1:] {
		catLogID := line[2]
		groupID := line[0]

		groupList = append(groupList, &model.PMSGroup{
			VesselIMO:         2371283782,
			VesselName:        "test",
			CatLogID:          catLogID,
			CatLogDescription: line[3],
			GroupID:           groupID,
			GroupDescription:  line[1],
			Editor:            "system",
			Source:            "system",
			ModifyTime:        time.Now().UTC(),
			CreateTime:        time.Now().UTC(),
		})
	}

	if err := BatchInsertPMSGroup(groupList); err != nil {
		t.Errorf("BatchInsertPMSGroup error: %v", err)
	}
}

func TestBatchInsertPMSEquipment(t *testing.T) {
	dir := "../../resources/db_files"
	if err := InitDatabase(dir); err != nil {
		t.Errorf("InitDatabase: %v", err)
		return
	}

	file, err := os.Open(EquipmentPath)
	if err != nil {
		t.Errorf("open EquipmentPath file error: %v", err)
		return
	}
	defer file.Close()

	var lines [][]string

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			logrus.Errorf("read file content error: %v", err)
			break
		}

		lines = append(lines, record)
	}

	var equipmentList []*model.PMSEquipment
	for _, line := range lines[1:] {
		catLogID := line[3]
		groupID := line[2]
		equipmentID := strings.Split(line[0], ".")[1]

		catLog, err := QueryPMSCatLog(catLogID)
		if err != nil {
			logrus.Errorf("QueryPMSCatLog error: %v", err)
			continue
		}
		if len(catLog) == 0 {
			continue
		}

		group, err := QueryPMSGroup(catLogID, groupID)
		if err != nil {
			logrus.Errorf("QueryPMSGroup error: %v", err)
			continue
		}
		if len(group) == 0 {
			continue
		}

		equipmentList = append(equipmentList, &model.PMSEquipment{
			VesselIMO:            2371283782,
			VesselName:           "test",
			CatLogID:             catLogID,
			CatLogDescription:    catLog[0].CatLogDescription,
			GroupID:              groupID,
			GroupDescription:     group[0].GroupDescription,
			EquipmentID:          equipmentID,
			EquipmentDescription: line[1],
			Editor:               "system",
			Source:               "system",
			ModifyTime:           time.Now().UTC(),
			CreateTime:           time.Now().UTC(),
		})
	}

	if err := BatchInsertPMSEquipment(equipmentList); err != nil {
		t.Errorf("BatchInsertPMSEquipment error: %v", err)
	}
}

func TestBatchInsertPMSComponent(t *testing.T) {
	dir := "../../resources/db_files"
	if err := InitDatabase(dir); err != nil {
		t.Errorf("InitDatabase: %v", err)
		return
	}

	file, err := os.Open(ComponentPath)
	if err != nil {
		t.Errorf("open ComponentPath file error: %v", err)
		return
	}
	defer file.Close()

	var lines [][]string

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}

			logrus.Errorf("read file content error: %v", err)
			break
		}

		lines = append(lines, record)
	}

	var componentList []*model.PMSComponent
	for _, line := range lines[1:] {
		catLogID := line[5]
		groupID := line[6]

		if len(strings.Split(line[0], ".")) != 3 {
			continue
		}

		equipmentID := strings.Split(line[0], ".")[1]
		componentID := strings.Split(line[0], ".")[2]

		hour, err := strconv.Atoi(line[2])
		if err != nil {
			logrus.Errorf("convert string to int error: %v", err)
			continue
		}

		catLog, err := QueryPMSCatLog(catLogID)
		if err != nil {
			logrus.Errorf("QueryPMSCatLog error: %v", err)
			continue
		}
		if len(catLog) == 0 {
			continue
		}

		group, err := QueryPMSGroup(catLogID, groupID)
		if err != nil {
			logrus.Errorf("QueryPMSGroup error: %v", err)
			continue
		}
		if len(group) == 0 {
			continue
		}

		equipment, err := QueryPMSEquipment(catLogID, groupID, equipmentID)
		if err != nil {
			logrus.Errorf("QueryPMSEquipment error: %v", err)
			continue
		}
		if len(equipment) == 0 {
			continue
		}

		componentList = append(componentList, &model.PMSComponent{
			VesselIMO:            2371283782,
			VesselName:           "test",
			CatLogID:             catLogID,
			CatLogDescription:    catLog[0].CatLogDescription,
			GroupID:              groupID,
			GroupDescription:     group[0].GroupDescription,
			EquipmentID:          equipmentID,
			EquipmentDescription: equipment[0].EquipmentDescription,
			ComponentID:          componentID,
			ComponentDescription: line[1],
			DailyRunningHours:    hour,
			Location:             "",
			Editor:               "system",
			Source:               "system",
			ModifyTime:           time.Now().UTC(),
			CreateTime:           time.Now().UTC(),
		})
	}

	if err := BatchInsertPMSComponent(componentList); err != nil {
		t.Errorf("BatchInsertPMSComponent error: %v", err)
	}
}
