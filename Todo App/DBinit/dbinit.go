package DBinit

import(
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

//table latout
type Inputlog struct {
	ID uint64 `gorm:"primaryKey;autoIncrement:false;unique"`
	Log  string
}

//connect to DB
func Connect()(*gorm.DB , error){
	pswd := os.Getenv("MYSQL_PASSWORD")
	dsn := "root:"+pswd+"@tcp(localhost:3306)/todo_db"
	db , err := gorm.Open(mysql.Open(dsn) , &gorm.Config{})
	if err != nil{
		return nil , err
	}else{
		return db , nil
	}
}


//insert a value
func Insert(id uint64 , l string){
	db , err := Connect()
	if err != nil{
		fmt.Printf("Error while trying to connect to DB: %v" , err)
	}
	db.AutoMigrate(&Inputlog{})
	db.Create(&Inputlog{ID: id , Log: l})
}


//delete a value
func Delete(id uint64){
	db , err := Connect()
	if err != nil{
		fmt.Printf("Error while trying to connect to DB: %v" , err)
	}
	db.Delete(&Inputlog{}, id)
	fmt.Println("Deleted the Log with the given ID")
}


//read all of the values
func Read_all(){
	db , err := Connect()
	if err != nil{
		fmt.Printf("Error while trying to connect to DB: %v" , err)
	}
	var logs []Inputlog
	db.Find(&logs)
	for _, log := range logs {
		if log.ID < 10{
			fmt.Printf("ID: %d        Log: %s\n", log.ID, log.Log)
		}else if log.ID < 100{
			fmt.Printf("ID: %d       Log: %s\n", log.ID, log.Log)
		}else{
			fmt.Printf("ID: %d      Log: %s\n", log.ID, log.Log)
		}
	}
}


//read a value
func Read(id uint64){
	db , err := Connect()
	if err != nil{
		fmt.Printf("Error while trying to connect to DB: %v" , err)
	}
	var logs []Inputlog
	//grabing a value with the desired id
	db.First(&logs, id)
	for _, log := range logs {
		fmt.Printf("ID: %d        Log: %s\n", log.ID, log.Log)
	}
}

//update a value
func Update(id uint64 , l string){
	db , err := Connect()
	if err != nil{
		fmt.Printf("Error while trying to connect to DB: %v" , err)
	}

	var logs Inputlog
	//grab an existing value and changing it
	db.First(&logs, id)
	logs.Log  = l
	db.Save(&logs)
}