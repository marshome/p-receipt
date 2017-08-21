package receipt

import (
	"github.com/marshome/p-receipt/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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
	s = strings.Trim(strings.ToLower(s), ": ")
	zap.L().Debug("isTotalPriceName " + s)
	_, has := ctx.totalPriceNames[s]
	return has
}

func (ctx *ExtractContext) isPrice(s string) (price float64, is bool) {
	zap.L().Debug("isPrice " + s)
	price, err := strconv.ParseFloat(strings.Replace(s, ",", "", -1), 64)
	if err != nil {
		return 0, false
	}

	return price, true
}

func (ctx *ExtractContext) processParagraph(paragraph *models.Paragraph) {
	zap.L().Debug("processParagraph")

	if paragraph.Words != nil {
		var w string
		var line string
		for _, word := range paragraph.Words {
			if word.Symbols != nil {
				for _, symbol := range word.Symbols {
					w += symbol.Text
					line += symbol.Text
					if symbol.Property != nil && symbol.Property.DetectedBreak != nil {
						breakType := symbol.Property.DetectedBreak.Type
						if breakType == "SPACE" {
							zap.L().Debug("SPACE")
							line += " "

							//check total price name,last one
							if ctx.isTotalPriceName(w) {
								zap.L().Debug("total price name found " + w)
								ctx.totalPriceNameFound = true
							}

							//total price
							if ctx.totalPriceNameFound {
								price, isPrice := ctx.isPrice(w)
								if isPrice {
									zap.L().Debug("total price found " + w)
									ctx.totalPrice = price
									ctx.totalPriceNameFound = false
								}
							}

							w = ""
						} else if breakType == "EOL_SURE_SPACE" || breakType == "LINE_BREAK" {
							zap.L().Debug(line + "\\n")

							//first line would be title,but welcomes,todo
							if ctx.Result.Title == "" {
								ctx.Result.Title = line
							}

							//check total price name,last one
							if ctx.isTotalPriceName(w) {
								zap.L().Debug("total price name found " + w)
								ctx.totalPriceNameFound = true
							}

							//total price
							if ctx.totalPriceNameFound {
								price, isPrice := ctx.isPrice(w)
								if isPrice {
									zap.L().Debug("total price found " + w)
									ctx.totalPrice = price
									ctx.totalPriceNameFound = false
								}
							}

							line = ""
							w = ""
						} else {
							panic(errors.New("unknown BreakType: " + breakType))
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
			zap.L().Debug("processBlock break=" + block.Property.DetectedBreak.Type)
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
