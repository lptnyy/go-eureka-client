package eureka

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//xml +="<instance>"
//xml +="<instanceId>demo-order2:11101</instanceId>"
//xml +="<hostName>127.0.0.1</hostName>"
//xml +="<app>DEMO-ORDER2</app>"
//xml +="<ipAddr>127.0.0.1</ipAddr>"
//xml +="<status>UP</status>"
//xml +="<overriddenstatus>UNKNOWN</overriddenstatus>"
//xml +="<port enabled=\"true\">11101</port>"
//xml +="<securePort enabled=\"false\">443</securePort>"
//xml +="<countryId>1</countryId>"
//xml +="<dataCenterInfo class=\"com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo\">"
//xml +="<name>MyOwn</name>"
//xml +="</dataCenterInfo>"
//xml +="<metadata class=\"java.util.Collections$EmptyMap\"/>"
//xml +="<vipAddress>demo-order2</vipAddress>"
//xml +="<secureVipAddress>demo-order2</secureVipAddress>"
//xml +="<isCoordinatingDiscoveryServer>false</isCoordinatingDiscoveryServer>"
//xml +="<lastUpdatedTimestamp>1540186708769</lastUpdatedTimestamp>"
//xml +="<lastDirtyTimestamp>1540186708747</lastDirtyTimestamp>"
//xml +="</instance>"

/**
获取所有注册在eureka 的服务
*/
func Region(regionXml RegionXml) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://127.0.0.1:8761/eureka/apps/demo-order2", strings.NewReader(createRegionXml(regionXml)))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", "application/xml")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
}

/*
	生成注册信息
*/
func createRegionXml(regionXml RegionXml) string {
	var xml string
	xml += "<instance>"
	xml += "<instanceId>" + regionXml.InstanceId + "</instanceId>"
	xml += "<hostName>" + regionXml.HostName + "</hostName>"
	xml += "<app>" + regionXml.App + "</app>"
	xml += "<ipAddr>" + regionXml.IpAddr + "</ipAddr>"
	xml += "<status>" + regionXml.Status + "</status>"
	xml += "<overriddenstatus>" + regionXml.Overriddenstatus + "</overriddenstatus>"
	xml += "<port enabled=\"true\">" + regionXml.Port + "</port>"
	xml += "<securePort enabled=\"false\">" + regionXml.SecurePort + "</securePort>"
	xml += "<countryId>" + regionXml.CountryId + "</countryId>"
	xml += "<dataCenterInfo class=\"com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo\">"
	xml += "<name>" + regionXml.Name + "</name>"
	xml += "</dataCenterInfo>"
	xml += "<metadata class=\"java.util.Collections$EmptyMap\"/>"
	xml += "<vipAddress>" + regionXml.VipAddress + "</vipAddress>"
	xml += "<secureVipAddress>" + regionXml.SecureVipAddress + "</secureVipAddress>"
	xml += "<isCoordinatingDiscoveryServer>" + regionXml.IsCoordinatingDiscoveryServer + "</isCoordinatingDiscoveryServer>"
	xml += "<lastUpdatedTimestamp>" + regionXml.LastUpdatedTimestamp + "</lastUpdatedTimestamp>"
	xml += "<lastDirtyTimestamp>" + regionXml.LastDirtyTimestamp + "</lastDirtyTimestamp>"
	xml += "</instance>"
	return xml
}

/**
注册信息保存
*/
type RegionXml struct {
	InstanceId                    string
	HostName                      string
	App                           string
	IpAddr                        string
	Status                        string
	Overriddenstatus              string
	Port                          string
	SecurePort                    string
	CountryId                     string
	DataCenterInfo                string
	Name                          string
	VipAddress                    string
	SecureVipAddress              string
	IsCoordinatingDiscoveryServer string
	LastUpdatedTimestamp          string
	LastDirtyTimestamp            string
}
