package receipt

import (
	"github.com/marshome/p-vision/models"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var totalPriceNames = []string{
	"合计",
	"总计",
	"总金额",
	"總金額",
	"合計",
	"総金額",
	"총 금액은",
	"합계",
	"Total",
}

type priceArray []float64

type ExtractContext struct {
	TextAnnotation *models.TextAnnotation
	Result         *models.ReceiptInfo

	totalPriceNames     map[string]bool
	totalPrice          float64
	totalPriceNameFound bool
	priceArray          priceArray
}

func NewExtractContext(textAnnotation *models.TextAnnotation) (ctx *ExtractContext) {
	ctx = &ExtractContext{
		TextAnnotation: textAnnotation,
	}

	ctx.totalPriceNames = make(map[string]bool)
	for _, v := range totalPriceNames {
		ctx.totalPriceNames[strings.ToLower(v)] = true
	}

	return ctx
}

func (ctx *ExtractContext) isTotalPriceName(s string) bool { //todo 税率，折扣
	_, has := ctx.totalPriceNames[strings.ToLower(s)]
	return has
}

func (ctx *ExtractContext) isPrice(s string) (price float64, is bool) {
	price, err := strconv.ParseFloat(strings.Replace(s, ",", "", -1), 64)
	if err != nil {
		return 0, false
	}

	return price, true
}

func (ctx *ExtractContext) processParagraph(paragraph *models.Paragraph) {
	logrus.Info("processParagraph")

	if paragraph.Words != nil {
		var s string
		for _, word := range paragraph.Words {
			if word.Symbols != nil {
				for _, symbol := range word.Symbols {
					s += symbol.Text
					if symbol.Property != nil && symbol.Property.DetectedBreak != nil {
						breakType := symbol.Property.DetectedBreak.Type
						if breakType == "SPACE" {
							s += " "
						} else if breakType == "EOL_SURE_SPACE" {
							logrus.Info(s + "\\n")

							//first line would be title,but welcomes,todo
							if ctx.Result.Title == "" {
								ctx.Result.Title = s
							}

							//check total price name,last one
							if ctx.isTotalPriceName(s) {
								logrus.Infoln("total price name found " + s)
								ctx.totalPriceNameFound = true
							}

							//total price
							if ctx.totalPriceNameFound {
								price, isPrice := ctx.isPrice(s)
								if isPrice {
									logrus.Infoln("total price found " + s)
									ctx.totalPrice = price
								}
							}

							s = ""
						} else {
							logrus.Fatalln(breakType)
						}
					}
				}
			}
		}
	}
}

func (ctx *ExtractContext) processBlock(block *models.Block) {
	if block.Property != nil {
		if block.Property.DetectedBreak != nil {
			logrus.Infoln("processBlock break=" + block.Property.DetectedBreak.Type)
		}
	}

	if block.Paragraphs != nil {
		for _, paragraph := range block.Paragraphs {
			ctx.processParagraph(paragraph)
		}
	}
}

func (ctx *ExtractContext) processPage(page *models.Page) {
	//lang
	if page.Property != nil && page.Property.DetectedLanguages != nil {
		for _, lang := range page.Property.DetectedLanguages {
			ctx.Result.Lang = lang.LanguageCode
		}
	}

	if page.Blocks != nil {
		for _, block := range page.Blocks {
			ctx.processBlock(block)
		}
	}
}

func (ctx *ExtractContext) Process() {
	ctx.Result = &models.ReceiptInfo{}

	//full text
	ctx.Result.FullText = ctx.TextAnnotation.Text

	if ctx.TextAnnotation.Pages != nil {
		for _, page := range ctx.TextAnnotation.Pages {
			ctx.processPage(page)
		}
	}

	ctx.Result.TotalPrice = ctx.totalPrice
}
