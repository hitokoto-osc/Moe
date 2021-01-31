package database

type APIRecord struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	URL  string `db:"url"`
	Desc string `db:"desc"`
}

func GetHitokotoAPIHostList() ([]APIRecord, error) {
	var records []APIRecord
	err := DB.Select(&records, "SELECT `id`, `name`, `url`, `desc` FROM `hitokoto_api_v1_server`")
	return records, err
}
