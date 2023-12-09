package activity

import (
	"fmt"
)

func Example_getEntryFromPath() {
	location := "file://[cwd]/activitytest/resource/activity-entry-v1.json"

	//buf, status := getEntryFromPath(location)
	//fmt.Printf("test: getEntryFromPath() -> [buf:%v] [status:%v]\n", len(buf), status)

	entries, status2 := getEntryFromPath(location)
	fmt.Printf("test: getEntryFromPath() -> [entries:%v] [status:%v]\n", len(entries), status2)

	//Output:
	//test: getEntryFromPath() -> [entries:2] [status:OK]

}
