package notifications

type Notification struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type NotificationWithUser struct {
	Id       int    `json:"id"`
	Text     string `json:"text"`
	Username string `json:"username"`
}
