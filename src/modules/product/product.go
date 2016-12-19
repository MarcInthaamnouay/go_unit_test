package product

import "modules/receiver"

/*
 * @author       Intha-amnouay Marc
 * @mail         marc.inthaamnouay@gmail.com
 * @goversion    1.7
 * @run once     go run src/modules/product/product.go
 * @run routine  $GOPATH/bin/product
 */

type Product struct {
	Name   string
	Status int
	Owner  *receiver.UserStruct
}

type P interface {
	IsValid() bool
	GetName() string
	SetName() string
	GetOwner() string
	SetOwner() string
	GetStatus() string
	CreateProduct(n string, s int, ow *receiver.UserStruct) (string, int, *receiver.UserStruct)
}

/*
 * CreateProduct
 *      Create a product
 * @Params {name} string
 * @Params {status} int
 * @Params {owner} *UserStruct
 */
func CreateProduct(name string, status int, ow *receiver.UserStruct) *Product {
	pr := &Product{
		Name:   name,
		Status: status,
		Owner:  ow,
	}

	return pr
}

/*
 * IsValid
 *      Check if the product is valid
 * @Params {p} product
 */
func (p *Product) IsValid() bool {
	if len(p.Name) > 0 && p.Owner != nil {
		if p.Owner.IsValid() {
			return true
		}
	} else {
		return false
	}

	return true
}

/*
 * SetName
 *      Set the name of the product
 * @Params {name} string
 */
func (p *Product) SetName(name string) {
	p.Name = name
}

/*
 * SetStatus
 *      Set the status of the product
 * @Params {status} int
 */
func (p *Product) SetStatus(status int) {
	p.Status = status
}

/*
 * SetOwner
 *      Set the owner of the product
 * @Params {owner} UserStruct (From start package | user.go)
 */
func (p *Product) SetOwner(owner *receiver.UserStruct) {
	p.Owner = owner
}

/*
 * GetName
 *      Return the name of the product
 * @Return {name} string
 */
func (p *Product) GetName() string {
	return p.Name
}

/*
 * GetStatus
 *      Return the status of the product
 * @Return {status} int
 */
func (p *Product) GetSatus() int {
	return p.Status
}

/*
 * GetOwner
 *      Get the owner of the product
 * @Return {owner} UserStruct (From start package | user.go)
 */
func (p *Product) GetOwner() *receiver.UserStruct {
	return p.Owner
}
