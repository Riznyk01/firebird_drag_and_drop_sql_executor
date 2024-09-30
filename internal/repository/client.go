package repository

import (
	"database/sql"
	"log"
	"time"
)

type FirebirdClient struct {
	db *sql.DB
}

func NewFirebirdClient(db *sql.DB) *FirebirdClient {
	return &FirebirdClient{
		db: db,
	}
}

func (c *FirebirdClient) UpdateDBCorrectionDate(currentTime time.Time) error {
	fc := "UpdateDBCorrectionDate"

	query := "UPDATE V_PARAM_VALUES SET F_DATE_VAL = ? WHERE F_PARAM_NAME = ?"

	_, err := c.db.Exec(query, currentTime, "GBackDate")
	if err != nil {
		log.Printf("%s: %v\n", fc, err)
		return err
	}
	return nil
}
func (c *FirebirdClient) GetDBCorrectionDate() (time.Time, error) {
	fc := "GetDBCorrectionDate"

	query := "SELECT F_DATE_VAL FROM V_PARAM_VALUES WHERE F_PARAM_NAME = ?"

	var dateVal time.Time
	err := c.db.QueryRow(query, "GBackDate").Scan(&dateVal)
	if err != nil {
		log.Printf("%s: %v\n", fc, err)
		return time.Time{}, err
	}
	return dateVal, nil
}
