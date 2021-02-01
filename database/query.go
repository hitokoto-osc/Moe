package database

// APIRecord 定义了 `hitokoto_api_v1_server` 的映射结构
type APIRecord struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	URL  string `db:"url"`
	Desc string `db:"desc"`
}

// GetHitokotoAPIHostList 用于查询接口记录集合
func GetHitokotoAPIHostList() ([]APIRecord, error) {
	var records []APIRecord
	err := DB.Select(&records, "SELECT `id`, `name`, `url`, `desc` FROM `hitokoto_api_v1_server`")
	return records, err
}
