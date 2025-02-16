package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	goLib_db "github.com/quack337/goLib/db"
	goLib_sqlite "github.com/quack337/goLib/db/sqlite"
)

var db *sql.DB
var port int = 8282
var tables []string

func parseArgs() {
	portSet := false
	dbSet := false
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-p" || os.Args[i] == "-P" {
			i++
			p, err := strconv.Atoi(os.Args[i])
			if err != nil {				
				panic("port 번호 오류. 사용법: goRest <db file> [-p <port number>]")
			} else {
				port = p
				portSet = true
				if dbSet { break }
			}	
		} else {
			_db, err := sql.Open("sqlite3", "file:" + os.Args[i] + "?_foreign_keys=on")
			if err != nil {
				panic("db 파일 열기 에러: " + os.Args[i] + " " + err.Error())
			} else {
				db = _db
				_table, err := goLib_sqlite.GetTableNames(db)
				if err != nil {
					panic("테이블명 조회 에러: " + err.Error())
				} else {
					tables = _table
				}
				dbSet = true
				if portSet { break }
			}
		}
	}
	if !dbSet {
		panic("db 파일 오류. 사용법: goRest <db file> [-p <port number>]")
	}
}

func main() {
	parseArgs()
	fmt.Println("서버 시작됨. port 번호:", port)

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	router.GET("/api/:table", func(ctx *gin.Context) {
		table := ctx.Param("table")
		if _, found := slices.BinarySearch(tables, table); !found {
			ctx.JSON(http.StatusNotFound, gin.H{"error": table + " 테이블을 찾을 수 없습니다."})
		} else {
			if rows, err := db.Query("SELECT * FROM " + table); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "데이터 조회 에러: " + err.Error()})
			} else {
				if json, err := goLib_db.Jsonify(rows); err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": "데이터 조회 에러: " + err.Error()})
				} else {
					ctx.String(http.StatusOK, json)
				}
			}
		}
	})

	router.Run(":" + strconv.Itoa(port))
}