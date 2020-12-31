package prototypepattern

import (
	"errors"
	"fmt"
)

type ShirtColor byte

type Shirt struct {
   Price float32
   SKU   string
   Color ShirtColor
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

var blackPrototype *Shirt = &Shirt{
	Price: 20.00,
	SKU:   "empty",
	Color: Black,
}

type ItemInfoGetter interface {
	GetInfo() string
}

func (s *Shirt) GetInfo()string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n",s.SKU, s.Color, s.Price)
}

type ShirtCloner interface{
	GetClone(m int) (ItemInfoGetter,error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)
	
type ShirtsCache struct {}
	
func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
		case White:
				newItem := *whitePrototype
				return &newItem, nil
		case Black:
				newItem := *blackPrototype
				return &newItem, nil
		case Blue:
				newItem := *bluePrototype
				return &newItem, nil
		default:
				return nil, errors.New("Shirt model not recognized")
	}
}

func GetShirtsCloner() ShirtCloner {
	return &ShirtsCache{}
}
	
func (i *Shirt) GetPrice() float32 {
	return i.Price
}