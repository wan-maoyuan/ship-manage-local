package proto

type QueryPMSGroupReq struct {
	CatLogID string `form:"cat_log_id"`
	GroupID  string `form:"group_id"`
}

type QueryPMSEquipmentReq struct {
	CatLogID    string `form:"cat_log_id"`
	GroupID     string `form:"group_id"`
	EquipmentID string `form:"equipment_id"`
}

type QueryPMSComponentReq struct {
	CatLogID    string `form:"cat_log_id"`
	GroupID     string `form:"group_id"`
	EquipmentID string `form:"equipment_id"`
	ComponentID string `form:"component_id"`
}
