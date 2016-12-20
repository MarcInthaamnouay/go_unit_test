package exchange

import (
	"modules/db"
	"modules/product"
	"modules/receiver"
	"time"
)

type Exchange struct {
	receiver  *receiver.UserStruct
	product   product.Product
	startDate string
	endDate   string
	sender    mailSender.MailContructor
	db        db.database
}

type eInterface interface {
	save() boolean
}

func (e Exchange) save() boolean {
	dateNow := time.Now()
	formatStartDate, _errStart := time.Parse(RFC3339, e.startDate)
	formatEndDate, _errEnd := time.Parse(RFC3339, e.endDate)

	if e.receiver.IsValid() && e.product.IsValid() {
		// convert the start date and the end date to a real date
		if formatStartDate.After(dateNow) && formatEndDate.After(dateNow) {
			if formatStartDate.Before(formatEndDate) && formatEndDate.After(formatStartDate) {
				if receiver.GetAge() < 18 {
					e.sender.sendMail()
				}
				return true
			} else {
				return false
			}
		}
	}

	return true
}
