package main

import (
	"fmt"
	"golang/rest-api-presensi/helper"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/robfig/cron"
)

func TestEpoch(t *testing.T) {
	waktu := time.Now().UnixNano() / 1000000
	fmt.Println(waktu)
	fmt.Println(reflect.TypeOf(waktu))
}

func TestConvertEpoch(t *testing.T) {
	time := time.Unix(-2211750432000, 0)
	fmt.Println(time)
}

func TestCron(t *testing.T) {
	epochMillis := int64(1674482340000)

	// Convert milliseconds to time
	scheduledTime := time.Unix(0, epochMillis*int64(time.Millisecond))
	fmt.Println(scheduledTime)
	fmt.Println(scheduledTime.Format("15:04"))

	c := cron.New()
	c.AddFunc("0 0 "+scheduledTime.Format("15:04")+" 23 1 * 2023", func() {
		_, err := http.Get("https://api.telegram.org/bot5836520934:AAGQ_iQvY7Hbm5goVRPbwH-k57p25dA-Gns/SendMessage?chat_id=1359729975&text=Halorendrainitescron")
		if err != nil {
			panic(err)
		}
	})
	c.Start()
	select {}
}

func TestDataType(t *testing.T) {
	OtpTime := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println("tipe data :", reflect.TypeOf(OtpTime))
}

func TestParsing(t *testing.T) {
	jam := "7.45"
	parts := strings.Split(jam, ".")
	hour := parts[0]
	minute := parts[1]

	fmt.Println("hour:", hour)
	fmt.Println("minute:", minute)
}

func TestTimeNow(t *testing.T) {
	// now := time.Now()
	// nowhour := now.Hour()
	// nowminute := now.Minute()

	// fmt.Println(nowhour)
	// fmt.Println(nowminute)

	loc, err := time.LoadLocation("Asia/Jakarta") // mengatur zona waktu ke WIB
	if err != nil {
		fmt.Println(err)
		return
	}

	nowTime := time.Now().In(loc) // mengambil waktu saat ini dalam zona waktu WIB

	fmt.Println("Waktu saat ini di WIB:", nowTime)
}
func TestUploadImages(t *testing.T) {
	path, err := helper.UploadImageProfile()
	helper.PanicIfError(err)
	fmt.Println(path)

}
