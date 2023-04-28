package project

import "github.com/osg_template/internal/usecase/project"

type ProjectController struct {
	pro project.UseCase
}

func NewProjectController(pro project.UseCase) *ProjectController {
	return &ProjectController{pro: pro}
}
