package sdk

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"os"
	"sms/io"

	"strings"
)

var conf, _ = io.ReadFile()

func CreateClient() (_result *dysmsapi20170525.Client, _err error) {

	config := &openapi.Config{
		AccessKeyId:     tea.String(conf["AccessKeyId"].(string)),
		AccessKeySecret: tea.String(conf["AccessKeySecret"].(string)),
	}
	config.Endpoint = tea.String(conf["Endpoint"].(string))
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func send(args []*string, phone string, code string) (_err error) {
	client, _err := CreateClient()
	if _err != nil {
		return _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(conf["SignName"].(string)),
		TemplateCode:  tea.String(conf["TemplateCode"].(string)),
		TemplateParam: tea.String(code),
	}

	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}

		fmt.Println(tea.StringValue(error.Message))
		// 诊断地址
		var data interface{}
		d := json.NewDecoder(strings.NewReader(tea.StringValue(error.Data)))
		d.Decode(&data)
		if m, ok := data.(map[string]interface{}); ok {
			recommend, _ := m["Recommend"]
			fmt.Println(recommend)
		}
		_, _err = util.AssertAsString(error.Message)
		if _err != nil {
			return _err
		}
	}

	return _err
}

func SendSms(phone string, code string) {

	err := send(tea.StringSlice(os.Args[1:]), phone, code)
	if err != nil {
		panic(err)
	}
}
