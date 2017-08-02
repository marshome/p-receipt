package receipt

import (
	"github.com/marshome/p-vision/models"
	"github.com/sirupsen/logrus"
)

type ExtractContext struct {
	TextAnnotation *models.TextAnnotation
	Result         *models.ReceiptInfo
}

func NewExtractContext(textAnnotation *models.TextAnnotation) (ctx *ExtractContext) {
	return &ExtractContext{
		TextAnnotation: textAnnotation,
	}
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
							if ctx.Result.Title == "" {
								ctx.Result.Title = s
							}
							logrus.Info(s + "\\n")
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
	ctx.Result.FullText = ctx.TextAnnotation.Text

	if ctx.TextAnnotation.Pages != nil {
		for _, page := range ctx.TextAnnotation.Pages {
			ctx.processPage(page)
		}
	}
}
