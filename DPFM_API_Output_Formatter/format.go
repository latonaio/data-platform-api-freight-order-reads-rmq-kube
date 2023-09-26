package dpfm_api_output_formatter

import (
	"data-platform-api-freight-order-reads-rmq-kube/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func ConvertToHeader(rows *sql.Rows) (*[]Header, error) {
	defer rows.Close()
	header := make([]Header, 0)

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Header{}

		err := rows.Scan(
			&pm.FreightOrder,
			&pm.FreightOrderType,
			&pm.ExternalFreightNumber,
			&pm.PlannedFreight,
			&pm.FreightAgreement,
			&pm.FreightAgreementItem,
			&pm.FreightAgreementItemAvailableFreight,
			&pm.FreightType,
			&pm.FreightSpec,
			&pm.FreightCalendar,
			&pm.PlannedFreightDepartureDate,
			&pm.PlannedFreightDepartureTime,
			&pm.ActualFreightDepartureDate,
			&pm.ActualFreightDepartureTime,
			&pm.LogisticsPartner,
			&pm.DeliverToParty,
			&pm.DeliverToPlant,
			&pm.DeliverFromParty,
			&pm.DeliverFromPlant,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.FreightCapacityWeight,
			&pm.FreightCapacityWeightUnit,
			&pm.CreationDateTime,
			&pm.LastChangeDateTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &header, err
		}

		data := pm
		header = append(header, Header{
			FreightOrder:                         data.FreightOrder,
			FreightOrderType:                     data.FreightOrderType,
			ExternalFreightNumber:                data.ExternalFreightNumber,
			PlannedFreight:                       data.PlannedFreight,
			FreightAgreement:                     data.FreightAgreement,
			FreightAgreementItem:                 data.FreightAgreementItem,
			FreightAgreementItemAvailableFreight: data.FreightAgreementItemAvailableFreight,
			FreightType:                          data.FreightType,
			FreightSpec:                          data.FreightSpec,
			FreightCalendar:                      data.FreightCalendar,
			PlannedFreightDepartureDate:          data.PlannedFreightDepartureDate,
			PlannedFreightDepartureTime:          data.PlannedFreightDepartureTime,
			ActualFreightDepartureDate:           data.ActualFreightDepartureDate,
			ActualFreightDepartureTime:           data.ActualFreightDepartureTime,
			LogisticsPartner:                     data.LogisticsPartner,
			DeliverToParty:                       data.DeliverToParty,
			DeliverToPlant:                       data.DeliverToPlant,
			DeliverFromParty:                     data.DeliverFromParty,
			DeliverFromPlant:                     data.DeliverFromPlant,
			MRPArea:                              data.MRPArea,
			MRPController:                        data.MRPController,
			FreightCapacityWeight:                data.FreightCapacityWeight,
			FreightCapacityWeightUnit:            data.FreightCapacityWeightUnit,
			CreationDateTime:                     data.CreationDateTime,
			LastChangeDateTime:                   data.LastChangeDateTime,
			IsMarkedForDeletion:                  data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &header, nil
	}

	return &header, nil
}
