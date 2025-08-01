package utils

type Reponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	TopicID uint   `json:"topic_id"`
	Index   int    `json:"index"`
}

func IsDuplicateIndex(indexs []int) bool {
	indexMap := make(map[int]bool)
	for _, index := range indexs {
		if indexMap[index] {
			return true
		}
		indexMap[index] = true
	}
	return false
}
