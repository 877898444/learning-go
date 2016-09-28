package main

import(
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  // _ "github.com/go-sql-driver/mysql"
  "fmt"
)

func main() {
  db, err := gorm.Open("mysql", "root:@/say_morning_development?charset=utf8&parseTime=True&loc=Local")
  if err != nil{
    panic(err)
  }
  fmt.Println(db)
  defer db.Close()

  // db.AutoMigrate(&User{})
  // db.AutoMigrate(&Uer{}, &Order{})
  //
  // db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
  //
  // db.HasTable(&User{})

  // todos table
  type Todo struct {
    gorm.Model
    ID int
    Title string
    Content string
  }
  db.Model(&Todo{}).AddIndex("idx_title", "title")

}
// Auto Migration
//
// Automatically migrate your schema, to keep your schema update to date.
//
// WARNING: AutoMigrate will ONLY create tables, missing columns and missing indexes, and WON'T change existing column's type or delete unused columns to protect your data.
//
// db.AutoMigrate(&User{})
//
// db.AutoMigrate(&User{}, &Product{}, &Order{})
//
// // Add table suffix when create tables
// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
// Has Table
//
// // Check model `User`'s table exists or not
// db.HasTable(&User{})
//
// // Check table `users` exists or not
// db.HasTable("users")
// Create Table
//
// // Create table for model `User`
// db.CreateTable(&User{})
//
// // will append "ENGINE=InnoDB" to the SQL statement when creating table `users`
// db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
// Drop table
//
// // Drop model `User`'s table
// db.DropTable(&User{})
//
// // Drop table `users
// db.DropTable("users")
//
// // Drop model's `User`'s table and table `products`
// db.DropTableIfExists(&User{}, "products")
// ModifyColumn
//
// Modify column's type to given value
//
// // change column description's data type to `text` for model `User`
// db.Model(&User{}).ModifyColumn("description", "text")
// DropColumn
//
// // Drop column description from model `User`
// db.Model(&User{}).DropColumn("description")
// Add Foreign Key
//
// // Add foreign key
// // 1st param : foreignkey field
// // 2nd param : destination table(id)
// // 3rd param : ONDELETE
// // 4th param : ONUPDATE
// db.Model(&User{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")
// Indexes
//
// // Add index for columns `name` with given name `idx_user_name`
// db.Model(&User{}).AddIndex("idx_user_name", "name")
//
// // Add index for columns `name`, `age` with given name `idx_user_name_age`
// db.Model(&User{}).AddIndex("idx_user_name_age", "name", "age")
//
// // Add unique index
// db.Model(&User{}).AddUniqueIndex("idx_user_name", "name")
//
// // Add unique index for multiple columns
// db.Model(&User{}).AddUniqueIndex("idx_user_name_age", "name", "age")
//
// // Remove index
// db.Model(&User{}).RemoveIndex("idx_user_name")
