package boredapi

import "strconv"

type Activity struct {
	Activity      string  `json:"activity"`
	Accessibility float64 `json:"accessibility"`
	Type          string  `json:"type"`
	Participants  uint    `json:"participants"`
	Price         float64 `json:"price"`
	Link          string  `json:"link"`
	Key           string  `json:"key"`
}

func RandomActivity() (Activity, error) {
	var activity Activity
	if err := queryApi(&ActivityQuery{}, &activity); err != nil {
		return Activity{}, err
	}
	return activity, nil
}

func FindActivity(query *ActivityQuery) (Activity, error) {
	var activity Activity
	if err := queryApi(query, &activity); err != nil {
		return Activity{}, err
	}
	return activity, nil
}

type ActivityQuery struct {
	params map[string]string
}

func NewActivityQuery() *ActivityQuery {
	query := ActivityQuery{}
	query.params = make(map[string]string)
	return &query
}

func (q *ActivityQuery) SetKey(key string) {
	q.params["key"] = key
}

func (q *ActivityQuery) SetType(activityType string) {
	q.params["type"] = activityType
}

func (q *ActivityQuery) SetParticipants(participants uint) {
	q.params["participants"] = strconv.FormatUint(uint64(participants), 10)
}

func (q *ActivityQuery) SetPrice(price float64) {
	q.params["price"] = strconv.FormatFloat(price, 'f', -1, 64)
}

func (q *ActivityQuery) SetMinPrice(minPrice float64) {
	q.params["minprice"] = strconv.FormatFloat(minPrice, 'f', -1, 64)
}

func (q *ActivityQuery) SetMaxPrice(maxPrice float64) {
	q.params["maxprice"] = strconv.FormatFloat(maxPrice, 'f', -1, 64)
}

func (q *ActivityQuery) SetAccessibility(accessibility float64) {
	q.params["accessibility"] = strconv.FormatFloat(accessibility, 'f', -1, 64)
}

func (q *ActivityQuery) SetMinAccessibility(minAccessibility float64) {
	q.params["minaccessibility"] = strconv.FormatFloat(minAccessibility, 'f', -1, 64)
}

func (q *ActivityQuery) SetMaxAccessibility(maxAccessibility float64) {
	q.params["maxaccessibility"] = strconv.FormatFloat(maxAccessibility, 'f', -1, 64)
}

func (q *ActivityQuery) Endpoint() string {
	return "activity"
}

func (q *ActivityQuery) Params() map[string]string {
	return q.params
}
