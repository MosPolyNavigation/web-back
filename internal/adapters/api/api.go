package api

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MosPolyNavigation/web-back/internal/entity"
	"os"
	"strings"

	"github.com/go-resty/resty/v2"

	log "github.com/sirupsen/logrus"
)

const apiBase = "https://api.github.com/repos/MosPolyNavigation/navigationData/contents"

type api struct {
	client *resty.Client
	log    *log.Logger
}

func New(logger *log.Logger) *api {
	client := resty.New()
	return &api{
		client: client,
		log:    logger,
	}
}

type content struct {
	Content string `json:"content"`
}

func (a *api) GetPlan(campus, corpus string, floor int) (entity.Plan, error) {
	pathInRepository := fmt.Sprintf("%s/%s/%s-%d", campus, corpus, corpus, floor)
	planData, err := a.getJSON(pathInRepository)
	if err != nil {
		a.log.Debugf("Error getting plan from github repository: %v", err)
		return entity.Plan{}, err
	}
	planSVG, err := a.getSVG(pathInRepository)
	if err != nil {
		a.log.Debugf("Error getting plan SVG: %v", err)
		return entity.Plan{}, err
	}

	var plan entity.Plan
	err = json.Unmarshal([]byte(planData), &plan)
	if err != nil {
		a.log.Debugf("Error unmarshalling plan: %v", err)
		return entity.Plan{}, err
	}

	plan.SVG = planSVG

	return plan, nil
}

func (a *api) getSVG(pathInRepository string) (string, error) {
	svgFile, err := a.getFile(pathInRepository, "svg")
	if err != nil {
		return "", fmt.Errorf("failed to get file: %w", err)
	}

	return svgFile, nil
}

func (a *api) getJSON(pathInRepository string) (string, error) {
	jsonFile, err := a.getFile(pathInRepository, "json")
	if err != nil {
		return "", err
	}

	return jsonFile, nil
}

func (a *api) getFile(filePath string, format string) (string, error) {
	url := fmt.Sprintf("%s/%s.%s", apiBase, filePath, format)

	resp, err := a.client.R().EnableTrace().
		SetAuthToken(os.Getenv("GITHUB_SECRET")).
		Get(url)
	if err != nil {
		a.log.Errorf("Error getting %s: %s\n", url, err)
		return "", err
	}

	if resp.StatusCode() != 200 {
		a.log.Errorf("Error getting %s: %s\n", url, resp.Status())
		return "", errors.New(resp.Status())
	}

	body := resp.Body()
	a.log.Debugf("Body: %s\n", string(body))

	var c content
	if err := json.Unmarshal(body, &c); err != nil {
		a.log.Debugf("Error get content from response %s: %s\n", filePath, err)
		return "", err
	}

	return readContent(c.Content)
}

func readContent(content string) (string, error) {
	content = strings.TrimSpace(content)
	content = strings.Replace(content, "\n", "", -1)
	decodedContent, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return "", err
	}
	return string(decodedContent), nil
}
