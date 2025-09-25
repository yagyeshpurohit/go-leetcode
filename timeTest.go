package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"google.golang.org/protobuf/types/known/timestamppb"
	"math"
	"time"
)

func main() {
	//startTime := time.Unix(0, 1749407400000*1000000)
	//ledgerStartTime := startTime.AddDate(0, 0, -14).UnixNano() / 1000000
	//println(ledgerStartTime)
	//actionTime1 := time.Now().Unix() * 1000
	//actionTime2 := time.Now().UnixNano() / 1000000
	//fmt.Printf("a1:%d a2:%d", actionTime1, actionTime2)

	//amountWithoutTDS := 137.58
	//tdsPercentage := float64(8)
	//tdsAmount := (amountWithoutTDS * tdsPercentage) / 100
	//netAmount := amountWithoutTDS - tdsAmount
	//netAmount = RoundFloat(netAmount, 2)
	//println(netAmount)

	//signature := makeSignature("8005807652")
	//fmt.Println(signature)

	//lastRunTime := GetLastRunAt()
	//fmt.Println("Last Run At:", lastRunTime)
	milliTime := MakeTimeFromMillis(1755455403000).Format("02-Jan-2006")
	println(milliTime)

}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func makeSignature(data string) string {
	secret := "37a31be743d24f2ac2b0a6103604b825813df50797fd507e41dd77241336368c"
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func GetLastRunAt() int64 {
	var lastRunAt int64
	currentTime := timestamppb.Now().Seconds * 1000
	prevDayTime := currentTime - 86400000
	lastRunAt = GetDayStartTimeForDate(prevDayTime)
	return lastRunAt
}

func GetDayStartTimeForDate(timeInMillis int64) int64 {
	year, month, day := MakeTimeFromMillis(timeInMillis).Date()
	return MakeTimeToMillis(time.Date(year, month, day, 0, 0, 0, 0, time.Local))
}
func MakeTimeFromMillis(millis int64) time.Time {
	return time.Unix(0, millis*int64(time.Millisecond))
}

func MakeTimeToMillis(time time.Time) int64 {
	return MakeTimestampInMillis(timestamppb.New(time))
}

func MakeTimestampInMillis(passedTime *timestamppb.Timestamp) int64 {
	if passedTime == nil {
		passedTime = timestamppb.Now()
	}
	goTime := passedTime.AsTime()
	return goTime.UnixNano() / int64(time.Millisecond)
}
