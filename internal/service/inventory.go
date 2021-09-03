package service

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var invFile = "./storage/inventory.data"

// Initialize the Inventory by retrieving data from the prefer database (PlainText for now),
// and load it into the Inventory Struct
func (inv models.Inventory) Init() {
	file, err := os.Open(invFile)
	if err != nil {
		log.Fatalf("File Not Found!")
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatalf("Error'd")
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inv.AppendItem(toItem(scanner.Text()))
	}
}

// Convert Raw String into its corresponding datatypes
func toItem(rawData string) (string, *Item) {
	id := 0
	str := strings.Split(rawData, "#")
	name := str[0]
	sum, err := strconv.ParseInt(str[1], 10, 0)
	if err != nil {
		log.Fatal(err)
	}

	price, err := strconv.ParseUint(str[2], 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	return name, &Item{id, name, int(sum), uint(price)}
}

func (inv Inventory) ItemCount() int {
	return len(inv)
}

func (inv Inventory) GetItemByName(name string) *Item {
	if itm, exist := inv[name]; exist {
		return itm
	} else {
		return nil
	}
}

func (invLst Inventory) AppendItem(name string, itm *Item) error {
	if _, exist := invLst[name]; exist {
		return errors.New("data already existed")
	} else {
		invLst[name] = itm
		return nil
	}
}

func (invLst Inventory) RemoveItem(name string) error {
	if _, exist := invLst[name]; exist {
		delete(invLst, name)
		return nil
	} else {
		return errors.New("data isn't existed")
	}
}

func (invLst *Inventory) Export() {
	file, err := os.Create("./storage/mem.txt")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for name, itm := range *invLst {
		file.WriteString(fmt.Sprintf("%s#%d%d\n", name, itm.Price, itm.Sum))
	}
}

func (inv Inventory) Verbose() {
	fmt.Print("Item List\n")
	for name, item := range inv {
		fmt.Printf("┬─%s\n├ SCH %d,00\n└ Qty: %d\n", name, item.Price, item.Sum)
	}
}

func (inv Inventory) VerboseItem(itm string) {
	for name, item := range inv {
		if strings.EqualFold(name, itm) {
			// fmt.Printf("%s %d %d\n", name, item.price, item.sum)
			fmt.Printf("┬─%s\n├ SCH %d,00\n└ Qty: %d\n", name, item.Price, item.Sum)
		}
	}
}

func (itm *Item) Verbose() {
	fmt.Printf("┬─%s\n├ SCH %d,00\n└ Qty: %d\n", itm.Name, itm.Price, itm.Sum)
}
