package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/zhengzhou1992/cmdtools/internal/sql2struct"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql tools",
	Long:  "sql tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql to struct",
	Long:  "sql to struct converter",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}

		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "DB username")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "DB password")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "DB host")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "DB charset")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "DB type (mysql, postgress, etc...")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "DB name")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "DB table name")
}
