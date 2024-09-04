package utils

import (
	"gorm.io/gorm"
	"reflect"
	"sort"
	"strings"
)

func CheckValueIn(target string, strArray []string, defaultValue string) string {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return target
	}
	return defaultValue
}

func CheckInModel(column string, model interface{}, defaultColumn string) string {
	modelV := reflect.TypeOf(model)
	if _, ok := modelV.FieldByName(column); ok {
		return column
	}
	return defaultColumn
}

func GetSortByStr(sort, order string, model interface{}) string {
	col := ToCamelString(sort)
	sort = CheckInModel(col, model, "created_at")
	orders := []string{"DESC", "ASC"}
	order = CheckValueIn(strings.ToTitle(order), orders, orders[0])
	sortBy := ToSnakeString(sort) + " " + order
	return sortBy
}

func Paginate(page int64, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
			break
		case pageSize <= 0:
			pageSize = 10
			break
		default:
			pageSize = 20
		}
		offset := (page - 1) * pageSize
		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
