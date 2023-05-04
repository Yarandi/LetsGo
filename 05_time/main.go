package main

import (
	"fmt"
	"time"
)

func main()  {

//present time
presentTime := time.Now().Format("01-02-2006 Monday 15:04:05")
fmt.Println(presentTime)

//generate a date time
createdTime := time.Date(1987, time.June, 11, 9, 50, 35, 123, time.UTC)
fmt.Println(createdTime)

//formatted created date time
formattedCreatedDate := createdTime.Format("01/02/2006 Monday 15:04:05")
fmt.Println("Ha Ha this is my birthDate, ignore the nono seconds :) ", formattedCreatedDate)


}