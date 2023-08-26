package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-freight-agreement-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-freight-agreement-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var header *[]dpfm_api_output_formatter.Header
	var item *[]dpfm_api_output_formatter.Item

	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				header = c.Header(mtx, input, output, errs, log)
			}()
		case "HeadersByBuyer":
			func() {
				header = c.HeadersByBuyer(mtx, input, output, errs, log)
			}()
		case "HeadersBySeller":
			func() {
				header = c.HeadersBySeller(mtx, input, output, errs, log)
			}()
		case "Item":
			func() {
				item = c.Item(mtx, input, output, errs, log)
			}()
		case "Items":
			func() {
				item = c.Items(mtx, input, output, errs, log)
			}()

		default:
		}
		if len(*errs) != 0 {
			break
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Header: header,
		Item:   item,
	}

	return data
}

func (c *DPFMAPICaller) Header(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := fmt.Sprintf("WHERE header.FreightOrder = %d ", input.Header.FreightOrder)

	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND header.IsCancelled = %v ", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND header.IsMarkedForDeletion = %v", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_freight_order_header_data AS header
		` + where + ` ORDER BY header.IsMarkedForDeletion ASC, header.IsCancelled ASC, header.FreightOrder DESC;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersByBuyer(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.Buyer != nil {
		where = fmt.Sprintf("%s\nAND Buyer = %v", where, *input.Header.Buyer)
	}
	if input.Header.HeaderBlockStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderBlockStatus = %t", where, *input.Header.HeaderBlockStatus)
	}
	if input.Header.HeaderBillingStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderBillingStatus != '%s'", where, *input.Header.HeaderBillingStatus)
	}
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND IsCancelled = %t", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_freight_order_header_data
		` + where + `;`,
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) HeadersBySeller(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Header {
	where := "WHERE 1 = 1"
	if input.Header.Seller != nil {
		where = fmt.Sprintf("%s\nAND Seller = %v", where, *input.Header.Seller)
	}
	if input.Header.HeaderBlockStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderBlockStatus = %t", where, *input.Header.HeaderBlockStatus)
	}
	if input.Header.HeaderBillingStatus != nil {
		where = fmt.Sprintf("%s\nAND HeaderBillingStatus != '%s'", where, *input.Header.HeaderBillingStatus)
	}
	if input.Header.IsCancelled != nil {
		where = fmt.Sprintf("%s\nAND IsCancelled = %t", where, *input.Header.IsCancelled)
	}
	if input.Header.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.Header.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_freight_order_header_data
		` + where + `;`,
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	data, err := dpfm_api_output_formatter.ConvertToHeader(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
