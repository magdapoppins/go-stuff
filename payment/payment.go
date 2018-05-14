package payment

type CreditCard struct {
	ownerName		string
	cardNo			string
	expMonth		int
	expYear			int
	ccv				int
	availableCred	float32
}

// This is the ctor function
func CreateCreditAccount(ownerName, cardNo string, expMonth, expMonth, ccv int) *CreditCard {
	return &CreditCard{
		ownerName:		ownerName,
		cardNo:			cardNo,
		expMonth: 		expMonth,
		expYear:		expYear,
		ccv:			ccv,
	}
}

func (c *CreditCard) OwnerName() string {
	return c.ownerName
}

func (c *CreditCard ) SetOwnerName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid owner name provided.")
	}

	c.ownerName = value
	return nil
}

func (c CreditCard) CardNo() string {
	return c.cardNo
}