package specification

import (
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"time"
)

func BuildGetMealPlanSpecification(params *dto.MealPlanParams) (string, []interface{}) {
	rawQuery := "WHERE 1 = 1 "
	args := make([]interface{}, 0)
	if params.UserID != nil {
		rawQuery += "AND user_id = ? "
		args = append(args, params.UserID)
	}
	if params.StartedDate != nil {
		rawQuery += "AND schedule >= ? "
		args = append(args, time.Unix(*params.StartedDate, 0).UTC())
	}
	if params.EndedDate != nil {
		rawQuery += "AND schedule <= ? "
		args = append(args, time.Unix(*params.EndedDate, 0).UTC())
	}
	return rawQuery, args
}
