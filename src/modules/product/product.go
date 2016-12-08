package product

/*
 * @author       Intha-amnouay Marc
 * @mail         marc.inthaamnouay@gmail.com
 * @goversion    1.7
 * @run once     go run src/product/product.go
 * @run routine  $GOPATH/bin/product
 */

import "start"

type product struct {
	Name   string
	Status int
	Owner  start.UserStruct
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
func createProduct(name string, status int, owner *UserStruct) product {
	product := &product{
		Name:   name,
		Status: status,
		Owner:  owner,
	}

	return product
}

/*
 * IsValid
 *      Check if the product is valid
 * @Params {p} product
 */
func (p product) isValid() bool {
	if len(p.Name) > 0 && len(p.Owner) > 0 {
		if p.Owner.isValid() {
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
func (p *product) setName(name string) {
	p.Name = name
}

/*
 * SetStatus
 *      Set the status of the product
 * @Params {status} int
 */
func (p *product) setStatus(status int) {
	p.Status = status
}

/*
 * SetOwner
 *      Set the owner of the product
 * @Params {owner} UserStruct (From start package | user.go)
 */
func (p *product) setOwner(owner *start.UserStruct) {
	p.Owner = owner
}

/*
 * GetName
 *      Return the name of the product
 * @Return {name} string
 */
func (p *product) getName() string {
	return p.Name
}

/*
 * GetStatus
 *      Return the status of the product
 * @Return {status} int
 */
func (p *product) getSatus() int {
	return p.Status
}

/*
 * GetOwner
 *      Get the owner of the product
 * @Return {owner} UserStruct (From start package | user.go)
 */
func (p *product) getOwner() start.UserStruct {
	return p.Owner
}
