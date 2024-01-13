package model

import (
	"log"
	"time"
	"strconv"

)

type Dividend struct {
	ID           int
	Code         string
	DateCom      time.Time
	DatePag      time.Time
	Value        float64
	DividendType string
	DateAtu      string
}

func (dividend *Dividend) SaveDividend(dividends []Dividend) error {
	db := GetConnection()
	defer db.Close()

	for _, dividend := range dividends {

		// Check if the stock code already exists in the database
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM dividend WHERE hash = $1",
			dividend.createHashCode()).Scan(&count)
		if err != nil {
			log.Printf("Error checking if stock code exists: %s", err)
			return err
		}
		if count == 0 {
			// Update the existing stock
			_, err = db.Exec("insert into dividend (hash, stock_code, datecom, datepag, value, tipo) values ($1, $2, $3, $4, $5, $6)",
				dividend.createHashCode(), dividend.Code, removeTime(dividend.DateCom), removeTime(dividend.DatePag), dividend.Value, dividend.DividendType)
			if err != nil {
				log.Printf("Error updating stock: %s", err)
				return err
			}
		}
	}
	return nil
}



func (dividend *Dividend) createHashCode() (hash string) {
	return dividend.Code + removeTime(dividend.DateCom).String() + removeTime(dividend.DatePag).String() + strconv.FormatFloat(dividend.Value, 'f', -1, 64)
}


func removeTime(t time.Time) time.Time {
    return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

