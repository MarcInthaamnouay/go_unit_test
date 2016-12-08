package exchange

import ("Receiver"
        "Product"
)

// Interface of the golang

type Receiver struct{
    email string
    firstname string
    name string
    age int
}

type ExchangeStruct struct{
    receiver Receiver
    product Product
}

/*
 * Exchange Interface
 *   
 */

type Exchange interface{
    constructObj() User
    isValid() bool
    save() bool
}




