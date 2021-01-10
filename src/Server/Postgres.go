package Server

import (
	"Zombie/src/Utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func PostgresConnect(User string, Password string, info Utils.IpInfo) (err error, result bool, db *sql.DB) {
	dataSourceName := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", User,
		Password, info.Ip, info.Port, "postgres", "disable")
	db, err = sql.Open("postgres", dataSourceName)

	if err != nil {
		result = false
	}
	return err, result, db
}

func PostgresConnectTest(User string, Password string, info Utils.IpInfo) (err error, result bool) {
	err, result, db := PostgresConnect(User, Password, info)
	defer db.Close()
	if err == nil {
		defer db.Close()
		err = db.Ping()
		if err == nil {
			result = true
		}
	}

	return err, result
}

func PostgresQuery(User string, Password string, info Utils.IpInfo, Query string) (err error, Qresult []map[string]string) {
	err, _, db := PostgresConnect(User, Password, info)
	if err != nil {
		fmt.Println("connect failed,please check your input.")
	} else {
		defer db.Close()
		err = db.Ping()
		if err == nil {
			rows, err := db.Query(Query)
			if err == nil {
				Qresult = DoRowsMapper(rows)

			} else {
				fmt.Println("please check your query.")
				return err, Qresult
			}
		} else {
			fmt.Println("connect failed,please check your input.")
			return err, Qresult
		}
	}
	return err, Qresult
}