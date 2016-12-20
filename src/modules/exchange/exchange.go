package exchange

import (
	"errors"
	"fmt"
	"modules/mailSender"
	"modules/product"
	"modules/receiver"
	"time"
)

type Exchange struct {
	Receiver  *receiver.UserStruct
	Product   product.Product
	StartDate string
	EndDate   string
}

type eInterface interface {
	save() bool
}

func (e Exchange) Save() (bool, error) {

	const shortForm = "2006-Jan-02"
	dateNowStr := time.Now().Format(shortForm)
	formatStartDate, _errStart := time.Parse(shortForm, e.StartDate)
	formatEndDate, _errEnd := time.Parse(shortForm, e.EndDate)
	dateNow, _errNow := time.Parse(shortForm, dateNowStr)

	if _errStart != nil || _errEnd != nil || _errNow != nil {
		return false, errors.New("time is not valid")
	}

	fmt.Println("before start date ? ", formatStartDate.Before(formatEndDate))
	fmt.Println("after start date ? ", formatEndDate.After(formatStartDate))

	if e.Receiver.IsValid() && e.Product.IsValid() {
		// convert the start date and the end date to a real date
		if formatStartDate.After(dateNow) && formatEndDate.After(dateNow) {
			if formatStartDate.Before(formatEndDate) && formatEndDate.After(formatStartDate) {
				if e.Receiver.GetAge() < 18 {
					mail.SendMail(e.Receiver.GetEmail(), "You're under 18", "Order information")
				}
				return true, nil
			} else {
				return false, errors.New("Please check the date")
			}
		} else {
			return false, errors.New("Please use a date that's after today")
		}
	}

	return false, errors.New("Product or User isn't valid")

}
