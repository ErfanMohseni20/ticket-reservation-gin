package helpers

import (
	"github.com/ErfanMohseni20/ticket-reservation-gin/internal/database"
	// "strconv"
)

func GetTotalCount(model interface{}) (int64) {
	var total int64
	database.DB.Model(&model).Count(&total)
	return total
}

// func ResponsePaginator(data []interface{}, responseList []interface{}) (responsetime []interface{}) {
// 	for _,var := range data {
// 		responseList = append(responseList , responseList{

//			})
//		}
//	}
// func GetNumber(perPage, page string , model interface{}) (perPage , page , offset ,total ,  totalPages) {
// 	perPageStr := perPage
// 	pageStr := page
// 	perPage, err := strconv.Atoi(perPageStr)
// 	if err != nil || perPage < 1 || perPage > 100 {
// 		perPage = 15
// 	}
// 	page, err := strconv.Atoi(pageStr)
// 	if err != nil || page < 1 {
// 		page = 1
// 	}
// 	offset := (page - 1) * perPage
// 	var total int64  
// 	database.DB.Model(&model).Count(&total)
// 	totalPages := int((total + int64(perPage) - 1) / int64(perPage))

// 	return (perPage,page,offset,total,totalPages)

// }
