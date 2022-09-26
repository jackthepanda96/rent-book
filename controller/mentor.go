package controller

import "learntbd/model"

type MentorController struct {
	Model model.MentorModel
}

func (mc *MentorController) GetAllMentor() ([]model.Mentor, error) {
	var result []model.Mentor

	return result, nil
}

func Search() {

}
