// B"H
package main
import (
	"fmt"
	"strings"
	"strconv"
	// "reflect"
	"math"
	"os"
)
func main() {
	// fmt.Println("Hello, 世界")
	// fmt.Println(addHoures(12, 76))
	// var hours, minutes  = getTimeFromMinutes(76)
	// fmt.Printf("getTimeFromMinutes: %v:%v", hours, minutes)
	minus := []string{
		"0:00",
		"1:39",
		"1:03",
		"1:56",
		"0:00",
		"0:00",
		"0:00",
		"0:36",
		"1:07",
		"0:43",
		"0:30",
		"0:00",
		"0:00",
		"0:00",
		"0:46",
		"0:00",
		"0:01",
		"0:00",
		"0:00",
		"0:00",
		"0:00",
		"0:02",
		"0:00",
		"1:42",
		"0:28",
		"0:00",
	}
	plus := []string{
		"00:19",
		"00:00",
		"00:00",
		"00:00",
		"00:22",
		"00:00",
		"00:00",
		"00:00",
		"00:00",
		"00:00",
		"00:00",
		"04:22",
		"00:00",
		"00:00",
		"00:00",
		"00:17",
		"00:00",
		"01:55",
		"02:48",
		"00:00",
		"00:00",
		"00:00",
		"01:24",
		"00:00",
		"00:00",
		"00:00",
	}

	times := []string{
		"9:19", 
		"7:21",
		"7:57",
		"7:04",
		"9:22",
		"0:00",
		"0:00",
		"8:24",
		"7:53",
		"8:17",
		"8:30",
		"13:22",
		"0:00",
		"0:00",
		"8:14",
		"9:17",
		"8:59",
		"10:55",
		"11:48",
		"0:00",
		"0:00",
		"8:58",
		"10:24",
		"7:18",
		"8:32",
		"0:00",

	}
	// var hours1, minuts2 = addTime("12:27", "05:45")
	// fmt.Println()
	// fmt.Printf("addTime: %v:%v \n", hours1, minuts2)
	// // fmt.Println()
	var timesResMinus = addTimes(minus)
	fmt.Println()
	fmt.Printf("addTime missing: %v \n", timesResMinus)

	var timesResPlus = addTimes(plus)
	fmt.Println()
	fmt.Printf("addTime extra: %v \n", timesResPlus)

	var timesResTimes = addTimes(times)
	fmt.Println()
	fmt.Printf("addTime times: %v \n", timesResTimes)
	
	var hours, minuts = substractTime(timesResPlus, timesResMinus)
	fmt.Println()
	fmt.Printf("substractTime extra minus missing: %v:%v \n", hours, minuts)
	
}

// Time format "10:25"
func convertTimeToInts(time string) (int, int){
	timeSplited := strings.Split(time, ":")

	hourInt, err := strconv.Atoi(timeSplited[0])
	minutesInt, err := strconv.Atoi(timeSplited[1])

	if err != nil {
        fmt.Println(err)
        os.Exit(2)
	}

	return hourInt, minutesInt
}


func addHoures(hour int, hourPlus int) int {
	return hour + hourPlus
}


func addMinutes(minutes int, minutesPlus int) int {
	return minutes + minutesPlus
}


func addTime(time string, timeToAdd string) (int, int) {

	var hourInt, minutesInt = convertTimeToInts(time)
	var hourToAddInt, minutesToAddInt = convertTimeToInts(timeToAdd)
	
	var hours int = addHoures(hourInt, hourToAddInt)
	var minutes int = addMinutes(minutesInt, minutesToAddInt)
	
	
	if(minutes > 59){
		var hoursRemains, minutesRemains = getTimeFromMinutes(minutes)
		minutes = minutesRemains
		hours += hoursRemains
	}
	
	
	// fmt.Println("time", time)
	// fmt.Println("timeToAdd", timeToAdd)
	// fmt.Println("addTime.minutes:  ", minutes)
	return hours, minutes
}


func substractHoures(hour int, hourMinus int) int {
	return hour - hourMinus
}

func substractMinutes(minutes int, minutesMinus int) int {
	return minutes - minutesMinus
}


func substractTime(time string, timeToSubstract string) (int, int) {

	var hourInt, minutesInt = convertTimeToInts(time)
	var hourToSubtractInt, minutesToSubtractInt = convertTimeToInts(timeToSubstract)

	var hours int = substractHoures(hourInt, hourToSubtractInt)
	var minutes int = substractMinutes(minutesInt, minutesToSubtractInt)

	var minutesRemaining int

	if (hours > 0 && minutes < 0 ){
		minutesRemaining = int(math.Abs(float64(minutes)))
		// fmt.Println(minutesRemaining)
		hours--
		minutes = 60 - minutesRemaining
	}

	return hours, minutes
}


func addTimes(time []string) (string) {
	// fmt.Println(time)
	var sumOfTimes string = time[0]
	for index := 0; index < len(time) - 1 ; index++ {
		var sumOfTimesHoursTmp, sumOfTimesMinutesTmp = addTime(sumOfTimes, time[index + 1])
		sumOfTimes = fmt.Sprintf("%v:%v", sumOfTimesHoursTmp, sumOfTimesMinutesTmp)
		// fmt.Println("sumOfTimesHoursTmp: ", sumOfTimesHoursTmp, "sumOfTimesMinutesTmp: ", sumOfTimesMinutesTmp, "sumOfTimes: " , sumOfTimes)
		
	}
	return sumOfTimes
}


func getTimeFromMinutes(minutes int) (int, int) {
	var hours int = minutes / 60	
	var remainingMinutes int = minutes % 60
	// fmt.Println("getTimeFromMinutes.remainingMinutes: ", remainingMinutes)
	return hours, remainingMinutes
}
