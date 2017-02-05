package gruff

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"time"
)

type Model struct {
	ID        uint64     `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"-" sql:"DEFAULT:current_timestamp"`
	UpdatedAt time.Time  `json:"-" sql:"DEFAULT:current_timestamp"`
	DeletedAt *time.Time `json:"-" settable:"false"`
}

func ModelToJson(model interface{}) string {
	j, err := json.Marshal(model)
	if err != nil {
		panic(fmt.Sprintf("Error %v encoding JSON for %v", err, model))
	}

	jsonStr := string(j)
	v := reflect.Indirect(reflect.ValueOf(model))
	ot := v.Type()
	t := ot
	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		t = t.Elem()
	} else if t.Kind() == reflect.Interface {
		t = v.Elem().Type()
	}
	return jsonStr
}

func ModelToJsonMap(modl interface{}) map[string]interface{} {
	jsonStr := ModelToJson(modl)
	m := JsonToMap(jsonStr)
	return m
}

func JsonToMap(jsonStr string) map[string]interface{} {
	jsonMap := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(fmt.Sprintf("Error %v unmarshaling JSON for %v", err, jsonStr))
	}

	return jsonMap
}

func JsonToMapArray(jsonStr string) []map[string]interface{} {
	var arr []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &arr)

	if err != nil {
		panic(fmt.Sprintf("Error %v unmarshaling JSON for %v", err, jsonStr))
	}

	return arr
}

func JsonToModel(jsonStr string, item interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), &item)

	if err == nil {
		v := reflect.Indirect(reflect.ValueOf(item))
		ot := v.Type()
		t := ot
		if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
			t = t.Elem()
		} else if t.Kind() == reflect.Interface {
			t = v.Elem().Type()
		}
	}
	return err
}

type ServerContext struct {
	Database *gorm.DB
	Test     bool
}
