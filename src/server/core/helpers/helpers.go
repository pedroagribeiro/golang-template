package helpers

import (
	"fmt"
	"template/core/db"
)

type DAO[DataType any] interface {
	Insert(DataType) (DataType, int, error)
	Get(DataType) (DataType, bool, error)
	Find(DataType) ([]DataType, int, error)
	Update(DataType) (DataType, int, error)
	Remove(DataType) (int, error)
	Save(DataType) error
	GetNextResult(DataType) ([]DataType, error)
}

func PreloadEntry[DataType any](session db.IDbSession, ids DataType, additionalPreloads ...string) (DataType, error) {
	results := []DataType{}
	if err := session.Preload(&results, &ids, additionalPreloads...); err != nil {
		return ids, err
	}
	return results[0], nil
}

func GetEntry[DataType any](session db.IDbSession, ids DataType) (DataType, bool, error) {
	results := []DataType{}
	if err := session.FindOne(&results, &ids); err != nil {
		return ids, false, err
	} else if len(results) == 0 {
		return ids, false, nil
	}
	return results[0], true, nil
}

func GetEntryWithPreloads[DataType any](session db.IDbSession, ids DataType, preloads ...string) (DataType, bool, error) {
	results := []DataType{}
	if err := session.FindOneWithPreload(&results, &ids, preloads...); err != nil {
		return ids, false, err
	} else if len(results) == 0 {
		return ids, false, nil
	}
	return results[0], true, nil
}

func ExistsEntry[DataType any](session db.IDbSession, ids DataType) (bool, error) {
	results := []DataType{}
	if err := session.FindOne(&results, &ids); err != nil {
		return false, err
	} else if len(results) == 0 {
		return false, nil
	}
	return true, nil
}

func FindEntry[DataType any](session db.IDbSession, ids DataType) ([]DataType, error) {
	results := []DataType{}
	if err := session.Find(&results, &ids); err != nil {
		return results, err
	} else {
		return results, nil
	}
}

func FindEntryWithPreloads[DataType any](session db.IDbSession, ids DataType, preloads ...string) ([]DataType, error) {
	results := []DataType{}
	if err := session.FindWithPreload(&results, &ids, preloads...); err != nil {
		return results, err
	} else {
		return results, nil
	}
}

func DeleteEntry[DataType any](session db.IDbSession, ids DataType) (DataType, error) {
	var placeholder DataType
	if err := session.Delete(&placeholder, &ids); err != nil {
		return ids, err
	} else {
		ids = placeholder
		return ids, nil
	}
}

func ReplaceEntry[DataType any](session db.IDbSession, newData DataType, ids ...any) (DataType, error) {
	if err := session.Update(&newData); err != nil {
		return newData, err
	} else {
		return newData, nil
	}
}

func CreateEntry[DataType any](session db.IDbSession, data DataType) (DataType, error) {
	if err := session.CreateWithPreload(&data); err != nil {
		return data, err
	} else {
		return data, nil
	}
}

func CreateInBatch[DataType any](session db.IDbSession, data DataType, batchSize int) (DataType, error) {
	if err := session.CreateInBatch(data, batchSize); err != nil {
		return data, err
	} else {
		return data, nil
	}
}

func SaveEntry[DataType any](session db.IDbSession, newData DataType) error {
	return session.Save(&newData)
}

func GetAllResult[DataType any](session db.IDbSession, values map[string]interface{}, data DataType) ([]DataType, error) {
	var service DataType
	results := []DataType{}
	res, err := session.QueryRows(values, &data)
	if err != nil {
		return results, err
	}

	defer res.Close()
	// Iterate over the rows and scan into the struct
	for res.Next(&service) {
		results = append(results, service)
	}

	return results, nil
}

func EntryNotFound(tableName string, constraints any) error {
	return fmt.Errorf("could not find %s entry with constraints: %+v", tableName, constraints)
}
