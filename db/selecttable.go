package db

import (
	"ginvue/moudle"
	"strconv"
)

func Selecttb(starttime, endtime string) (time_values moudle.TimeVlues) {

	var start_datas []string
	var end_dates []string
	var start_values = make([]int, 12)
	var end_values = make([]int, 12)

	//var end_dates = make([]string, 0)
	for i := 0; i < 9; i++ {
		start_datas = append(start_datas, starttime+"-0"+strconv.Itoa(i+1))
		end_dates = append(end_dates, endtime+"-0"+strconv.Itoa(i+1))
	}
	for i := 9; i < 12; i++ {
		start_datas = append(start_datas, starttime+"-"+strconv.Itoa(i+1))
		end_dates = append(end_dates, endtime+"-"+strconv.Itoa(i+1))
	}

	query := "select ym,count from (select concat(year,'-',month) as\n    ym,count(*) as count from time_count group by ym )\n    as yearmothon where ym like ? order by ym;"
	starttime = starttime + "-%"
	starts, err := Mysqldb.Query(query, starttime)
	//ends, err := db.Query(query,start_date)
	if err != nil {
		return
		//log.Fatal(err)
	}
	defer starts.Close()

	// 遍历查询结果
	for starts.Next() {
		//var id int
		var ym string
		var count int
		err := starts.Scan(&ym, &count)
		if err != nil {
			return
		}

		for i, date := range start_datas {
			if date == ym {
				start_values[i] = count
			}
		}
	}

	endtime = endtime + "-%"
	ends, err := Mysqldb.Query(query, endtime)
	//ends, err := db.Query(query,start_date)
	if err != nil {
		return
		//log.Fatal(err)
	}
	defer ends.Close()
	for ends.Next() {
		//var id int
		var ym string
		var count int
		err := ends.Scan(&ym, &count)
		if err != nil {
			return
		}

		for i, date := range end_dates {
			if date == ym {
				end_values[i] = count
			}
		}
	}
	time_values.StartValues = start_values
	time_values.EndValues = end_values
	time_values.StartDates = start_datas
	time_values.EndDates = end_dates
	return

	if err = starts.Err(); err != nil {
		return
	}

	return
}
