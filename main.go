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
	names := []string{
		"12:27", "5:45","06:03", "4:14",
	}
	var hours1, minuts2 = addTime("12:27", "05:45")
	fmt.Println()
	fmt.Printf("addTime: %v:%v \n", hours1, minuts2)
	// fmt.Println()
	var times = addTimes(names)
	fmt.Println()
	fmt.Printf("addTime: %v \n", times)

	// var hours3, minuts3 = substractTime("5:35", "5:45")
	// fmt.Println()
	// fmt.Printf("substractTime: %v:%v \n", hours3, minuts3)
	
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
		fmt.Println(minutesRemaining)
		hours--
		minutes = 60 - minutesRemaining
	}

	return hours, minutes
}


func addTimes(time []string) (string) {
	fmt.Println(time)
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

