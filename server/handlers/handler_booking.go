package handlers

import (
	"net/http"
)

func ListBooking(w http.ResponseWriter, r *http.Request) {
	//session, err := service.CookieStore.Get(r, "session")
	//if err != nil{
	//	log.Error(err)
	//}
	//msg := service.IndexPageDataStruct{}
	//deletePastRecords()
	//query := elastic.NewMatchAllQuery()
	//res, err := el.Client.Search().
	//	Index(config.MAIN_INDEX).
	//	Pretty(true).
	//	Query(query).
	//	Size(500).
	//	Sort("time",true).
	//	Do(context.Background())
	//if err != nil {
	//	log.Error(err)
	//}
	//num := 1
	//for _, hit := range res.Hits.Hits {
	//	data,err := hit.Source.MarshalJSON()
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	var order2 order.Order
	//	order2.Number = num
	//	order2.Id = hit.Id
	//	err = json.Unmarshal(data,&order2)
	//	ourTime := time.Unix(order2.Time.Local().Unix() - 60 * 60 * 6,0)
	//	timeStr := ourTime.Format(time.RFC850)[:len(order2.Time.Format(time.RFC850)) - 7]
	//	order2.TimeString = timeStr
	//	if err != nil{
	//		log.Error(err)
	//	}
	//	msg.Orders = append(msg.Orders, order2)
	//	num++
	//}
	//
	//flashes := session.Flashes()
	//messages := service.GetFlashMessagesFromSession(flashes)
	//msg.SuccessMessages = messages.SuccessMessages
	//msg.ErrorMessages = messages.ErrorMessages
	//service.ExecuteTemplateWithHeader("index", r, w, msg)
	_, _ = w.Write([]byte("asdasdasd"))
}
