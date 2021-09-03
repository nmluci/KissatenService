package cashier

import (
	"errors"
	"fmt"
	"log"

	inv "github.com/nmluci/KissatenService/internal/Inventory"
	mem "github.com/nmluci/KissatenService/internal/Membership"
)

var ItemStorage = make(inv.Inventory)
var MemberList = make(mem.Members)

func Init() {
	ItemStorage.Init()
	MemberList.Init()
	fmt.Println("init done!")
}

func (cart *UserList) Buy(user *mem.Member, name string, sum int) error {
	tempCart := &UserShoppingCart{}

	if user == nil {
		user = &mem.Member{
			ID:    0,
			Name:  "Meta",
			Point: 177013,
		}
	}

	if c := cart.GetCart(user); c != nil {
		tempCart = c
		log.Print(fmt.Sprintf("Logged as %s", tempCart.User.Name))
	} else {
		tempCart.User = user
		tempCart.Cart = make(inv.Inventory)
		log.Print(fmt.Sprintf("Signed in as %s", tempCart.User.Name))
	}

	itm := ItemStorage.GetItemByName(name)
	if itm == nil {
		err := errors.New("item not existed")
		log.Println(err)
		return err
	}

	if itm.Sum < sum {
		log.Print("attempted to buy item more than what it had")
		sum = itm.Sum
	}

	if _, exist := tempCart.Cart[itm.Name]; exist {
		tempCart.Cart[itm.Name].Sum += sum
		log.Print(fmt.Sprintf("Added %s x%d", itm.Name, sum))
	} else {
		tempCart.Cart[itm.Name] = &inv.Item{itm.Id, itm.Name, sum, itm.Price}
		log.Print(fmt.Sprintf("Appended %s x%d", itm.Name, sum))
		*cart = append(*cart, tempCart)
	}
	return nil
}

func (cart *UserList) RemoveItem(usr *mem.Member, name string, sum int) error {
	crt := cart.GetCart(usr)
	if crt == nil {
		return errors.New("user doesn't have a cart")
	}

	if itm, exists := crt.Cart[name]; !exists {
		return errors.New("item doesn't exist")
	} else {
		if sum > itm.Sum {
			log.Print("attempted to remove more than what it bought")
			delete(crt.Cart, name)
		} else {
			itm.Sum -= sum
		}
	}
	return nil
}

func (cart UserList) GetCart(user *mem.Member) *UserShoppingCart {
	if user == nil {
		return nil
	}

	for _, name := range cart {
		if name.User == user {
			return name
		}
	}
	return nil
}

func (cart UserList) GetCartByName(name string) *UserShoppingCart {
	for _, usc := range cart {
		if usc.User.Name == name {
			return usc
		}
	}
	return nil
}

func (cart *UserList) RemoveCart(user *mem.Member) error {
	if user == nil {
		return errors.New("user not found")
	}

	idx := -1
	for id, usr := range *cart {
		if usr.User == user {
			idx = id
		}
	}
	if idx == -1 {
		return errors.New("user not found")
	} else {
		*cart = append((*cart)[:idx], (*cart)[idx+1:]...)
	}
	return nil
}

func (cart *UserList) Checkout(name string) error {
	var price uint
	usc := cart.GetCartByName(name)
	if usc == nil {
		return errors.New("user not found")
	}

	for key, itm := range usc.Cart {
		if ItemStorage[key].Sum > itm.Sum {
			ItemStorage[key].Sum -= itm.Sum
		} else {
			ItemStorage.RemoveItem(key)
		}
		price += itm.Price
	}

	if usc.User.Point >= uint32(price) {
		usc.User.Point -= uint32(price) / 10
	} else {
		log.Print("cash not enough for this transaction")
		return errors.New("cash not enough for this transaction")
	}

	cart.RemoveCart(usc.User)
	ItemStorage.Export()
	return nil
}

func (cart UserList) Verbose() {
	if len(cart) == 0 {
		fmt.Printf("No Data!")
	} else {
		fmt.Println("itm list")
		for _, usr := range cart {
			usr.Verbose()
		}
	}
}

func (usr UserShoppingCart) Verbose() {
	fmt.Println(usr.User.Name)
	usr.Cart.Verbose()
}
