package actions

import (
	"regexp"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/slices"
	"github.com/markbates/gobular/models"
	"github.com/pkg/errors"
)

func NewChecker(c buffalo.Context) error {
	x := models.Expression{
		Expression: exp,
		TestString: testString,
	}
	c.Set("expression", x)
	return c.Render(200, r.HTML("checker/new.html"))
}

func RunChecker(c buffalo.Context) error {
	x := &models.Expression{}
	if err := c.Bind(x); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)

	// if this expression and test string combo already exists, load it and use it
	exq := tx.Where("expression = ? and test_string = ?", x.Expression, x.TestString)
	if b, err := exq.Exists(x); err == nil && b {
		err = exq.First(x)
		if err != nil {
			return errors.WithStack(err)
		}
		results := &models.Results{}
		err = tx.Where("expression_id = ?", x.ID).Order("num asc").All(results)
		if err != nil {
			return errors.WithStack(err)
		}

		c.Set("expression", x)
		c.Set("results", results)
		return c.Render(200, r.Template("text/html", "checker/_results.html"))
	}

	rx, err := regexp.Compile(x.Expression)
	if err != nil {
		c.Set("compile_error", err.Error())
		if c.Request().Method == "GET" {
			return c.Render(422, r.HTML("checker/new.html"))
		}
		return c.Render(200, r.Template("text/html", "checker/_results.html"))
	}

	err = tx.Create(x)
	if err != nil {
		return errors.WithStack(err)
	}

	results := models.Results{}
	for i, s := range strings.Split(x.TestString, "\n") {
		s = strings.TrimSpace(s)
		res := rx.FindAllStringSubmatch(s, -1)
		if len(res) > 0 {
			rr := models.Result{
				Num:          i + 1,
				Line:         s,
				Matches:      slices.String{},
				ExpressionID: x.ID,
			}
			for _, r := range res {
				if len(r) > 1 {
					rr.Matches = append(rr.Matches, r[1])
				}
			}
			if len(rr.Matches) > 0 {
				err = tx.Create(&rr)
				if err != nil {
					return errors.WithStack(err)
				}
				results = append(results, rr)
			}
		}
	}

	c.Set("expression", x)
	c.Set("results", results)

	return c.Render(200, r.Template("text/html", "checker/_results.html"))
}

func ReRunChecker(c buffalo.Context) error {
	x := &models.Expression{}
	tx := c.Value("tx").(*pop.Connection)
	err := tx.Find(x, c.Param("expression_id"))
	if err != nil {
		return errors.WithStack(err)
	}

	results := &models.Results{}
	err = tx.Where("expression_id = ?", x.ID).Order("num asc").All(results)
	if err != nil {
		return errors.WithStack(err)
	}

	c.Set("expression", x)
	c.Set("results", results)
	return c.Render(200, r.HTML("checker/new.html"))
}

const exp = `(Go|start|buffalo)`
const testString = `Welcome to Gobular!!

This is an online regular expression tester for Go, sometimes known as Golang.

All you need to do is to start typing an expression and set up your own test string
and Gobular will do the rest!

This project is powered by http://gobuffalo.io. We hope you enjoy.
`
