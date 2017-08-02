package handler

import (
	api "github.com/marshome/p-vision/api/receipt/server/models"
	"github.com/marshome/p-vision/models"
)

func fromReceiptInfo(p *models.ReceiptInfo) (r *api.ReceiptInfo) {
	if p == nil {
		return nil
	}

	r = &api.ReceiptInfo{}
	r.Lang = p.Lang
	r.FullText = p.FullText
	r.Title = p.Title
	r.TotalPrice = p.TotalPrice

	return r
}

func fromDetectedLanguage(p *models.DetectedLanguage) (r *api.DetectedLanguage) {
	if p == nil {
		return nil
	}

	r = &api.DetectedLanguage{}
	r.LanguageCode = p.LanguageCode
	r.Confidence = p.Confidence

	return r
}

func fromTextProperty(p *models.TextProperty) (r *api.TextProperty) {
	if p == nil {
		return nil
	}

	r = &api.TextProperty{}
	if p.DetectedBreak != nil {
		r.DetectedBreak = &api.DetectedBreak{
			Type:     p.DetectedBreak.Type,
			IsPrefix: p.DetectedBreak.IsPrefix,
		}
	}
	if p.DetectedLanguages != nil {
		r.DetectedLanguages = make([]*api.DetectedLanguage, 0)
		for _, v := range p.DetectedLanguages {
			r.DetectedLanguages = append(r.DetectedLanguages, fromDetectedLanguage(v))
		}
	}

	return r
}

func fromBoundingPoly(p *models.BoundingPoly) (r *api.BoundingPoly) {
	if p == nil {
		return nil
	}

	r = &api.BoundingPoly{}
	if p.Vertices != nil {
		r.Vertices = make([]*api.Vertex, 0)
		for _, v := range p.Vertices {
			r.Vertices = append(r.Vertices, &api.Vertex{
				X: v.X,
				Y: v.Y,
			})
		}
	}

	return r
}

func fromSymbol(p *models.Symbol) (r *api.Symbol) {
	if p == nil {
		return nil
	}

	r = &api.Symbol{}
	r.Property = fromTextProperty(p.Property)
	r.BoundingBox = fromBoundingPoly(p.BoundingBox)
	r.Text = p.Text

	return r
}

func fromWord(p *models.Word) (r *api.Word) {
	if p == nil {
		return nil
	}

	r = &api.Word{}
	r.Property = fromTextProperty(p.Property)
	r.BoundingBox = fromBoundingPoly(p.BoundingBox)
	if p.Symbols != nil {
		r.Symbols = make([]*api.Symbol, 0)
		for _, v := range p.Symbols {
			r.Symbols = append(r.Symbols, fromSymbol(v))
		}
	}

	return r
}

func fromParagraph(p *models.Paragraph) (r *api.Paragraph) {
	if p == nil {
		return nil
	}

	r = &api.Paragraph{}
	r.Property = fromTextProperty(p.Property)
	r.BoundingBox = fromBoundingPoly(p.BoundingBox)
	if p.Words != nil {
		r.Words = make([]*api.Word, 0)
		for _, v := range p.Words {
			r.Words = append(r.Words, fromWord(v))
		}
	}

	return r
}

func fromBlock(p *models.Block) (r *api.Block) {
	if p == nil {
		return nil
	}

	r = &api.Block{}
	r.BlockType = p.BlockType
	r.Property = fromTextProperty(p.Property)
	r.BoundingBox = fromBoundingPoly(p.BoundingBox)
	if p.Paragraphs != nil {
		r.Paragraphs = make([]*api.Paragraph, 0)
		for _, v := range p.Paragraphs {
			r.Paragraphs = append(r.Paragraphs, fromParagraph(v))
		}
	}

	return r
}

func fromPage(p *models.Page) (r *api.Page) {
	if p == nil {
		return nil
	}

	r = &api.Page{}
	r.Width = p.Width
	r.Height = p.Height
	r.Property = fromTextProperty(p.Property)
	if p.Blocks != nil {
		r.Blocks = make([]*api.Block, 0)
		for _, v := range p.Blocks {
			r.Blocks = append(r.Blocks, fromBlock(v))
		}
	}

	return r
}

func fromTextAnnotation(p *models.TextAnnotation) (r *api.TextAnnotation) {
	if p == nil {
		return nil
	}

	r = &api.TextAnnotation{}
	r.Text = p.Text
	if p.Pages != nil {
		r.Pages = make([]*api.Page, 0)
		for _, v := range p.Pages {
			r.Pages = append(r.Pages, fromPage(v))
		}
	}

	return r
}
