package component

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type PostBody struct {
	MustContain     string         `json:"must_contain"`
	RegexpRule      string         `json:"regexp_rule"`
	Related         string         `json:"related"`
	Target          []SearchTarget `json:"target"`
	Count           int            `json:"count"`
	MoreTargetUrl   string         `json:"more_target_url"`
	MoreTargetParam interface{}    `json:"more_target_param"`
	IsEnd           bool           `json:"is_end"`
}

type SearchTarget struct {
	Index   int    `json:"index"`
	Content string `json:"content"`
}

type Client struct {
	CliCrt     tls.Certificate
	CaCertPool *x509.CertPool
}

func (sel *Client) Run() {
	var postBody PostBody
	postBody.MustContain = "must"
	var urlPrefix = Config.MyRequest.Url + Config.MyRequest.Port + "/"

	ret, err := sel.Post(urlPrefix, "add", postBody)
	if err != nil {
		fmt.Printf("Post search failed error: %v\n", err)
	} else {
		fmt.Println(ret)
	}

	var getParams = make(map[string]string)
	getParams["id"] = "sdfa"
	ret, err = sel.Get(urlPrefix, "search", getParams)
	if err != nil {
		fmt.Printf("Get search failed error: %v\n", err)
	} else {
		fmt.Println(ret)
	}
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
