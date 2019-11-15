package main

import (
	"fmt"
	"time"
)

func main(){
	// time package provides suuprt for times and durations
	now := time.Now()
	fmt.Println(now)
	// times are always associated with timeZone location for example in this case its CET


	birthday := time.Date(1994, 10, 01, 1,1,1,1, time.Local)
	fmt.Println(birthday)
	fmt.Println(birthday.Year())
	fmt.Println(birthday.Date())
	fmt.Println(birthday.Day())

	fmt.Println(birthday.Before(now)) /// birthday was before the todays date
	fmt.Println(birthday.After(now))
	fmt.Println(birthday.Equal(now))
	age := now.Sub(birthday)
	fmt.Println(age.Hours())

}
