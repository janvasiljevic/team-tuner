package out

type QuestionerStatus struct {
	FinishedBfiQuestioner bool `json:"finished_bfi_questioner" validate:"required"`
}

func NewQuestionerStatus(finishedBfiQuestioner bool) *QuestionerStatus {
	return &QuestionerStatus{
		FinishedBfiQuestioner: finishedBfiQuestioner,
	}
}
