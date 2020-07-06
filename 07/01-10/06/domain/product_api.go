package domain

type ProductAPI interface {
	ProductList() (ProductList, error)
}
