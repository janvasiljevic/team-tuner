package store

import (
	"context"
	"jv/team-tone-tuner/model"
	"jv/team-tone-tuner/model/user"

	"github.com/google/uuid"
)

func (us UserStore) CreateANewStudentWithAnswers(githubUsername string, ctx context.Context) *model.User {
	newStudentUser := us.db.User.Create().
		SetGithubUsername(githubUsername).
		SaveX(ctx)

	// get all questions from database
	questions := us.db.BfiQuestion.Query().AllX(ctx)

	// create all the empty answers for the new user
	bulkCreateAnswer := make([]*model.BfiAnswerCreate, len(questions))

	for i, question := range questions {
		bulkCreateAnswer[i] = us.db.BfiAnswer.Create().
			SetBfiQuestion(question).
			SetStudentID(newStudentUser.ID)
	}

	// save all the answers
	us.db.BfiAnswer.CreateBulk(bulkCreateAnswer...).SaveX(ctx)

	return newStudentUser
}

func (us UserStore) GetUserInfo(userId uuid.UUID, ctx context.Context) (*model.User, error) {
	user, err := us.db.User.Query().WithBfiReport().Where(user.IDEQ(userId)).Only(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}
