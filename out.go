package main
type CasbinRule struct {
	 // id 
	 Id	int64	`json:"id"`
	 // ptype 
	 Ptype	string	`json:"ptype"`
	 // v0 
	 V0	string	`json:"v0"`
	 // v1 
	 V1	string	`json:"v1"`
	 // v2 
	 V2	string	`json:"v2"`
	 // v3 
	 V3	string	`json:"v3"`
	 // v4 
	 V4	string	`json:"v4"`
	 // v5 
	 V5	string	`json:"v5"`
}
func (model CasbinRule) TableName() string {
	return "casbin_rule"
}