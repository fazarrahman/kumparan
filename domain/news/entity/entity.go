package entity

// News ...
type News struct {
	Id     int64  `db:"id"`
	Author string `db:"author"`
	Body   string `db:"body"`
}
