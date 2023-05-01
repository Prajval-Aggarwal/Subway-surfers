package db

import (
	"subway/server/response"

	"gorm.io/gorm"
)

var db *gorm.DB

func Transfer(connection *gorm.DB) {
	db = connection
}

func CreateRecord(data interface{}) error {

	err := db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func FindById(data interface{}, id interface{}, columName string) error {
	column := columName + "=?"
	err := db.Where(column, id).First(data).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateRecord(data interface{}, id interface{}, columName string) *gorm.DB {
	column := columName + "=?"
	result := db.Where(column, id).Updates(data)

	return result
}

func RawQuery(query string, data interface{}, args ...interface{}) error {

	err := db.Raw(query, args...).Scan(data).Error
	if err != nil {
		return err
	}

	// return nil if there were no errors
	return nil
}

func ExecuteQuery(query string, args ...interface{}) error {
	err := db.Exec(query, args...).Error
	if err != nil {
		return err
	}
	return nil
}
func DeleteRecord(data interface{}, id interface{}, columName string) error {
	column := columName + "=?"
	result := db.Where(column, id).Delete(data)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func RecordExist(tableName string, columnName string, value string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT * FROM " + tableName + " WHERE " + columnName + " = '" + value + "')"
	db.Raw(query).Scan(&exists)
	return exists
}

// rename it
func Fun(query string, args ...interface{}) response.PlayerDetails {
	playerDetails := &response.PlayerDetails{}
	row := db.Raw(query, args...).Row()

	row.Scan(&playerDetails.P_ID, &playerDetails.P_Name, &playerDetails.Email, &playerDetails.HighScore, &playerDetails.TotalDistance, &playerDetails.Coins)
	return *playerDetails

}
