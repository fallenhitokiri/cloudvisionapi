package cloudvisionapi

import (
	"context"
	"encoding/base64"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/vision/v1"
)

func detect(img []byte, name string) (*APIResponse, error) {
	response := &APIResponse{Name: name}

	service, err := newService()

	if err != nil {
		return nil, err
	}

	req := newBatchRequest(img)

	data, err := service.Images.Annotate(req).Do()

	if err != nil {
		return nil, err
	}

	parseLabels(response, data)
	parseProperties(response, data)
	parseLogo(response, data)

	return response, nil
}

func newService() (*vision.Service, error) {
	ctx := context.Background()

	client, err := google.DefaultClient(ctx, vision.CloudPlatformScope)

	if err != nil {
		return nil, err
	}

	return vision.New(client)
}

func newBatchRequest(img []byte) *vision.BatchAnnotateImagesRequest {
	req := &vision.AnnotateImageRequest{
		Image: &vision.Image{
			Content: base64.StdEncoding.EncodeToString(img),
		},
		Features: []*vision.Feature{
			{
				Type:       "LABEL_DETECTION",
				MaxResults: 5,
			},
			{
				Type:       "IMAGE_PROPERTIES",
				MaxResults: 5,
			},
			{
				Type:       "LOGO_DETECTION",
				MaxResults: 5,
			},
			{
				Type:       "TEXT_DETECTION",
				MaxResults: 5,
			},
		},
	}

	return &vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{req},
	}
}

func parseLabels(response *APIResponse, data *vision.BatchAnnotateImagesResponse) {
	if labels := data.Responses[0].LabelAnnotations; len(labels) > 0 {
		for _, label := range labels {
			l := &Label{
				Description: label.Description,
				Score:       label.Score,
				Confidence:  label.Confidence,
			}
			response.Labels = append(response.Labels, l)
		}
	}
}

func parseProperties(response *APIResponse, data *vision.BatchAnnotateImagesResponse) {
	if colors := data.Responses[0].ImagePropertiesAnnotation.DominantColors.Colors; len(colors) > 0 {
		for _, color := range colors {
			c := &Color{
				Alpha: color.Color.Alpha,
				Blue:  color.Color.Blue,
				Green: color.Color.Green,
				Red:   color.Color.Red,
				Score: color.Score,
			}
			response.Colors = append(response.Colors, c)
		}
	}
}

func parseLogo(response *APIResponse, data *vision.BatchAnnotateImagesResponse) {
	if logos := data.Responses[0].LogoAnnotations; len(logos) > 0 {
		for _, logo := range logos {
			l := &Logo{
				Description: logo.Description,
				Locale:      logo.Locale,
				Score:       logo.Score,
				Confidence:  logo.Confidence,
			}
			response.Logos = append(response.Logos, l)
		}
	}
}

func parseText(response *APIResponse, data *vision.BatchAnnotateImagesResponse) {
	if texts := data.Responses[0].LogoAnnotations; len(texts) > 0 {
		for _, text := range texts {
			t := &Text{
				Description: text.Description,
				Locale:      text.Locale,
				Score:       text.Score,
				Confidence:  text.Confidence,
			}
			response.Texts = append(response.Texts, t)
		}
	}
}
