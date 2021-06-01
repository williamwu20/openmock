package openmock

import (
	"fmt"
	"io/ioutil"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

var schema *graphql.Schema

func (om *OpenMock) startGraphQL() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		logrus.WithFields(logrus.Fields{
			"http_path":   c.Path(),
			"http_method": c.Request().Method,
			"http_host":   c.Request().Host,
			"http_req":    string(reqBody),
			"http_res":    string(resBody),
		}).Info()
	}))
	if om.CorsEnabled {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowCredentials: true,
			AllowHeaders:     []string{"*"},
			AllowMethods:     []string{"*"},
		}))
	}

	mocks := om.repo.GraphQLMocks

	for g, ms := range mocks {
		//TODO combine all schemas as a giant string and send it as parameter to MustParseSchema
		func(g ExpectGraphQL, ms MocksArray) {
			c := Context{
				GraphQLSchema: g.Schema,
				om:            om,
			}
			ms.DoActions(c)

		}(g, ms)
	}
	// schema := graphql.MustParseSchema("placeholder", &query{})
	// // graphql
	// h, err := graphql.NewHandler(db)
	// logFatal(err)
	// e.POST("/graphql", echo.WrapHandler(h))

	// mux := http.NewServeMux()

	// mux.Handle("/query", &relay.Handler{Schema: schema})

	// log.WithFields(log.Fields{"time": time.Now()}).Info("starting server")
	// log.Fatal(http.ListenAndServe("localhost:8080", logged(mux)))
	e.Logger.Fatal(
		e.Start(fmt.Sprintf("%s:%d", om.HTTPHost, om.HTTPPort)),
	)
}
