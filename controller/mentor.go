package controller

import "learntbd/model"

type MentorController struct {
	Model model.MentorModel
}

<<<<<<< HEAD
func (mc MentorController) GetAll([]model.Mentor, error) {
	// Ini Get All
	result, err := mc.Model.GetAll()
	if err != nil {
		return nil, err
	}
=======
func (mc *MentorController) GetAllMentor() ([]model.Mentor, error) {
	var result []model.Mentor
<<<<<<< HEAD
>>>>>>> GetAllMentor Controller
=======

>>>>>>> GetAllMentor 1
	return result, nil
}

func Search() {

}
