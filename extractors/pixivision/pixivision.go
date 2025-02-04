package pixivision

import (
	"github.com/iawia002/lux/extractors"
	"github.com/iawia002/lux/parser"
	"github.com/iawia002/lux/request"
	"github.com/iawia002/lux/utils"
)

func init() {
	extractors.Register("pixivision", New())
}

type extractor struct{}

// New returns a pixivision extractor.
func New() extractors.Extractor {
	return &extractor{}
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string, option extractors.Options) ([]*extractors.Data, error) {
	html, err := request.Get(url, url, nil)
	if err != nil {
		return nil, err
	}
	title, urls, err := parser.GetImages(html, "am__work__illust  ", nil)
	if err != nil {
		return nil, err
	}

	parts := make([]*extractors.Part, 0, len(urls))
	for _, u := range urls {
		_, ext, err := utils.GetNameAndExt(u)
		if err != nil {
			return nil, err
		}
		size, err := request.Size(u, url)
		if err != nil {
			return nil, err
		}
		parts = append(parts, &extractors.Part{
			URL:  u,
			Size: size,
			Ext:  ext,
		})
	}

	streams := map[string]*extractors.Stream{
		"default": {
			Parts: parts,
		},
	}

	return []*extractors.Data{
		{
			Site:    "pixivision pixivision.net",
			Title:   title,
			Type:    extractors.DataTypeImage,
			Streams: streams,
			URL:     url,
		},
	}, nil
}
