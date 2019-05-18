package main

import (
	"fmt"
	"net/http"

	"appScrip/apis/authapi"
	"appScrip/apis/awardsapi"
	"appScrip/apis/blogapi"
	"appScrip/apis/categoryapi"
	"appScrip/apis/contactusapi"
	"appScrip/apis/portfolioapi"
	"appScrip/apis/videosapi"

	"appScrip/app"

	"appScrip/daos/about"
	"appScrip/daos/awards"
	"appScrip/daos/blog"
	"appScrip/daos/category"
	"appScrip/daos/contactus"
	"appScrip/daos/portfolio"
	"appScrip/daos/service"
	"appScrip/daos/testimonial"
	"appScrip/daos/user"
	"appScrip/daos/videos"

	"appScrip/errors"

	"appScrip/services/serviceabout"
	"appScrip/services/serviceawards"
	"appScrip/services/serviceblog"
	"appScrip/services/servicecategory"
	"appScrip/services/servicecontactus"
	"appScrip/services/serviceportfolio"
	"appScrip/services/serviceservice"
	"appScrip/services/servicetestimonial"
	"appScrip/services/serviceuser"
	"appScrip/services/servicevideos"

	"github.com/Sirupsen/logrus"
	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
	_ "github.com/lib/pq"

	"gopkg.in/mgo.v2"
)

func main() {
	// load application configurations
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// load error messages
	if err := errors.LoadMessages(app.Config.ErrorFile); err != nil {
		panic(fmt.Errorf("Failed to read the error message file: %s", err))
	}

	// create the logger
	logger := logrus.New()

	db, err := mgo.Dial(app.Config.DSN) //dbx.MustOpen("postgres", app.Config.DSN)
	if err != nil {
		panic(err)
	}
	// db.LogFunc = logger.Infof

	// wire up API routing
	http.Handle("/", buildRouter(logger, db))

	// start the server
	address := fmt.Sprintf(":%v", app.Config.ServerPort)
	logger.Infof("server %v is started at %v", app.Version, address)
	panic(http.ListenAndServe(address, nil))
}

func buildRouter(logger *logrus.Logger, db *mgo.Session) *routing.Router {
	router := routing.New()

	router.To("GET,HEAD", "/ping", func(c *routing.Context) error {
		c.Abort() // skip all other middlewares/handlers
		return c.Write("OK " + app.Version)
	})

	router.Use(
		app.Init(logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.Options{
			AllowOrigins: "*",
			AllowHeaders: "*",
			AllowMethods: "*",
		}),
		app.Transactional(db),
	)

	rg := router.Group("")

	//make all daos
	userDAO := user.NewDAO()
	aboutDAO := about.NewDAO()
	categoryDAO := category.NewDAO()
	awardsDAO := awards.NewDAO()
	portfolioDAO := portfolio.NewDAO()
	blogDAO := blog.NewDAO()
	videosDAO := videos.NewDAO()
	contactusDAO := contactus.NewDAO()
	testimonialDAO := testimonial.NewDAO()
	serviceDAO := service.NewDAO()

	//intiate router for api without auth
	authapi.ServerAuthResource(rg, serviceuser.NewService(userDAO))
	categoryapi.ServeCategoryResource(
		rg,
		servicecategory.NewService(categoryDAO),
		serviceawards.NewService(awardsDAO),
		serviceabout.NewService(aboutDAO),
		servicetestimonial.NewService(testimonialDAO),
		serviceservice.NewService(serviceDAO),
	)
	awardsapi.ServeAwardsResource(rg, serviceawards.NewService(awardsDAO))
	portfolioapi.ServePortfolioResource(rg, serviceportfolio.NewService(portfolioDAO))
	blogapi.ServeBlogResource(rg, serviceblog.NewService(blogDAO))
	videosapi.ServeVideosResource(rg, servicevideos.NewService(videosDAO))
	contactusapi.ServeContactUsResource(rg, servicecontactus.NewService(contactusDAO))

	//intiate router for api with auth
	categoryapi.ServeCategoryResourceWithAuth(
		rg,
		servicecategory.NewService(categoryDAO),
		serviceawards.NewService(awardsDAO),
		serviceabout.NewService(aboutDAO),
		servicetestimonial.NewService(testimonialDAO),
		serviceservice.NewService(serviceDAO),
	)
	awardsapi.ServeAwardsResourceWithAuth(rg, serviceawards.NewService(awardsDAO))
	portfolioapi.ServePortfolioResourceWithAuth(rg, serviceportfolio.NewService(portfolioDAO))
	blogapi.ServeBlogResourceWithAuth(rg, serviceblog.NewService(blogDAO))
	videosapi.ServeVideosResourceWithAuth(rg, servicevideos.NewService(videosDAO))
	contactusapi.ServeContactUsResourceWithAuth(rg, servicecontactus.NewService(contactusDAO))

	// wire up more resource APIs here

	return router
}
