package component

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"https_client/model"
	"https_client/utils"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type PostBody struct {
	MustContain string `json:"must_contain"`
	RegexpRule  string `json:"regexp_rule"`
	Related     string `json:"related"`
}

type Client struct {
	CliCrt     tls.Certificate
	CaCertPool *x509.CertPool
}

func (sel *Client) Run() {
	// loror 虽然为了测试，我没有设置输入的所有属性，但除了 id、isdeprecated、createtime 和 updatedtime 之外，其他属性都需要设置。
	// loror 目前还没有用于修改 del 的socket
	var input model.Account
	var ret addRes
	input.Account = "pig"
	input.ShortName = "pork"
	input.Describe = "pork is smelly"
	input.IsDeprecated = false
	ret = sel.addAccount(input)
	fmt.Printf("addAccount %+v\n", ret)

	var input1 model.EmailAccount
	input1.Address = "fruit@gg.com"
	input1.Password = "grapes"
	input1.Describe = "I like tables"
	input1.Risk = "weird hobby"
	input1.IsDeprecated = false
	ret = sel.addEmailAccount(input1)
	fmt.Printf("addEmailAccount %+v\n", ret)

	var input2 model.PhoneCard
	input2.Number = "1234456789"
	input2.Describe = "sweet vege"
	input2.Risk = "I like criminal"
	input2.IsDeprecated = false
	ret = sel.addPhoneCard(input2)
	fmt.Printf("addPhoneCard %+v\n", ret)

	var input3 model.Cellphone
	input3.Brand = "apple"
	input3.Model = "ad111"
	input3.System = "ad"
	input3.Risk = "I like sweet vegetable"
	input3.IsDeprecated = false
	ret = sel.addCellphone(input3)
	fmt.Printf("addCellphone %+v\n", ret)

	var input4 model.Link
	input4.Table0 = "accounts"
	input4.ID0 = 16
	input4.Table1 = "cellphones"
	input4.ID1 = 1
	sel.addLink(input4)

	var input5 model.Link
	input5.Table0 = "accounts"
	input5.ID0 = 16
	input5.Table1 = "email_accounts"
	input5.ID1 = 1
	sel.addLink(input5)

	var input6 model.Link
	input6.Table0 = "phone_cards"
	input6.ID0 = 1
	input6.Table1 = "accounts"
	input6.ID1 = 16
	sel.addLink(input6)

	var postBody PostBody
	// loror 用户必须设置以下三个属性之一，才能进行搜索
	postBody.MustContain = "beef"
	postBody.RegexpRule = "^.{1}a.{1}_.*S$"
	postBody.Related = "cloud server"
	var ret1 = sel.searchAccount(postBody)
	rj, _ := json.Marshal(ret1)
	fmt.Printf("--> %s", rj)
}

type addRes struct {
	Id  uint   `json:"id"`
	Err string `json:"err"`
}

func (sel *Client) addAccount(input model.Account) addRes {
	var err error
	input.Account, err = utils.EncryptStringWithRSA("./public_key.pem", input.Account, "salt")
	input.Password, err = utils.EncryptStringWithRSA("./public_key.pem", input.Password, "salt")

	var urlPrefix = Config.MyRequest.Url + Config.MyRequest.Port + "/"
	ret, err := sel.Post(urlPrefix, "addaccount", input)
	if err != nil {
		fmt.Printf("Post addaccount failed error: %v\n", err)
	} else {
		fmt.Println(ret)
	}
	var res addRes
	err = json.Unmarshal([]byte(ret), &res)
	if err != nil {
		panic(err)
	}
	return res
}

func (sel *Client) addEmailAccount(input model.EmailAccount) addRes {
	var err error
	input.Address, err = utils.EncryptStringWithRSA("./public_key.pem", input.Address, "salt")
	input.Password, err = utils.EncryptStringWithRSA("./public_key.pem", input.Password, "salt")
	input.SafeQuestions, err = utils.EncryptStringWithRSA("./public_key.pem", input.SafeQuestions, "salt")

	var urlPrefix = Config.MyRequest.Url + Config.MyRequest.Port + "/"
	ret, err := sel.Post(urlPrefix, "addemailaccount", input)
	if err != nil {
		fmt.Printf("Post addEmailAccount failed error: %v\n", err)
	} else {
		fmt.Println(ret)
	}
	var res addRes
	err = json.Unmarshal([]byte(ret), &res)
	if err != nil {
		panic(err)
	}
	return res
}

func (sel *Client) addPhoneCard(input model.PhoneCard) addRes {
	var err error
	input.Number, err = utils.EncryptStringWithRSA("./public_key.pem", input.Number, "salt")
	input.SimInfo, err = utils.EncryptStringWithRSA("./public_key.pem", input.SimInfo, "salt")

	var urlPrefix = Config.MyRequest.Url + Config.MyRequest.Port + "/"
	ret, err := sel.Post(urlPrefix, "addphonecard", input)
	if err != nil {
		fmt.Printf("Post addPhoneCard failed error: %v\n", err)
	} else {
		fmt.Println(ret)
	}
	var res addRes
	err = json.Unmarshal([]byte(ret), &res)
	if err != nil {
		panic(err)
	}
	return res
}

func (sel *Client) addCellphone(input model.Cellphone) addRes {
	var urlPrefix = Config.MyRequest.Url + Config.MyRequest.Port + "/"
	ret, err := sel.Post(urlPrefix, "addcellphone", input)
	if err != nil {
		fmt.Printf("Post addCellphone failed error: %v\n", err)
	} else {
		fmt.Println(ret)
	}
	var res addRes
	err = json.Unmarshal([]byte(ret), &res)
	if err != nil {
		panic(err)
	}
	return res
}

func (sel *Client) addLink(input model.Link) addRes {
	var urlPrefix = Config.MyRequest.Url + Config.MyRequest.Port + "/"
	ret, err := sel.Post(urlPrefix, "addlink", input)
	if err != nil {
		fmt.Printf("Post addLink failed error: %v\n", err)
	} else {
		fmt.Println(ret)
	}
	var res addRes
	err = json.Unmarshal([]byte(ret), &res)
	if err != nil {
		panic(err)
	}
	return res
}

type AccountDetail struct {
	Account         string `json:"account"`
	UserName        string `json:"user_name"`
	AccountPassword string `json:"account_password"`
	ShortName       string `json:"short_name"`
	AccountDescribe string `json:"account_describe"`
	AccountRisk     string `json:"account_risk"`

	Number      string `json:"number"`
	SimDescribe string `json:"sim_describe"`
	SimInfo     string `json:"sim_info"`

	EmailAddress  string `json:"email_address"`
	EmailPassword string `json:"email_password"`
	EmailRisk     string `json:"email_risk"`
	EmailDescribe string `json:"email_describe"`

	CellphoneBrand    string `json:"cellphone_brand"`
	CellphoneModel    string `json:"cellphone_model"`
	CellphoneSystem   string `json:"cellphone_system"`
	CellphoneDescribe string `json:"cellphone_describe"`
	CellphoneRisk     string `json:"cellphone_risk"`
}

func (sel *Client) searchAccount(postBody PostBody) []AccountDetail {

	var urlPrefix = Config.MyRequest.Url + Config.MyRequest.Port + "/"

	ret, err := sel.Post(urlPrefix, "searchaccountinfo", postBody)
	if err != nil {
		fmt.Printf("Post search failed error: %v\n", err)
	}
	var output []AccountDetail
	err = json.Unmarshal([]byte(ret), &output)
	if err != nil {
		panic(err)
	}
	sel.decryptAccountDetailList(output)
	return output
}

func (sel *Client) decryptAccountDetailList(input []AccountDetail) []AccountDetail {
	for i, _ := range input {
		var element = input[i]
		var err error
		if len(element.Account) > 300 {
			if element.Account, err = utils.DecryptStringWithRSA(Config.RsaConfig.PrivateKey, element.Account, "salt"); err != nil {
				panic(err)
			}
		}
		if len(element.AccountPassword) > 300 {
			if element.AccountPassword, err = utils.DecryptStringWithRSA(Config.RsaConfig.PrivateKey, element.AccountPassword, "salt"); err != nil {
				panic(err)
			}
		}
		if len(element.Number) > 300 {
			if element.Number, err = utils.DecryptStringWithRSA(Config.RsaConfig.PrivateKey, element.Number, "salt"); err != nil {
				panic(err)
			}
		}
		if len(element.SimInfo) > 300 {
			if element.SimInfo, err = utils.DecryptStringWithRSA(Config.RsaConfig.PrivateKey, element.SimInfo, "salt"); err != nil {
				panic(err)
			}
		}
		if len(element.EmailAddress) > 300 {
			if element.EmailAddress, err = utils.DecryptStringWithRSA(Config.RsaConfig.PrivateKey, element.EmailAddress, "salt"); err != nil {
				panic(err)
			}
		}
		if len(element.EmailPassword) > 300 {
			if element.EmailPassword, err = utils.DecryptStringWithRSA(Config.RsaConfig.PrivateKey, element.EmailPassword, "salt"); err != nil {
				panic(err)
			}
		}
		input[i] = element
	}
	return input
}

func NewClient() Client {
	pool := x509.NewCertPool()
	caCertPath := "ca.pem"
	caCrt, err := os.ReadFile(caCertPath)
	if err != nil {
		panic("os.ReadFile(caCertPath) faild")
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("client.pem", "client.key")
	if err != nil {
		panic("tls.LoadX509KeyPair faild")
	}
	var output Client
	output.CliCrt = cliCrt
	output.CaCertPool = pool
	return output
}

// Post 客户端请求
func (sel *Client) Post(urlPrefix string, url string, body interface{}) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      sel.CaCertPool,
			Certificates: []tls.Certificate{sel.CliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	newUrl := urlPrefix + url
	// 将结构体转换为 JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	// 发送 HTTP POST 请求
	response, err := client.Post(newUrl, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	// 读取响应体
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	// 将响应体转换为字符串
	responseString := string(responseBody)
	return responseString, nil
}

// Get 发送 HTTP GET 请求
func (sel *Client) Get(urlPrefix string, url string, params map[string]string) (string, error) {
	url = urlPrefix + url
	// 构建完整的 URL
	completeURL := buildCompleteURL(url, params)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      sel.CaCertPool,
			Certificates: []tls.Certificate{sel.CliCrt},
		},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(completeURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 将响应体转换为字符串
	responseString := string(responseBody)

	return responseString, nil
}

// 拼接URL
func buildCompleteURL(baseURL string, params map[string]string) string {
	var queryString strings.Builder
	queryString.WriteString(baseURL)

	if len(params) > 0 {
		queryString.WriteString("?")

		i := 0
		for key, value := range params {
			queryString.WriteString(url.QueryEscape(key))
			queryString.WriteString("=")
			queryString.WriteString(url.QueryEscape(value))

			if i < len(params)-1 {
				queryString.WriteString("&")
			}
			i++
		}
	}

	return queryString.String()
}
