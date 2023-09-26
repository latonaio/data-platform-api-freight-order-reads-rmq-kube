package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string      `json:"connection_key"`
	Result              bool        `json:"result"`
	RedisKey            string      `json:"redis_key"`
	Filepath            string      `json:"filepath"`
	APIStatusCode       int         `json:"api_status_code"`
	RuntimeSessionID    string      `json:"runtime_session_id"`
	BusinessPartnerID   *int        `json:"business_partner"`
	ServiceLabel        string      `json:"service_label"`
	APIType             string      `json:"api_type"`
	Message             interface{} `json:"message"`
	APISchema           string      `json:"api_schema"`
	Accepter            []string    `json:"accepter"`
	Deleted             bool        `json:"deleted"`
	SQLUpdateResult     *bool       `json:"sql_update_result"`
	SQLUpdateError      string      `json:"sql_update_error"`
	SubfuncResult       *bool       `json:"subfunc_result"`
	SubfuncError        string      `json:"subfunc_error"`
	ExconfResult        *bool       `json:"exconf_result"`
	ExconfError         string      `json:"exconf_error"`
	APIProcessingResult *bool       `json:"api_processing_result"`
	APIProcessingError  string      `json:"api_processing_error"`
}

type Message struct {
	Header *[]Header `json:"Header"`
}

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
