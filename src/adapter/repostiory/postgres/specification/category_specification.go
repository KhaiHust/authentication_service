package specification

import (
	"github.com/KhaiHust/authen_service/core/entity/dto"
)

func ToCategorySpecification(spe *dto.CategorySpec) string {
	rawQuery := "WHERE 1=1 "
	if spe.Name != nil {
		rawQuery += "AND name LIKE " + *spe.Name
	}
	if spe.OrderBy != nil {
		rawQuery += "ORDER BY " + *spe.OrderBy
	} else {
		rawQuery += "ORDER BY updated_at "
	}
	if spe.Direct != nil {
		rawQuery += *spe.Direct
	} else {
		rawQuery += "DESC"
	}
	return rawQuery
}
func ToCountCategorySpecification(spe *dto.CategorySpec) string {
	rawQuery := "WHERE 1=1 "
	if spe.Name != nil {
		rawQuery += "AND name LIKE " + *spe.Name
	}
	return rawQuery
}
