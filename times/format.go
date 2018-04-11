package times

import (
	"fmt"
	"time"
)

func Timestamp() {
	now := time.Now()
	fmt.Println(now, now.Add(60 * time.Second), now.Add(60 * time.Second).Sub(now))

	timestamp := now.Unix()

	before := time.Unix(timestamp - 60, 0)
	fmt.Println(before.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(before.Format("10/Apr/2018:18:30:45 +0800"))
	fmt.Println(before.Format("2016-06-10"))

	if now.After(before) {
		fmt.Println("time passed")
	}

	//fmt.Println(now)
	//fmt.Println(now.AddDate(0, 1, -1).Format("2016-02-03"))
	t, err := time.Parse(time.UnixDate, "Sat Mar  7 15:06:39 PST 2015")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.UTC(), t.Format("2006-01-02 15:04:05"), t.Format("20060102150405"), t.Format("10/Apr/2018:18:30:45 +0800"))

	//ss := "10/Apr/2018:18:30:45 +0800"

}