package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Antonio-Gabriel/go_repository_pattern/repository"
	serviceInstance "github.com/Antonio-Gabriel/go_repository_pattern/service"
)

func main() {

	/// In memory repository

	// db := repository.CategoriesMemoryDb{Categories: []entity.Category{}}
	// repositoryMemory := repository.NewCategoryRepositoryMemory(db)

	/// Sqlite repository

	db, _ := sql.Open("sqlite3", "./sqlite.repository")
	sqliteRepository := repository.NewCategorySqlite(db)

	service := serviceInstance.NewCategoryService(sqliteRepository)
	category, _ := service.Create("Pc Laptop")

	transformToJson, _ := json.Marshal(category)

	fmt.Println(string(transformToJson))
}
