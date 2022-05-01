package controller

import (
	"fmt"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// parameters of mysql at localhost
// which username and passwd are default setting
const (
	USERNAME = "root"
	PASSWORD = "password"
	NETWORK  = "tcp"
	SERVER   = "127.0.0.1"
	PORT     = 3306
	DATABASE = "iplog"
)

// Ip represents data about a log record.
type Ip struct {
	IP     string  `json:"IP"` // json 回傳時IP 的 Key 對應到 IP
	Time   string  `json:"時間"` // json 回傳時Time 的 Key 對應到 時間
	Url    string  `json:"網址"`
	Status float64 `json:"狀態"`
}


// Get all iplogs LIMIT 100
func GetDatas(c *gin.Context) {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		c.JSON(http.StatusBadGateway, "Connect DB failed")
		return
	}
	if err := db.Ping(); err != nil {
		c.JSON(http.StatusBadGateway, "Connect DB failed")
		return
	}
	fmt.Println("DB iplog connected")

	// An datas slice to hold data from returned rows.
	var datas []Ip
	rows, err := db.Query("SELECT * FROM ip_log LIMIT 100")

	if err != nil {
		c.JSON(http.StatusBadRequest, "Connot Find Datas")
		return
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() { // rows.Next() 前往下一個項目。如果成功（還有下一項的話）返回True、失敗（沒有下一項可讀）則返回False
		var data Ip // 宣告一個 type 為 Ip struct 的變數
		err = rows.Scan(&data.IP, &data.Time, &data.Url, &data.Status)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Cannot Find Data")
			return
		}
		datas = append(datas, data)
	}
	c.IndentedJSON(http.StatusOK, datas)
	defer db.Close()
}

// find iplogs of a single ip
func GetDataByIP(c *gin.Context) {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		c.JSON(http.StatusBadGateway, "Connect DB failed")
		return
	}
	if err := db.Ping(); err != nil {
		c.JSON(http.StatusBadGateway, "Connect DB failed")
		return
	}
	fmt.Println("DB iplog connected")

	ip := c.Param("ip")
	fmt.Printf("iP = %v", ip)
	// Loop over the list of iplogs, looking for
	// a log whose Ip value matches the parameter.
	var datas []Ip
	sqlIp := "SELECT * FROM ip_log WHERE IP = '" + ip + "' ;"
	rows, err := db.Query(sqlIp)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Connot Find Datas")
		return
	}
	fmt.Printf("Ip = %v data found", ip)

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() { // rows.Next() 前往下一個項目。如果成功（還有下一項的話）返回True、失敗（沒有下一項可讀）則返回False
		var data Ip // 宣告一個 type 為 Ip struct 的變數
		err = rows.Scan(&data.IP, &data.Time, &data.Url, &data.Status)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Cannot Find Data")
			return
		}
		fmt.Println(data)
		datas = append(datas, data)
	}
	c.IndentedJSON(http.StatusOK, datas)

	defer rows.Close()
	defer db.Close()
}

