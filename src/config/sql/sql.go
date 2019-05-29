package sql

import (
	"fmt"
	"time"
	"tutor_platform/src/common"
)

const (
	QueryWeekMessageSQL    = "select * from message"
	QueryTutorSQL          = "select * from tutor "
	QueryMessageSQL        = "select * from message "
	QueryPageMessageSQL    = "select * from message"
	QueryPageMessageNumSQL = "select count(*) as num from message where city='上海市';"
	QueryStudentSQL        = "select * from student"
)

func QueryPageWhereCitySQL(city string, page int, pageSize int) string {
	sql := fmt.Sprintf("%v where city='%v' order by create_time desc limit %v,%v", QueryPageMessageSQL, city, page, pageSize)
	return sql
}

func QueryBatchIDSQL(idList []string, idName string) string {
	if len(idList) == 0 {
		return ""
	}
	sql := fmt.Sprintf("where %v in (", idName)
	for _, id := range idList {
		sql = fmt.Sprintf("%v%v,", sql, id)
	}
	return sql[:len(sql)-1] + ")"
}

func QueryWeekSQL() string {
	WeekTime := time.Now().AddDate(0, 0, -60).Unix()

	//WeekAgo := common.TimeToStr(WeekTime)
	//fmt.Println(WeekAgo)
	//WeekString := common.TimeFormat(WeekAgo)
	//fmt.Println(WeekString)
	Week := common.UnixToStr(WeekTime)
	sql := fmt.Sprintf("%v where status='发布中' and type_code=1 and create_time >= '%v' order by create_time desc limit 10", QueryWeekMessageSQL, Week)
	return sql
}

func QueryBatchStudentSQL(idList []string) string {
	return fmt.Sprintf("%v %v", QueryStudentSQL, QueryBatchIDSQL(idList, "user_id"))
}
