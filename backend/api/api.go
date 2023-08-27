package api

import (
	"jv/team-tone-tuner/model"
	"jv/team-tone-tuner/store"
)

type Api struct {
	db             *model.Client
	questionStore  *store.QuestionStore
	userStore      *store.UserStore
	bfiReportStore *store.BfiReportStore
}

func New(db *model.Client) *Api {
	return &Api{
		db:             db,
		questionStore:  store.NewQuestionStore(db),
		userStore:      store.NewUserStore(db),
		bfiReportStore: store.NewBfiReportStore(db),
	}
}
