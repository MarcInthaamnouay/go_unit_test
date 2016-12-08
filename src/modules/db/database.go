package db

/*
 * @author       Intha-amnouay Marc
 * @mail         marc.inthaamnouay@gmail.com
 * @goversion    1.7
 * @run once     go run src/db/database.go
 * @run routine  $GOPATH/bin/start
 */

type database struct {
	username string
	password string
	hostname string
	port     int
}

type dbInterface interface {
	saveProduct()
	saveUser()
	saveExchange()
}

/*
 * SaveProduct
 *      Save the product
 */
func (db database) saveProduct() {

}

/*
 * SaveUser
 *      Save the user
 */
func (db database) saveUser() {

}

/*
 * SaveExchange
 *      Save the exchange
 */
func (db database) saveExchange() {

}
