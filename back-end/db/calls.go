package db

import (
	"xorm.io/xorm"
)

func Create(engine *xorm.Engine, table interface{}) error {
	_, err := engine.Insert(table)
	return err
}

func Read(engine *xorm.Engine, table interface{}, condition interface{}) error {
	_, err := engine.Where(condition).Get(table)
	return err
}

func Update(engine *xorm.Engine, table interface{}, condition interface{}) error {
	_, err := engine.Where(condition).Update(table)
	return err
}

func Delete(engine *xorm.Engine, table interface{}, condition interface{}) error {
	_, err := engine.Where(condition).Delete(table)
	return err
}

func GetByID(engine *xorm.Engine, table interface{}, id int64) (bool, error) {
	found, err := engine.ID(id).Get(table)
	return found, err
}
func GetByStorageId(engine *xorm.Engine, table interface{}, storageID string) (bool, error) {
	found, err := engine.Where("storage = ?", storageID).Get(table)
	return found, err
}
func GetByEmail(engine *xorm.Engine, table interface{}, email string) (bool, error) {
	found, err := engine.Where("email = ?", email).Get(table)
	return found, err
}

func UpdateStorage(engine *xorm.Engine, storage string, newValues map[string]interface{}) error {
	_, err := engine.Table(new(Store)).Where("storage = ?", storage).Update(newValues)
	return err
}

func GetAll(engine *xorm.Engine, tableSlice interface{}) error {
	err := engine.Find(tableSlice)
	return err
}
