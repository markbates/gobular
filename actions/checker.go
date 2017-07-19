package actions

import (
	"regexp"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/markbates/gobular/models"
	"github.com/pkg/errors"
)

func NewChecker(c buffalo.Context) error {
	c.Set("checker", models.Checker{})
	return c.Render(200, r.HTML("checker/new.html"))
}

func RunChecker(c buffalo.Context) error {
	checker := models.Checker{}
	if err := c.Bind(&checker); err != nil {
		return errors.WithStack(err)
	}
	c.Set("checker", checker)

	rx, err := regexp.Compile(checker.Expression)
	if err != nil {
		c.Set("compile_error", err.Error())
		if c.Request().Method == "GET" {
			return c.Render(422, r.HTML("checker/new.html"))
		}
		return c.Render(200, r.Template("text/html", "checker/_results.html"))
	}

	results := []models.Result{}
	for i, s := range strings.Split(checker.TestString, "\n") {
		s = strings.TrimSpace(s)
		res := rx.FindAllStringSubmatch(s, -1)
		if len(res) > 0 {
			rr := models.Result{
				Num:     i + 1,
				Line:    s,
				Matches: []string{},
				// Matches: res[0][1:],
			}
			for _, r := range res {
				if len(r) > 1 {
					rr.Matches = append(rr.Matches, r[1])
				}
			}
			results = append(results, rr)
		}
	}

	c.Set("results", results)

	if c.Request().Method == "GET" {
		return c.Render(200, r.HTML("checker/new.html"))
	}
	return c.Render(200, r.Template("text/html", "checker/_results.html"))
}