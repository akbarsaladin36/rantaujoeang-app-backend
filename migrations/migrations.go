package migrations

import (
	"fmt"
	"rantaujoeang-app-backend/database"
	"rantaujoeang-app-backend/models"
)

func MigrateTables() {
	migrateUser()
	migrateBalanceTransactions()
	migrateDorm()
	migratePayments()
	migrateInvoices()
	migrateSessions()
	migrateMessages()
}

func migrateUser() {

	// database.DB.Migrator().DropTable(models.User{})

	if !database.DB.Migrator().HasTable(models.User{}) {
		err := database.DB.AutoMigrate(models.User{})
		if err != nil {
			fmt.Println("Error migrating table", err)
			return
		}
		fmt.Printf("Table %s is succesfully migrated!\n", "User")
	}
}

func migrateBalanceTransactions() {

	// database.DB.Migrator().DropTable(models.BalanceTransaction{})

	if !database.DB.Migrator().HasTable(models.BalanceTransaction{}) {
		err := database.DB.AutoMigrate(models.BalanceTransaction{})
		if err != nil {
			fmt.Println("Error migrating table", err)
			return
		}
		fmt.Printf("Table %s is succesfully migrated!\n", "Balance Transaction")
	}

	if !database.DB.Migrator().HasColumn(models.BalanceTransaction{}, "BalanceTransactionProcessCd") {
		err := database.DB.Migrator().AddColumn(models.BalanceTransaction{}, "BalanceTransactionProcessCd")
		if err != nil {
			fmt.Println("Error adding column", err)
			return
		}
		fmt.Printf("Column %s is succesfully added!\n", "Balance_Transaction_Process_Cd")
	}
}

func migrateDorm() {
	if !database.DB.Migrator().HasTable(models.Dorm{}) {
		err := database.DB.AutoMigrate(models.Dorm{})
		if err != nil {
			fmt.Println("Error migrating table", err)
			return
		}
		fmt.Printf("Table %s is succesfully migrated!\n", "Dorm")
	}
}

func migratePayments() {

	// database.DB.Migrator().DropTable(models.Payment{})

	if !database.DB.Migrator().HasTable(models.Payment{}) {
		err := database.DB.AutoMigrate(models.Payment{})
		if err != nil {
			fmt.Println("Error migrating table", err)
			return
		}
		fmt.Printf("Table %s is succesfully migrated!\n", "Payment")
	}
}

func migrateInvoices() {

	// database.DB.Migrator().DropTable(models.Invoice{})

	if !database.DB.Migrator().HasTable(models.Invoice{}) {
		err := database.DB.AutoMigrate(models.Invoice{})
		if err != nil {
			fmt.Println("Error migrating table", err)
			return
		}
		fmt.Printf("Table %s is succesfully migrated!\n", "Invoice")
	}
}

func migrateSessions() {

	// database.DB.Migrator().DropTable(models.Session{})

	if !database.DB.Migrator().HasTable(models.Session{}) {
		err := database.DB.AutoMigrate(models.Session{})
		if err != nil {
			fmt.Println("Error migrating table", err)
			return
		}
		fmt.Printf("Table %s is succesfully migrated!\n", "Session")
	}
}

func migrateMessages() {

	// database.DB.Migrator().DropTable(models.Message{})

	if !database.DB.Migrator().HasTable(models.Message{}) {
		err := database.DB.AutoMigrate(models.Message{})
		if err != nil {
			fmt.Println("Error migrating table", err)
			return
		}
		fmt.Printf("Table %s is succesfully migrated!\n", "Message")
	}
}
