package exchange

import (
	"errors"
	"modules/mailSender"
	"modules/product"
	"modules/receiver"
	"time"
)

// Original exchange struct
type Exchange struct {
	Receiver  *receiver.UserStruct
	Product   product.Product
	StartDate time.Time
	EndDate   time.Time
}

// Interface export for the gomock
type eInterface interface {
	save() bool
}

// Original exchange function
func (e Exchange) Save() (bool, error) {

	dateNow, _errNow := GetFormatTime(time.Now())
	formatStartDate, _errStart := GetFormatTime(e.StartDate)
	formatEndDate, _errEnd := GetFormatTime(e.EndDate)

	if _errStart != nil || _errEnd != nil || _errNow != nil {
		return false, errors.New("time is not valid")
	}

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

// GetFormatTime (time.Time)
// parse the time and send a format time based on the layout
func GetFormatTime(inputTime time.Time) (time.Time, error) {
	const shortForm = "2006-Jan-02"

	timeFormat, _errTime := time.Parse(shortForm, inputTime.Format(shortForm))

	if _errTime != nil {
		return inputTime, errors.New("an error has occured while converting the time")
	}

	return timeFormat, nil
}

// BuildError build a custom error by comparing params
// Return error
func BuildError(content string, c bool, h bool, a string, r string) error {
	var newStr string

	if !c && !h {
		newStr = a + " & " + r + " " + content
	} else if !c {
		newStr = a + " " + content
	} else if !h {
		newStr = r + " " + content
	}

	return errors.New(newStr)
}
