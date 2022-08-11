package parse

import (
	"encoding/json"
	"fmt"
	"io"

	common "github.com/jxcorra/peparse/internal/common"
	"golang.org/x/net/html"
)

func ParsePage(tokenizer *html.Tokenizer, schema *[]common.Search) (common.Parsed, error) {
	data := common.Parsed{}
	dataItemsForTextCounter := 1

	for {
		token := tokenizer.Next()

		if token == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				return data, nil
			}
			return nil, fmt.Errorf("cannot parse page: %v", tokenizer.Err())
		}

		if token == html.TextToken && len(data) == dataItemsForTextCounter {
			currentDataItem := data[len(data)-1]
			if currentDataItem.WithText {
				currentDataItem.Data["text"] = string(tokenizer.Text())
			}
			dataItemsForTextCounter++
		}

		tag, hasAttr := tokenizer.TagName()
		attrs := parseAttrs(tokenizer, hasAttr)

		fillParsedData(tag, schema, &data, &attrs)
	}
}

func Serialize(data *common.Parsed) (string, error) {
	collected := []common.Data{}
	for _, dataItem := range *data {
		collected = append(collected, dataItem.GetData())
	}
	serialized, err := json.Marshal(collected)
	if err != nil {
		return "", err
	}

	return string(serialized), nil
}

func parseAttrs(tokenizer *html.Tokenizer, hasAttr bool) map[string]string {
	attrs := map[string]string{}

	if hasAttr {
		for {
			attrName, attrValue, moreAttrs := tokenizer.TagAttr()
			attrs[string(attrName)] = string(attrValue)
			if !moreAttrs {
				break
			}
		}
	}

	return attrs
}

func fillParsedData(tag []byte, schema *[]common.Search, data *common.Parsed, attrs *map[string]string) {
	for _, search := range *schema {
		if search.Key.Element != nil && *search.Key.Element == string(tag) {
			if search.Key.Class != nil {
				class, found := (*attrs)["class"]
				if !found || *search.Key.Class != class {
					continue
				}
			}

			dataItem := common.DataItem{
				WithText: search.WithText,
				Data:     common.Data{},
			}
			*data = append(*data, dataItem)

			for _, parseItem := range search.Parse {
				if parseItem.Attribute != nil {
					value, found := (*attrs)[*parseItem.Attribute]
					if found {
						dataItem.Data[*parseItem.Attribute] = value
					}
				}
			}
		}
	}
}
