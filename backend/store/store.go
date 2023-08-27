package store

import "jv/team-tone-tuner/model"

type UserStore struct {
	db *model.Client
}

type QuestionStore struct {
	db *model.Client
}

type BfiReportStore struct {
	db *model.Client
}

func NewUserStore(db *model.Client) *UserStore {
	return &UserStore{db: db}
}

func NewQuestionStore(db *model.Client) *QuestionStore {
	return &QuestionStore{db: db}
}

func NewBfiReportStore(db *model.Client) *BfiReportStore {
	return &BfiReportStore{db: db}
}
