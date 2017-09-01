package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	bittrex "github.com/toorop/go-bittrex"
)

const (
	BITTREX_API_KEY    = "644a03c302374523869ba9e87421b3ae"
	BITTREX_API_SECRET = "MU72EQ6OTBDVGY65"
)

var bittrexapi = bittrex.New(BITTREX_API_KEY, BITTREX_API_SECRET)

var Db *sql.DB

func ServiceImplInit(config Config) {
	if config.DbConnectionString == "" {
		panic("env not defined=DbConnectionString")
	}
	db, err := sql.Open("mysql", config.DbConnectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	Db = db
}

type MarketServiceImpl struct{}

func (mrs MarketServiceImpl) RefreshAllTicks(eid int) error {

	if eid == 1 {
		var mCode string
		q1 := "SELECT m_code FROM market_list WHERE eid=?"
		stmt, err := Db.Prepare(q1)
		if err != nil {
			panic(err)
			return err
		}
		rows, err := stmt.Query(eid)

		for rows.Next() {
			rows.Scan(&mCode)
			fmt.Println("scan:" + mCode)
			btrxticker, err := bittrexapi.GetTicker(mCode)
			if err != nil {
				return errors.New("error in fetching tick:" + mCode + ", in eid:" + string(eid))
			}
			var ticker Ticker
			ticker.Bid = btrxticker.Bid
			ticker.Ask = btrxticker.Ask
			ticker.Last = btrxticker.Last
			ticker.ServerTime = 0

			mrs.RefreshTick(eid, mCode, ticker)
		}

		defer rows.Close()
		return nil
	} else if eid == 2 { // vipbitcoin
		// var vipbtc = msvc.GetVIPTicker("btc", "idr")
		mrs.RefreshTick(2, "btc_idr", GetVIPTicker("btc", "idr"))
		mrs.RefreshTick(2, "eth_idr", GetVIPTicker("eth", "idr"))
		return nil
		// fmt.Printf("VIP=?", vipticker)
	} else if eid == 3 { // luno
		// var vipbtc = msvc.GetVIPTicker("btc", "idr")
		ticker, err := GetLunoTicker("btc", "idr")
		if err != nil {
			log.Println(err)
			return err
		}
		mrs.RefreshTick(3, "XBTIDR", ticker)
		return nil
	} else {
		return errors.New("undefined exchange_id")
	}
}

func (mrs MarketServiceImpl) RefreshTick(eid int, mCode string, ticker Ticker) error {
	// Insert into history
	q := "INSERT INTO tick_history (SELECT * FROM current_tick WHERE eid=? AND m_code=?)"
	stmt, err := Db.Prepare(q)
	if err != nil {
		// panic(err)
		log.Fatal(err)
		return err
	}
	_, err = stmt.Exec(eid, mCode)
	if err != nil {
		// panic(err)
		log.Fatal(err)
		return err
	}

	// Insert new val
	q2 := "INSERT INTO current_tick (eid, m_code, last, ask, bid, server_time, ts) VALUES (?, ?, ?, ?, ?, ?, NOW()) ON DUPLICATE KEY UPDATE last=?, ask=?, bid=?, server_time=?, ts=NOW()"
	stmt2, err2 := Db.Prepare(q2)
	if err2 != nil {
		// panic(err)
		log.Fatal(err)
		return err2
	}
	fmt.Println("waktu server=?", ticker.ServerTime)
	_, err2 = stmt2.Exec(eid, mCode, ticker.Last, ticker.Ask, ticker.Bid, ticker.ServerTime, ticker.Last, ticker.Ask, ticker.Bid, ticker.ServerTime)
	if err2 != nil {
		// panic(err2)
		log.Fatal(err)
		return err2
	}

	defer stmt.Close()

	return nil
}

func (mrs MarketServiceImpl) GetCurrentTick(eid int, mCode string) (ticker Ticker, err error) {
	var ret Ticker
	q1 := "SELECT bid, last, ask FROM current_tick WHERE eid=? and m_code=?"
	stmt, err := Db.Prepare(q1)
	if err != nil {
		return ret, err
	}
	row := stmt.QueryRow(eid, mCode)
	row.Scan(&ret.Bid, &ret.Last, &ret.Ask)
	return ret, nil
}

func (mrs MarketServiceImpl) GetCurrentTick2(eid int, c1 string, c2 string) (ticker Ticker, err error) {
	var ret Ticker
	q1 := "SELECT bid, last, ask FROM current_tick A INNER JOIN market_list B on A.eid=B.eid AND A.m_code=B.m_code WHERE B.eid=? and B.c1_code=? AND B.c2_code=?"
	stmt, err := Db.Prepare(q1)
	if err != nil {
		return ret, err
	}
	row := stmt.QueryRow(eid, c1, c2)
	row.Scan(&ret.Bid, &ret.Last, &ret.Ask)
	return ret, nil
}

//func (mrs MarketServiceImpl) RefreshTick(eid int, mCode string, ticker Ticker) error {
