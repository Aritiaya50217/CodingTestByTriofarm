package response

type Reponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	TopicID uint   `json:"topic_id"`
	Index   int    `json:"index"`
}
