package goresource

import (
	"errors"
	crud "github.com/kordar/gocrud"
	"github.com/kordar/goi18n"
)

type ResourceService struct {
}

func (common *ResourceService) Search(body crud.SearchBody) crud.SearchVO {
	return crud.SearchVO{}
}

func (common *ResourceService) SearchOne(body crud.SearchBody) crud.SearchOneVO {
	return crud.SearchOneVO{}
}

func (common *ResourceService) Delete(body crud.RemoveBody) error {
	message := goi18n.GetSectionValue(body.Lang(), "resource.errors", "not.provided", "ini").(string)
	return errors.New(message)
}

func (common *ResourceService) Add(body crud.FormBody) (interface{}, error) {
	message := goi18n.GetSectionValue(body.Lang(), "resource.errors", "not.provided", "ini").(string)
	return nil, errors.New(message)
}

func (common *ResourceService) Update(body crud.FormBody) (interface{}, error) {
	message := goi18n.GetSectionValue(body.Lang(), "resource.errors", "not.provided", "ini").(string)
	return nil, errors.New(message)
}

func (common *ResourceService) Edit(body crud.EditorBody) error {
	message := goi18n.GetSectionValue(body.Lang(), "resource.errors", "not.provided", "ini").(string)
	return errors.New(message)
}

func (common *ResourceService) Configs(lang string) map[string]interface{} {
	return map[string]interface{}{}
}
