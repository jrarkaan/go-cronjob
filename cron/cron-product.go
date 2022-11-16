package cron

import (
	"encoding/json"
	"fmt"
	"github.com/jrarkaan/go-cronjob/dto"
	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type ICronProductService interface {
	CronProduct()
}

type NewCronProductService struct {
	conn       gorm.DB
	LengthData int
}

func InitCroService(conn *gorm.DB) ICronProductService {
	return &NewCronProductService{
		conn:       *conn,
		LengthData: 0,
	}
}
func (service *NewCronProductService) getProductMaster() <-chan dto.MProduct {
	var result dto.MProductRequest
	chanOut := make(chan dto.MProduct)
	url := "https://api.cfu.pharmalink.id/integra/v2/products"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error to request HTTP Product: %s", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error to read body request: %s", err)
	}
	errU := json.Unmarshal([]byte(body), &result)
	if errU != nil {
		log.Printf("Error unmarshal: %s", errU)
	}
	go func() {
		for _, value := range result.Data {
			chanOut <- dto.MProduct{
				ProCode:            value.ProCode,
				Sku:                value.Sku,
				ProName:            strconv.Quote(value.ProName),
				ProName2:           strconv.Quote(value.ProName2),
				GenerikName:        value.GenerikName,
				Kategori:           value.Kategori,
				Etalase:            value.Etalase,
				FotoUtama:          value.FotoUtama,
				FotoDepan:          value.FotoDepan,
				FotoSamping:        value.FotoSamping,
				FotoAtas:           value.FotoAtas,
				FotoDetail:         value.FotoDetail,
				Deskripsi:          value.Deskripsi,
				LinkVideo:          value.LinkVideo,
				Custom:             value.Custom,
				SellUnit:           value.SellUnit,
				SellPack:           value.SellPack,
				SellPackName:       value.SellPackName,
				CustomSellUnit:     value.CustomSellUnit,
				CustomSellPack:     value.CustomSellPack,
				CustomSellPackName: value.CustomSellPackName,
				Berat:              value.Berat,
				HargaMin:           value.HargaMin,
				HargaMax:           value.HargaMax,
				HargaPdn:           value.HargaMin,
				HargaIndoApotek:    value.HargaIndoApotek,
			}
		}
		close(chanOut)
	}()
	return chanOut
}

func (service *NewCronProductService) getProductMaster2() {
	var result dto.MProductRequest
	url := "https://api.cfu.pharmalink.id/integra/v2/products"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error to request HTTP Product: %s", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error to read body request: %s", err)
	}
	errU := json.Unmarshal([]byte(body), &result)
	if errU != nil {
		log.Printf("Error unmarshal: %s", errU)
	}
	go func() {
		//for _, value := range result.Data {
		//	fmt.Printf("Harga Max: %d, Harga Min: %d", value.HargaMax, value.HargaMax)
		//}
		for _, value := range result.Data {
			fmt.Println("Harga Max: ", value.HargaMax)
			//fmt.Printf("Harga Max: %d, Harga Min: %d", value.HargaMax, value.HargaMax)
		}
	}()

}

func (service *NewCronProductService) insertProductMaster(chanIn <-chan dto.MProduct) <-chan dto.MProduct {
	chanOut := make(chan dto.MProduct)
	go func() {
		for productMaster := range chanIn {
			//fmt.Printf("Harga Max: %d, Harga Min: %d", productMaster.HargaMax, productMaster.HargaMax)
			if err := service.conn.Create(&productMaster).Error; err != nil {
				fmt.Println("GORM ERROR CREATE: ", err)
				panic(err)
			}
			service.conn.Preload("m_product").Find(&productMaster)
			chanOut <- productMaster
		}
		close(chanOut)
	}()
	return chanOut
}

func (service *NewCronProductService) mergeChanInsertBulk(chanInMany ...<-chan dto.MProduct) <-chan dto.MProduct {
	wg := new(sync.WaitGroup)
	chanOut := make(chan dto.MProduct)

	wg.Add(len(chanInMany))
	for _, eachChan := range chanInMany {
		go func(eachChan <-chan dto.MProduct) {
			for eachChanData := range eachChan {
				chanOut <- eachChanData
			}
			wg.Done()
		}(eachChan)
	}

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}

func (service *NewCronProductService) CronProduct() {
	// set scheduler berdasarkan zona waktu sesuai kebutuhan
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))
	// stop scheduler tepat sebelum fungsi berakhir
	defer scheduler.Stop()
	scheduler.AddFunc("*/1 * * * *", func() {
		totalCpu := runtime.NumCPU()
		fmt.Println("Total CPU: ", totalCpu)

		totalThread := runtime.GOMAXPROCS(-1)
		fmt.Println("Total Threads: ", totalThread)
		log.Println("start")
		log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~ start program with concurrent/ worker threads ~~~~~~~~~~~~~~~~~~~~~~~~~~")
		start := time.Now()
		//service.getProductMaster2()
		resultOfProduct := service.getProductMaster()
		insertBulk1 := service.insertProductMaster(resultOfProduct)
		insertBulk2 := service.insertProductMaster(resultOfProduct)
		insertBulk3 := service.insertProductMaster(resultOfProduct)
		insertBulk4 := service.insertProductMaster(resultOfProduct)
		insertBulk5 := service.insertProductMaster(resultOfProduct)
		bulkInsertChannel := service.mergeChanInsertBulk(insertBulk1, insertBulk2, insertBulk3, insertBulk4, insertBulk5)

		// print output
		counterTotal := 0
		for productInfo := range bulkInsertChannel {
			if productInfo.ProCode != "" {
				counterTotal++
			}
		}
		log.Printf("%d data to be insert", counterTotal)

		duration := time.Since(start)
		log.Println(" ~~~~~~~~~~~~~~~~~~~~~~~~~~ done in", duration.Seconds(), "seconds ~~~~~~~~~~~~~~~~~~~~~~~~~~")
	})

	// start scheduler
	go scheduler.Start()

	// trap SIGINT untuk trigger shutdown.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
