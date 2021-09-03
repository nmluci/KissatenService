package models

type Member struct {
	ID    uint
	Name  string
	Point uint32
}

type Members map[string]*Member
