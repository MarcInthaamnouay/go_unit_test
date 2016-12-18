package db

import "errors"

/*
 * @author       Intha-amnouay Marc
 * @mail         marc.inthaamnouay@gmail.com
 * @goversion    1.7
 * @run once     go run src/db/database.go
 * @run routine  $GOPATH/bin/start
 */

type Database struct {
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
 * MakeError
 *		a function which throw an error
 */
func makeError() error {
	err := errors.New("Not implemented")

	return err
}

/*
 * SaveProduct
 *      Save the product
 */
func (db database) saveProduct() error {
	return makeError()
}

/*
 * SaveUser
 *      Save the user
 */
func (db database) saveUser() error {
	return makeError()
}

/*
 * SaveExchange
 *      Save the exchange
 */
func (db database) saveExchange() error {
	return makeError()
}
