package main

import (
	"fmt"
	"github.com/jrarkaan/go-cronjob/app/database"
	"time"
)

func main() {
	db, _ := database.Connect()
	//_ := config.Init()
	// migration table on database raka raka
	database.Migrate(db)
	// run cron
	//cron.InitCroService(db).CronProduct()
}

func SendAutomail(automailType string) {
	// ... instruksi untuk mengirim automail berdasarkan automailType
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " SendAutomail " + automailType + " telah dijalankan.\n")
}

func SendMonthlyBillingAutomail() {
	// ... instruksi untuk mengirim automail berisi tagihan
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " SendMonthlyBillingAutomail telah dijalankan.\n")
}

func NotifyDailyAgenda() {
	// ... instruksi untuk mengirim notifikasi agenda harian
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " NotifyDailyAgenda telah dijalankan.\n")
}

func NotifyNewOrder() {
	// ... instruksi untuk mengecek pesanan baru, lalu mengirimkan notifikasi
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05") + " NotifyNewOrder telah dijalankan.\n")
}
