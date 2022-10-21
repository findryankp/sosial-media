package helpers

import "sosialMedia/configs"

func ById(id, model interface{}) (bool, error) {
	query := configs.DB.First(&model, id)
	if query.Error != nil {
		return false, query.Error
	}
	return true, nil
}
