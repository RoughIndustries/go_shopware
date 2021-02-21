package models

type SalesChannelAnalytics struct {
	Entity
	TrackingID   string        `json:"trackingId"`
	Active       bool          `json:"active"`
	TrackOrders  bool          `json:"trackOrders"`
	SalesChannel *SalesChannel `json:"salesChannel"`
}
