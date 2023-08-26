package requests

type Header struct {
	FreightOrder                         int      `json:"FreightOrder"`
	FreightOrderType                     *string  `json:"FreightOrderType"`
	ExternalFreightNumber                *string  `json:"ExternalFreightNumber"`
	PlannedFreight                       int      `json:"PlannedFreight"`
	FreightAgreement                     int      `json:"FreightAgreement"`
	FreightAgreementItem                 int      `json:"FreightAgreementItem"`
	FreightAgreementItemAvailableFreight int      `json:"FreightAgreementItemAvailableFreight"`
	FreightType                          string   `json:"FreightType"`
	FreightSpec                          string   `json:"FreightSpec"`
	FreightCalendar                      string   `json:"FreightCalendar"`
	PlannedFreightDepartureDate          string   `json:"PlannedFreightDepartureDate"`
	PlannedFreightDepartureTime          string   `json:"PlannedFreightDepartureTime"`
	ActualFreightDepartureDate           *string  `json:"ActualFreightDepartureDate"`
	ActualFreightDepartureTime           *string  `json:"ActualFreightDepartureTime"`
	LogisticsPartner                     int      `json:"LogisticsPartner"`
	DeliverToParty                       int      `json:"DeliverToParty"`
	DeliverToPlant                       string   `json:"DeliverToPlant"`
	DeliverFromParty                     int      `json:"DeliverFromParty"`
	DeliverFromPlant                     string   `json:"DeliverFromPlant"`
	MRPArea                              *string  `json:"MRPArea"`
	MRPController                        *string  `json:"MRPController"`
	FreightCapacityWeight                *float32 `json:"FreightCapacityWeight"` ///float32かfloat64
	FreightCapacityWeightUnit            *string  `json:"FreightCapacityWeightUnit"`
	CreationDateTime                     string   `json:"CreationDateTime"`
	LastChangeDateTime                   string   `json:"LastChangeDateTime"`
	IsMarkedForDeletion                  *int     `json:"IsMarkedForDeletion"`
}
