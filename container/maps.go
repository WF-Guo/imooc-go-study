package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":    "gwf",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	m2 := make(map[string]int) //m2 == empty map
	var m3 map[string]int      //m3 == nil
	fmt.Println(m, m2, m3)

	for k, v := range m {
		fmt.Println(k, v) //unordered
	}

	fmt.Println("Getting values------")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if courseName1, ok := m["corse"]; ok {
		fmt.Println(courseName1, ok)
	} else {
		fmt.Println("key does not exist")
	}

	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
