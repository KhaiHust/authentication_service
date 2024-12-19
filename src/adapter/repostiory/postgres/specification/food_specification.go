package specification

import "github.com/KhaiHust/authen_service/core/entity/dto"

func ToGetFoodSpecification(params *dto.FoodParams) (string, []interface{}) {
	query := "SELECT * FROM foods WHERE 1 = 1 "
	args := make([]interface{}, 0)
	if params.UserID != nil {
		query += "AND created_by = ? "
		args = append(args, *params.UserID)
	}
	if params.Name != nil {
		query += "AND name = ? "
		args = append(args, *params.Name)
	}
	if params.Type != nil {
		query += "AND type = ? "
		args = append(args, *params.Type)
	}
	if params.OrderBy != nil {
		query += " ORDER BY " + *params.OrderBy
	} else {
		query += " ORDER BY updated_at"
	}
	if params.Direct != nil {
		query += " " + *params.Direct
	} else {
		query += " DESC"
	}
	return query, args
}
func ToCountFoodSpecification(params *dto.FoodParams) (string, []interface{}) {
	query := "SELECT COUNT(*) FROM foods WHERE 1 = 1 "
	args := make([]interface{}, 0)
	if params.UserID != nil {
		query += "AND created_by = ? "
		args = append(args, *params.UserID)
	}
	if params.Name != nil {
		query += "AND name = ? "
		args = append(args, *params.Name)
	}
	if params.Type != nil {
		query += "AND type = ? "
		args = append(args, *params.Type)
	}
	return query, args
}
