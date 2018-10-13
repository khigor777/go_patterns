package main

import "fmt"

type Producter interface {
	GetName() string
}

type  ProductApple struct {

}

func (ProductApple) GetName() string  {
	return "apple"
}

type  ProductGrapes string

func (ProductGrapes) GetName() string  {
	return "grapes "
}

type Factory interface {
	FactoryMethod() Producter
}

type FGrapes struct {
}

func (FGrapes) FactoryMethod() Producter  {
	return new(ProductGrapes)
}

type FApple struct {

}

func (FApple) FactoryMethod() Producter  {
	return &ProductApple{}
}

func main()  {
	r := []Factory{&FGrapes{}, &FApple{}}
	for _, v := range r{
		fmt.Println(v.FactoryMethod().GetName())
	}
}
