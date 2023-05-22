package db

import (
	"log"
	"strconv"
)

func Selecttb(time string) (time_datas []string, values []int) {

	time_datas = make([]string, 0)
	values = make([]int, 12)
	//var end_dates = make([]string, 0)
	for i := 0; i < 12; i++ {
		time_datas = append(time_datas, time+"-"+strconv.Itoa(i+1))
		//end_dates = append(end_dates, end_date+"-"+strconv.Itoa(i+1))
	}

	query := "select ym,count from (select concat(year,'-',month) as\n    ym,count(*) as count from time_count group by ym )\n    as yearmothon where ym like '?-%' order by ym;"
	starts, err := Mysqldb.Query(query, 2023)
	//ends, err := db.Query(query,start_date)
	if err != nil {
		return nil, nil
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
			log.Fatal(err)
		}

		for i, date := range time_datas {
			if date == ym {
				values[i] = count

			} else {
				values[i] = 0
			}
		}
	}

	if err = starts.Err(); err != nil {
		return
	}

	return
}
