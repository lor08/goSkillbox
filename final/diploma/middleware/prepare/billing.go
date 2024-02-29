package prepare

import (
	"github.com/lor08/goSkillbox/final/internal/models"
	"github.com/lor08/goSkillbox/final/internal/providers/billing"
	"log"
)

func GetBillingData(c chan models.BillingData) {
	defer close(c)
	billingProvider := billing.BillingProvider{
		Name: "Billing Status",
	}
	billingRes, err := billingProvider.GetStatus()
	if err != nil {
		log.Printf("can't get Billing status: %v", err)
		c <- models.BillingData{}
	}
	c <- billingRes
}
