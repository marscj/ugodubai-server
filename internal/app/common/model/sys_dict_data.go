package model

type DictType struct {
	DictName string `json:"name"`
	Remark   string `json:"remark"`
}

// DictDataRes 字典数据
type DictData struct {
	DictValue string `json:"key"`
	DictLabel string `json:"value"`
	IsDefault int    `json:"isDefault"`
	Remark    string `json:"remark"`
}
