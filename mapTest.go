package main

func main() {
	paymentTimePartnerIdJobIdSchemeMap := make(map[int]map[int]map[int]string)

	//jobIdSchemeMap := make(map[int]string)
	jobIdSchemeMap := map[int]string{
		101: "scheme1",
		102: "scheme2",
		103: "scheme3",
		104: "scheme4",
	}
	paymentTimeStart := 4000200100
	for jobId, scheme := range jobIdSchemeMap {
		paymentTimeStart++
		partnerIdJobIdSchemeMap := paymentTimePartnerIdJobIdSchemeMap[paymentTimeStart]

		if partnerIdJobIdSchemeMap == nil {
			partnerIdJobIdSchemeMap = make(map[int]map[int]string)
			paymentTimePartnerIdJobIdSchemeMap[paymentTimeStart] = partnerIdJobIdSchemeMap
		}

	}
	//arr := []string{"scheme1", "scheme2", "scheme3"}
	//
	//arr1 := make([]string,0)
}
