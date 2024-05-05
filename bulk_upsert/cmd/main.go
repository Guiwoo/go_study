package main

import (
	"bulk_upsert/database"
	"bulk_upsert/dto"
	"bulk_upsert/entity"
	"bulk_upsert/log"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var (
	jsonFilePath string
	fileName     = "/bulk_upsert.json"
)

func init() {
	if len(os.Args) > 2 {
		jsonFilePath = os.Args[1]
	} else {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		jsonFilePath = path + fileName
	}
}

func main() {
	_log := log.NewCustomLog()

	var data []dto.Vendor
	rawData, err := os.ReadFile(jsonFilePath)

	if err != nil {
		_log.WithError(err).Fatalf("fail to open json file %+v", jsonFilePath)
	}

	if err := json.Unmarshal(rawData, &data); err != nil {
		_log.WithError(err).Fatalf("fail to unmarshal json file %+v", jsonFilePath)
	}

	_log.Debugf("Have read data total len is %d", len(data))

	var (
		tbVendors      = make([]entity.Vendor, 0, len(data))
		tbAddresses    = make([]entity.Address, 0, len(data))
		tbPaymentTerms = make([]entity.PaymentTerm, 0, len(data))
		tbContracts    = make([]entity.Contract, 0, len(data))
		tbTaxCodes     = make([]entity.TaxCode, 0, len(data))
	)

	for _, vendor := range data {
		tbAddresses = append(tbAddresses, vendor.Address.ToEntity())

		tbPaymentTerms = append(tbPaymentTerms, vendor.PaymentTerm.ToEntity())

		tbTaxCodes = append(tbTaxCodes, vendor.TaxCode.ToEntity())

		tbContracts = append(tbContracts, vendor.Contract.ToEntity())

		tbVendors = append(tbVendors, vendor.ToEntity())
	}

	myDB := database.NewCustom()

	if err := myDB.AutoMigrate(
		entity.Address{}, entity.Contract{}, entity.TaxCode{}, entity.PaymentTerm{}, entity.Vendor{},
	); err != nil {
		panic(err)
	}

}
