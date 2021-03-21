package main

import (
	"fmt"
	"os"
	"shukai-api/domain"
	"shukai-api/infrastructure"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("env load error")
	}

	dsn := os.Getenv("DATABASE_URL")
	sqlHandler := infrastructure.NewSqlHandler(dsn)

	sqlHandler.Conn.AutoMigrate(&domain.UserSkillModel{})
	sqlHandler.Conn.AutoMigrate(&domain.RecruitmentConditionsModel{})

	sqlHandler.Conn.AutoMigrate(&domain.UserModel{})
	sqlHandler.Conn.AutoMigrate(&domain.NotificationModel{})
	sqlHandler.Conn.AutoMigrate(&domain.RecruitmentModel{})
	sqlHandler.Conn.AutoMigrate(&domain.ApprovalWaitModel{})
	sqlHandler.Conn.AutoMigrate(&domain.RecruitmentUsersModel{})
}
