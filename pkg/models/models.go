package models

type DBModel interface {
	IsModel() bool
}

type BaseModel struct{}

func (bm BaseModel) IsModel() bool {
	return true
}
