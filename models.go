package main

import (
	"github.com/jinzhu/gorm"
)

var (
	db, _ = gorm.Open("sqlite3", "gorm.db")
)

type Event struct {
	ID 			int			`json:"-" gorm:"primary_key";"AUTO_INCREMENT"`
	Name 		string 		`json:"name,omitempty" gorm:"primary_key"`
	Event 		string 		`json:"event,omitempty" sql:"size:255;unique;index"`
	Handlers 	[]Handler	`gorm:"foreignkey:EventID"`
	Published 	bool 		`sql:"DEFAULT:false"`
	data 		interface{}
}

type Handler struct{
	ID 			int			`json:"-" gorm:"primary_key";"AUTO_INCREMENT"`
	EventID		int			`json:"-" gorm:"unique_key"`
	Address 	string 		`json:"address,omitempty" sql:"unique;index"`

}

func ModelsInit() {
	db.LogMode(true)

	if !db.HasTable(&Event{}) {
		db.CreateTable(&Event{},&Handler{})
	}

	db.AutoMigrate(&Event{})
	db.AutoMigrate(&Handler{})
}

func Store(evt Event) {
	db.Where(Event{Name: evt.Name}).Assign(evt).FirstOrCreate(&evt)
	//if db.NewRecord(evt) {
	//	if err := db.Create(&evt).Error; err != nil {
	//		log.Fatal("Error:", err)
	//	}
	//}

}
func ModelRemove(name string) {
	db.Where(Event{Name: name}).Association("Handlers").Delete(&Event{Name: name})
	db.Where(Event{Name: name}).Delete(&Event{Handlers:[]Handler{}})
}

func ModelPublish(name string, data interface{}){
	var evt = Event{}
	db.Where(Event{Name: name}).Find(&evt)
	evt.Published = true
	evt.data = data
	db.Save(&evt).Assign(evt)

}

