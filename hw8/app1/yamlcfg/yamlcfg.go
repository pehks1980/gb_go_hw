package yamlcfg

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
)

type Conf struct {
	Port         uint16 `yaml:"port"`
	Db_url       string `yaml:"db_url"`
	Jaeger_url   string `yaml:"jaeger_url"`
	Sentry_url   string `yaml:"sentry_url"`
	Kafka_broker string `yaml:"kafka_broker`
	Some_app_id  string `yaml:"some_app_id"`
	Some_app_key string `yaml:"some_app_key"`
	Clients      struct {
		Client_name   string   `yaml:"client_name"`
		Client_emails []string `yaml:"client_emails"`
	}
	Valid bool
}

//load config as method
func (c *Conf) GetConf(filename string) *Conf {

	path, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	corrPath := path + "/hw8/app1/" + filename

	fmt.Printf("conf path=%s, please check and set it according to your location!!! \n", corrPath)

	yamlFile, err := ioutil.ReadFile(corrPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

//validate config as method
func (c *Conf) ValidateConfig() string {
	//validate e-mail
	for _, email := range c.Clients.Client_emails {
		match, _ := regexp.MatchString("^\\S+@\\S+\\.\\S+$", email)

		if !match {
			c.Valid = false
			//fmt.Println(reflect.TypeOf(*c).Field(7))
			t := reflect.TypeOf(*c)
			field, _ := t.FieldByName("Clients")
			//field1, _ := t.FieldByName("Clients/Client_name")
			// до email не смог добраться (( sub-struct
			return string(field.Name) + " = " +email
		}
	}
	//validate URLs
	match, _ := regexp.MatchString("^postgres:\\/\\/[-@\\/?=A-Za-z:0-9]+", c.Db_url)

	if !match {
		c.Valid = false
		return string(reflect.TypeOf(*c).Field(1).Tag) + " = " + c.Db_url
	}
	urlPattern := "^https*:\\/\\/[-@\\/?=A-Za-z]+:+[0-9]+"
	match, _ = regexp.MatchString(urlPattern, c.Jaeger_url)

	if !match {
		c.Valid = false
		return string(reflect.TypeOf(*c).Field(2).Tag) + " = " + c.Jaeger_url
	}
	match, _ = regexp.MatchString(urlPattern, c.Sentry_url)

	if !match {
		c.Valid = false
		return string(reflect.TypeOf(*c).Field(3).Tag) + " = " + c.Sentry_url
	}

	//validate port https://tools.ietf.org/html/rfc1340
	if c.Port < 7009 {
		c.Valid = false
		return string(reflect.TypeOf(*c).Field(0).Tag) + " = " + strconv.Itoa(int(c.Port))
	}

	c.Valid = true
	return "ok"
}
