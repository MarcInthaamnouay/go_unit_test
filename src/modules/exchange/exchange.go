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

	dateNow, _errNow := GetFormatTime("now")
	formatStartDate, _errStart := GetFormatTime(e.StartDate)
	formatEndDate, _errEnd := GetFormatTime(e.EndDate)

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

func GetFormatTime(t string) (time.Time, error) {
	const shortForm = "2006-Jan-02"
	if t == "now" {
		now, _enow := time.Parse(shortForm, time.Now().Format(shortForm))
		if _enow != nil {
			return time.Now(), errors.New("Error while converting time.now to format shortForm")
		}

		return now, nil
	}

	newTime, _e := time.Parse(shortForm, t)

	if _e != nil {
		return time.Now(), errors.New("time is not valid")
	}

	return newTime, nil
}
