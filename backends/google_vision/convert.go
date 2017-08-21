package google_vision

import (
	"github.com/marshome/p-receipt/models"
	"google.golang.org/api/vision/v1"
)

func fromDetectedLanguage(p *vision.DetectedLanguage) (r *models.DetectedLanguage) {
	if p == nil {
		return nil
	}

	r = &models.DetectedLanguage{}
	r.LanguageCode = p.LanguageCode
	r.Confidence = p.Confidence

	return r
}

func fromTextProperty(p *vision.TextProperty) (r *models.TextProperty) {
	if p == nil {
		return nil
	}

	r = &models.TextProperty{}
	if p.DetectedBreak != nil {
		r.DetectedBreak = &models.DetectedBreak{
			Type:     p.DetectedBreak.Type,
			IsPrefix: p.DetectedBreak.IsPrefix,
		}
	}
	if p.DetectedLanguages != nil {
		r.DetectedLanguages = make([]*models.DetectedLanguage, 0)
		for _, v := range p.DetectedLanguages {
			r.DetectedLanguages = append(r.DetectedLanguages, fromDetectedLanguage(v))
		}
	}

	return r
}

func fromBoundingPoly(p *vision.BoundingPoly) (r *models.BoundingPoly) {
	if p == nil {
		return nil
	}

	r = &models.BoundingPoly{}
	if p.Vertices != nil {
		r.Vertices = make([]*models.Vertex, 0)
		for _, v := range p.Vertices {
			r.Vertices = append(r.Vertices, &models.Vertex{
				X: v.X,
				Y: v.Y,
			})
		}
	}

	return r
}

func fromSymbol(p *vision.Symbol) (r *models.Symbol) {
	if p == nil {
		return nil
	}

	r = &models.Symbol{}
	r.Property = fromTextProperty(p.Property)
	r.BoundingBox = fromBoundingPoly(p.BoundingBox)
	r.Text = p.Text

	return r
}

func fromWord(p *vision.Word) (r *models.Word) {
	if p == nil {
		return nil
	}

	r = &models.Word{}
	r.Property = fromTextProperty(p.Property)
	r.BoundingBox = fromBoundingPoly(p.BoundingBox)
	if p.Symbols != nil {
		r.Symbols = make([]*models.Symbol, 0)
		for _, v := range p.Symbols {
			r.Symbols = append(r.Symbols, fromSymbol(v))
		}
	}

	return r
}

func fromParagraph(p *vision.Paragraph) (r *models.Paragraph) {
	if p == nil {
		return nil
	}

	r = &models.Paragraph{}
	r.Property = fromTextProperty(p.Property)
	r.BoundingBox = fromBoundingPoly(p.BoundingBox)
	if p.Words != nil {
		r.Words = make([]*models.Word, 0)
		for _, v := range p.Words {
			r.Words = append(r.Words, fromWord(v))
		}
	}

	return r
}

func fromBlock(p *vision.Block) (r *models.Block) {
	if p == nil {
		return nil
	}

	r = &models.Block{}
	r.BlockType = p.BlockType
	r.Property = fromTextProperty(p.Property)
	r.BoundingBox = fromBoundingPoly(p.BoundingBox)
	if p.Paragraphs != nil {
		r.Paragraphs = make([]*models.Paragraph, 0)
		for _, v := range p.Paragraphs {
			r.Paragraphs = append(r.Paragraphs, fromParagraph(v))
		}
	}

	return r
}

func fromPage(p *vision.Page) (r *models.Page) {
	if p == nil {
		return nil
	}

	r = &models.Page{}
	r.Width = p.Width
	r.Height = p.Height
	r.Property = fromTextProperty(p.Property)
	if p.Blocks != nil {
		r.Blocks = make([]*models.Block, 0)
		for _, v := range p.Blocks {
			r.Blocks = append(r.Blocks, fromBlock(v))
		}
	}

	return r
}

func fromTextAnnotation(p *vision.TextAnnotation) (r *models.TextAnnotation) {
	if p == nil {
		return nil
	}

	r = &models.TextAnnotation{}
	r.Text = p.Text
	if p.Pages != nil {
		r.Pages = make([]*models.Page, 0)
		for _, v := range p.Pages {
			r.Pages = append(r.Pages, fromPage(v))
		}
	}

	return r
}
