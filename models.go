package main

import (
	"github.com/jinzhu/gorm"
	"log"
)

var (
	db, _ = gorm.Open("sqlite3", "gorm.db")
)

type Event struct {
	ID			uint		`json:"-" gorm:"primary_key"`
	Name 		string 		`json:"name,omitempty"`
	Event 		string 		`json:"event,omitempty" sql:"size:255;unique;index"`
	Handlers 	[]Handler	`gorm:"foreignkey:EventID"`
	Published 	bool 		`sql:"DEFAULT:false"`
	Data 		EventData	`gorm:"foreignkey:EventID"`
}

type Handler struct{
	EventID		int			`json:"-" gorm:"unique_index:hadlerIndex"`
	Address 	string 		`json:"address,omitempty" gorm:"unique_index:hadlerIndex"`

}
type EventData struct {
	EventID		int			`json:"-" gorm:"unique"`
	Key			string		`json:"key,omitempty"`
	Value		string		`json:"value,omitempty"`
}

func ModelsInit() {
	db.LogMode(true)

	if !db.HasTable(&Event{}) {
		db.CreateTable(&Event{},&Handler{}, &EventData{})
	}

	db.AutoMigrate(&Event{})
	db.AutoMigrate(&Handler{})
	db.AutoMigrate(&EventData{})
}

func Store(evt Event) {
	if db.NewRecord(evt) {
		if err := db.Create(&evt).Error; err != nil {
			log.Fatal("Error:", err)
		}
	}

}
func ModelRemove(name string) {
	var evt Event
	db.Where(&evt, "name = ?",name)
	log.Print(evt)
	db.Delete(&evt)
}

func ModelPublish(name string, data interface{}){
	var evt = Event{}
	db.Where(Event{Name: name}).Find(&evt)
	evt.Published = true
	evt.Data = EventData{}
	db.Save(&evt).Assign(evt)

}

