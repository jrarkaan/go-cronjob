package playground

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// set scheduler berdasarkan zona waktu sesuai kebutuhan
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))

	// stop scheduler tepat sebelum fungsi berakhir
	defer scheduler.Stop()

	// set task yang akan dijalankan scheduler
	// gunakan crontab string untuk mengatur jadwal
	scheduler.AddFunc("*/1 * * * *", func() { SendAutomail("New Year") })
	scheduler.AddFunc("0 07 10 * *", SendMonthlyBillingAutomail)
	scheduler.AddFunc("0 09 * * 1-5", NotifyDailyAgenda)
	scheduler.AddFunc("*/10 * * * *", NotifyNewOrder)

	// start scheduler
	go scheduler.Start()

	// trap SIGINT untuk trigger shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
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
