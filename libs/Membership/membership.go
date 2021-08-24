package MembershipLibs

import (
	"fmt"
)

var MemberList = Members{
	{0, "Lynne Fuyuna", 800},
	{1, "Fuyuna", 230},
	{2, "Lynne", 141},
}

func MemberCount() int {
	return len(MemberList)
}

func PrintMember() {
	for _, mem := range MemberList {
		fmt.Printf("%d %s %d\n", mem.ID, mem.name, mem.point)
	}
}
