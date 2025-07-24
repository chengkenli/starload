/*
 * Copyright (c) 2025 @chengkenli. All rights reserved.
 * Alternative more formal version for the second sentence:
 * "The applicable license is available at https://github.com/chengkenli."
 *
 * @Project: starload
 * @File: defg
 * @Author: chengkenli
 * @Created: 2025/7/13 11:03
 * @Description:
 */

package util

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

const (
	StarRocksServer   = "生产SR连接地址"
	StarRocksServerQa = "测试SR连接地址"
)

var (
	LogScroll   *container.Scroll
	logContent  = binding.NewString()
	Software    fyne.App
	Logger      *Logs
	MetaLink    []map[string]interface{}
	MetaData    []DataMeta
	ClientIPDec []ClientIPData
	H           Hosts
)

type Hosts struct {
	Ip string
}

type ConnectParms struct {
	Host       string
	Port       int
	User       string
	Pass       string
	Base       string
	Uri        string
	AaccessKey string
	SecretKey  string
}
type ConnectData struct {
	User, Password string
	Host           string
	Port           int
	Schema         string
}
type ImportData struct {
	User            string
	Password        string
	Host            string
	Port            int
	Schema          string
	File            string
	FileFormat      string
	Spacer          string
	FilterRatio     string
	StripOuterArray string
	SkipHeader      string
	ItemStore       DataMeta
}

type JesultData struct {
	TxnID                  int    `json:"TxnId"`
	Label                  string `json:"Label"`
	Status                 string `json:"Status"`
	Message                string `json:"Message"`
	NumberTotalRows        int    `json:"NumberTotalRows"`
	NumberLoadedRows       int    `json:"NumberLoadedRows"`
	NumberFilteredRows     int    `json:"NumberFilteredRows"`
	NumberUnselectedRows   int    `json:"NumberUnselectedRows"`
	LoadBytes              int    `json:"LoadBytes"`
	LoadTimeMs             int    `json:"LoadTimeMs"`
	BeginTxnTimeMs         int    `json:"BeginTxnTimeMs"`
	StreamLoadPlanTimeMs   int    `json:"StreamLoadPlanTimeMs"`
	ReadDataTimeMs         int    `json:"ReadDataTimeMs"`
	WriteDataTimeMs        int    `json:"WriteDataTimeMs"`
	CommitAndPublishTimeMs int    `json:"CommitAndPublishTimeMs"`
}

type ClientIPData struct {
	Ts              string `bson:"ts"`
	ComputerName    string `bson:"computer_name"`
	UserName        string `bson:"user_name"`
	ComputerType    string `bson:"computer_type"`
	ComputerStatus  string `bson:"computer_status"`
	IpAddress       string `bson:"ip_address"`
	SerialNumber    string `bson:"serial_number"`
	Brand           string `bson:"brand"`
	Model           string `bson:"model"`
	ComputerVersion string `bson:"computer_version"`
	BusinessUnit    string `bson:"business_unit"`
	BusinessFormat  string `bson:"business_format"`
	DataSource      string `bson:"data_source"`
	LastUpdateTime  string `bson:"last_update_time"`
	AiTime          string `bson:"ai_time"`
	AddTime         string `bson:"add_time"`
	LastActiveTime  string `bson:"last_active_time"`
}

type DataMeta struct {
	Segment    float32 `bson:"segment"`
	Loadbytes  int64   `bson:"loadbytes"`
	Loadcount  int64   `bson:"loadcount"`
	Envoptions string  `bson:"envoptions"`
	Sproptions string  `bson:"sproptions"`
	Txtoptions string  `bson:"txtoptions"`
	Dialog     string  `bson:"dialog"`
	State      int     `bson:"state"`
}
