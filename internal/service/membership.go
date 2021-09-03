package service

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	mod "github.com/nmluci/KissatenService/internal/models"
)

func (mem *mod.Members) Init() {
	file, err := os.Open("./storage/membership.data")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mem.AppendMem(ToMember(scanner.Text()))
	}
}

func ToMember(rawData string) (string, *mod.Member) {
	parsed := strings.Split(rawData, "#")
	id := uint(0)
	name := parsed[0]
	cash, err := strconv.ParseUint(parsed[1], 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	return name, &Member{id, name, uint32(cash)}
}

func (mem Members) AppendMem(name string, m *Member) error {
	if _, exist := mem[name]; exist {
		return errors.New("member already registered")
	} else {
		mem[name] = m
		return nil
	}
}

func (mem Members) GetMember(name string) *Member {
	if itm, exist := mem[name]; exist {
		return itm
	} else {
		return nil
	}
}

func (mem Members) RemoveMember(name string) error {
	if _, exist := mem[name]; exist {
		delete(mem, name)
		return nil
	} else {
		return errors.New("member not found")
	}
}

func (mem Members) UpdateMember(name string, data *Member) error {
	if _, exist := mem[name]; exist {
		mem[name] = data
		return nil
	} else {
		return errors.New("member not found")
	}
}

func (mem Members) Export() {
	file, err := os.Create("./storage/membership.data")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	for name, itm := range mem {
		file.WriteString(fmt.Sprintf("%s#%d\n", name, itm.Point))
	}
}

func (mem Members) MemberCount() int {
	return len(mem)
}
