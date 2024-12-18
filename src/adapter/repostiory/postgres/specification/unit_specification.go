package specification

import "github.com/KhaiHust/authen_service/core/entity/dto"

func ToUnitSpecification(params *dto.UnitParamDto) string {
	rawQuery := "WHERE 1=1 "
	if params.Name != nil {
		rawQuery += "AND name LIKE " + *params.Name
	}
	if params.OrderBy != nil {
		rawQuery += "ORDER BY " + *params.OrderBy
	} else {
		rawQuery += "ORDER BY updated_at "
	}
	if params.Direct != nil {
		rawQuery += *params.Direct
	} else {
		rawQuery += "DESC"
	}
	return rawQuery
}
func ToCountUnitSpecification(params *dto.UnitParamDto) string {
	rawQuery := "WHERE 1=1 "
	if params.Name != nil {
		rawQuery += "AND name LIKE " + *params.Name
	}
	return rawQuery
}
