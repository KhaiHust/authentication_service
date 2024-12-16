package helper

func CalculateParameterForGetRequest(page, pageSize, total int64) (*int64, *int64, int64) {
	totalPage := int64(0)
	var nextPage *int64
	var prePage *int64

	if total%pageSize == 0 {
		totalPage = total / pageSize
	} else {
		totalPage = total/pageSize + 1
	}
	if page > 1 {
		temp := page - 1
		prePage = &temp
	}
	if page < totalPage {
		temp := page + 1
		nextPage = &temp
	}

	return nextPage, prePage, totalPage
}
