package controller

import "learntbd/model"

type MentorController struct {
	Model model.MentorModel
}

func (mc MentorController) GetAll([]model.Mentor, error) {
	// Ini Get All
	result, err := mc.Model.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Search() {

}
