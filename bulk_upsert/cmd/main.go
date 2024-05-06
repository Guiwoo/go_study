package main

import (
	"bulk_upsert/database"
	"bulk_upsert/dto"
	"bulk_upsert/entity"
	"bulk_upsert/log"
	"encoding/json"
	"os"
)

var (
	jsonFilePath string
	fileName     = "/config.json"
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

	cfgData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		_log.WithError(err).Fatalf("fail to read josn file %+v", jsonFilePath)
	}

	var cfg Config
	if err := json.Unmarshal(cfgData, &cfg); err != nil {
		_log.WithError(err).Fatal("fail to parse config file")
	}

	var data []dto.Vendor
	rawData, err := os.ReadFile(cfg.Source)

	if err != nil {
		_log.WithError(err).Fatalf("fail to open json file %+v", jsonFilePath)
	}

	if err := json.Unmarshal(rawData, &data); err != nil {
		_log.WithError(err).Fatalf("fail to unmarshal json file %+v", jsonFilePath)
	}

	_log.Debugf("have read data total len is %d", len(data))

	var (
		tbVendors      = make([]database.TableLabeler, 0, len(data))
		tbAddresses    = make([]database.TableLabeler, 0, len(data))
		tbPaymentTerms = make([]database.TableLabeler, 0, len(data))
		tbContacts     = make([]database.TableLabeler, 0, len(data))
		tbTaxCodes     = make([]database.TableLabeler, 0, len(data))
	)

	for _, vendor := range data {
		if vendor.Address != nil {
			tbAddresses = append(tbAddresses, vendor.Address.ToEntity())
		}

		if vendor.PaymentTerm != nil {
			tbPaymentTerms = append(tbPaymentTerms, vendor.PaymentTerm.ToEntity())
		}

		if vendor.TaxCode != nil {
			tbTaxCodes = append(tbTaxCodes, vendor.TaxCode.ToEntity())
		}

		if vendor.Contact != nil {
			tbContacts = append(tbContacts, vendor.Contact.ToEntity())
		}

		tbVendors = append(tbVendors, vendor.ToEntity())
	}

	myDB := database.NewMysql(cfg.Dns)

	if err := myDB.AutoMigrate(
		entity.Address{}, entity.Contact{}, entity.TaxCode{}, entity.PaymentTerm{}, entity.Vendor{},
	); err != nil {
		panic(err)
	}

	if err := myDB.BulkUpsert(tbPaymentTerms); err != nil {
		panic(err)
	}
	_log.Infof("success bulk upsert payment term")

	if err := myDB.BulkUpsert(tbTaxCodes); err != nil {
		panic(err)
	}
	_log.Infof("success bulk upsert tax code")

	if err := myDB.BulkUpsert(tbContacts); err != nil {
		panic(err)
	}
	_log.Infof("success bulk upsert contract")

	if err := myDB.BulkUpsert(tbAddresses); err != nil {
		panic(err)
	}
	_log.Infof("success bulk upsert address")

	if err := myDB.BulkUpsert(tbVendors); err != nil {
		panic(err)
	}
	_log.Infof("success bulk upsert vendor")
}
