package eureka

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
获取所有注册在eureka 的服务
*/
func Apps() Applications {
	client := &http.Client{}
	resp, err := client.Get("http://127.0.0.1:8761/eureka/apps")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var v Applications
	err1 := xml.Unmarshal([]byte(string(body)), &v) //将文件转化成对象
	if err != nil {
		fmt.Println(err1)
	}
	return v
}

type Applications struct {
	VersionsDelta string        `xml:"versions__delta"`
	AppsHashcode  string        `xml:"apps__hashcode"`
	Application   []Application `xml:"application"`
}

type Application struct {
	Name     string   `xml:"name"`
	Instance Instance `xml:"instance"`
}

type Instance struct {
	InstanceId                    string         `xml:"instanceId"`
	HostName                      string         `xml:"hostName"`
	App                           string         `xml:"app"`
	IpAddr                        string         `xml:"ipAddr"`
	Status                        string         `xml:"status"`
	Overriddenstatus              string         `xml:"overriddenstatus"`
	Port                          string         `xml:"port"`
	SecurePort                    string         `xml:"securePort"`
	CountryId                     string         `xml:"countryId"`
	HomePageUrl                   string         `xml:"homePageUrl"`
	StatusPageUrl                 string         `xml:"statusPageUrl"`
	HealthCheckUrl                string         `xml:"healthCheckUrl"`
	VipAddress                    string         `xml:"vipAddress"`
	SecureVipAddress              string         `xml:"secureVipAddress"`
	IsCoordinatingDiscoveryServer string         `xml:"isCoordinatingDiscoveryServer"`
	LastUpdatedTimestamp          string         `xml:"lastUpdatedTimestamp"`
	LastDirtyTimestamp            string         `xml:"lastDirtyTimestamp"`
	ActionType                    string         `xml:"actionType"`
	Metadata                      Metadata       `xml:"metadata"`
	LeaseInfo                     LeaseInfo      `xml:"leaseInfo"`
	DataCenterInfo                DataCenterInfo `xml:"dataCenterInfo"`
}

type Metadata struct {
	ManagementPort string `xml:"management.port"`
	JmxPort        string `xml:"jmx.port"`
}

type LeaseInfo struct {
	RenewalIntervalInSecs string `xml:"renewalIntervalInSecs"`
	DurationInSecs        string `xml:"durationInSecs"`
	RegistrationTimestamp string `xml:"registrationTimestamp"`
	LastRenewalTimestamp  string `xml:"lastRenewalTimestamp"`
	EvictionTimestamp     string `xml:"evictionTimestamp"`
	ServiceUpTimestamp    string `xml:"serviceUpTimestamp"`
}

type DataCenterInfo struct {
	Name string `xml:"name"`
}
