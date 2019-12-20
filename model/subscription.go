package model

import (
	"time"
)

type SubscriptionStatus uint8

const (
	Pending = iota + 1
	Active
	OnHold
	Cancelled
	Expired
)

var SubscriptionStatuss = []SubscriptionStatus{
	Pending,
	Active,
	OnHold,
	Cancelled,
	Expired,
}

func (s SubscriptionStatus) String() string {
	switch s {
	case Pending:
		return "Pending"
	case Active:
		return "Active"
	case OnHold:
		return "OnHold"
	case Cancelled:
		return "Cancelled"
	case Expired:
		return "Expired"
	default:
		return ""
	}
}

//SubscriptionCache cache field customer_number+service_name
type SubscriptionCache struct {
	UUID             string                 `json:"uuid"`
	PackageID        string                 `json:"package_id"`
	StartDate        time.Time              `json:"start_date"`
	EndDate          time.Time              `json:"end_date"`
	Type             uint8                  `json:"type"`
	Meta             map[string]interface{} `json:"meta"`
	Duration         uint16                 `json:"duration"`
	IntervalTime     uint16                 `json:"interval_time"`
	Quota            uint16                 `json:"quota"`
	OldPrice         float64                `json:"old_price"`
	CustomerNumber   string                 `json:"customer_number"`
	Service          string                 `json:"service"`
	Status           SubscriptionStatus     `json:"status"`
	HaveTrialPackage bool                   `json:"have_trial_package"`
	TrialDuration    uint16                 `json:"trial_duration"`
}