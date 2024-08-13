package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gookit/goutil/dump"
)

var client *elasticsearch.Client

func init() {
	config := elasticsearch.Config{
		//注意地址协议，如果是http不能写成https，否则会报错：first record does not look like a tls handshake
		Addresses: []string{"http://192.168.201.86:9200"},
		Username:  "elastic",
		Password:  "bP70esSLTKVXuGZ0O1Rd",
	}
	c, err := elasticsearch.NewClient(config)
	if err != nil {
		fmt.Println("connect es failed,err:", err)
	}
	client = c
}

func main() {
	//baseInfo()
	//node()
	index()
}

func baseInfo() {
	//获取es信息
	info, err := client.Info()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(info)
}

func node() {
	nodes := client.Nodes
	info, err := nodes.Info()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(info)
}

//使用Marshal序列化需要注意结构体中的字段一定要大写，否则无法序列化
type ScaleEnterprise struct {
	Area             string                   `json:"area,omitempty"`
	Country          string                   `json:"country,omitempty"`
	NeaIndList       map[string]interface{}   `json:"nea_ind_list,omitempty"`
	HsIndList        []map[string]interface{} `json:"hs_ind_list,omitempty"`
	ScaleEnterprises map[string]interface{}   `json:"scale_enterprises,omitempty"`
	CompanyName      string                   `json:"company_name,omitempty"`
	CompanyCode      string                   `json:"company_code,omitempty"`
}

//func (s ScaleEnterprise) String() string {
//	return fmt.Sprintf("%v", s.HsIndList, s.NeaIndList, s.ScaleEnterprises)
//}

func index() {
	//response, err := client.Search(func(request *esapi.SearchRequest) {
	//	request.Index = []string{"ads_lget_scale_enterprises"}
	//	var size *int
	//	size = new(int)
	//	*size = 10
	//	request.Size = size
	//})
	//if err != nil {
	//	fmt.Println("get index error,", err)
	//}
	//fmt.Println(response.Body)

	var buf bytes.Buffer

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{"term": map[string]interface{}{"is_delete": 0}},
				},
				"should": []map[string]interface{}{
					{"term": map[string]interface{}{"country": "中国"}},
					{"terms": map[string]interface{}{"nea_ind_list.cs_second": []string{"煤炭开采和洗选业"}}},
				},
			},
		},
	}
	//为什么NewEncoder传入的参数是指针类型，因为Buffer实现了Writer的Write方法，该方法的接收者是指针类型
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		fmt.Println("查询条件写入失败", err)
		return
	}
	fmt.Println("查询语句: \n", query)
	response, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("ads_lget_scale_enterprises"),
		client.Search.WithSize(1),
		client.Search.WithFrom(10),
		client.Search.WithBody(&buf),
		//返回json格式
		client.Search.WithPretty(),
	)
	if err != nil {
		fmt.Println("search error,", err)
	}
	defer response.Body.Close()
	var r map[string]interface{}
	//解码，即读取数据
	if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
		fmt.Println("读取数据失败,", err)
	}
	//(map[string]interface{})["hits"]:因为key是接口类型，所以使用断言来判断类型。获取hits中的hits，因为它也是一个map[string]interface{}
	//([]interface{}):获取对应的value，它是一个切片
	total := r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]
	fmt.Println("total:", total)
	scales := make([]*ScaleEnterprise, 0)
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		//依然是接口断言判断类型。尝试将hit断言为map[string]interface{}类型，并获取"_source"字段对应的值，
		//如果断言成功，i将会是"_source"字段对应的值，v是一个布尔值，表示断言是否成功
		if _, v := hit.(map[string]interface{})["_source"]; v {
			//fmt.Println("items:", hit.(map[string]interface{})["_source"])
			var model ScaleEnterprise
			//序列化
			body, err := json.Marshal(hit.(map[string]interface{})["_source"])
			if err != nil {
				fmt.Println("序列化失败", err)
				continue
			}
			//反序列化，即转换成对应的json格式
			if err := json.Unmarshal(body, &model); err != nil {
				fmt.Println("反序列化失败", err)
				continue
			}
			scales = append(scales, &model)
		}
	}
	//for _, v := range scales {
	//	//fmt.Printf("%#v\n", v)
	//	dump.P(v)
	//}
	dump.P(scales)
}
