package product

/*
 * @author       Intha-amnouay Marc
 * @mail         marc.inthaamnouay@gmail.com
 * @goversion    1.7
 * @run once     go run src/modules/product/product.go
 * @run routine  $GOPATH/bin/product
 */

import "start"

type product struct {
	Name   string
	Status int
	Owner  *start.UserStruct
}

type Product interface {
	isValid() bool
	setName()
	setStatus()
	setOwner()
	getName() string
	getStatus() int
	getOwner() string
}

/*
 * CreateProduct
 *      Create a product
 * @Params {name} string
 * @Params {status} int
 * @Params {owner} *UserStruct
 */
func CreateProduct(name string, status int, ow *start.UserStruct) *product {
	pr := &product{
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
func (p *product) IsValid() bool {
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
func (p *product) SetName(name string) {
	p.Name = name
}

/*
 * SetStatus
 *      Set the status of the product
 * @Params {status} int
 */
func (p *product) SetStatus(status int) {
	p.Status = status
}

/*
 * SetOwner
 *      Set the owner of the product
 * @Params {owner} UserStruct (From start package | user.go)
 */
func (p *product) SetOwner(owner *start.UserStruct) {
	p.Owner = owner
}

/*
 * GetName
 *      Return the name of the product
 * @Return {name} string
 */
func (p *product) GetName() string {
	return p.Name
}

/*
 * GetStatus
 *      Return the status of the product
 * @Return {status} int
 */
func (p *product) GetSatus() int {
	return p.Status
}

/*
 * GetOwner
 *      Get the owner of the product
 * @Return {owner} UserStruct (From start package | user.go)
 */
func (p *product) GetOwner() *start.UserStruct {
	return p.Owner
}
