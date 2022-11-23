package datatypes

import (
	// "acquia/decision-service/pkg/database"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type (
	Error struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}

	// ApiResponseErrorV3 version 1 api error response.
	ApiResponseErrorV3 struct {
		Errors []*Error `json:"errors"`
	}

	// ApiResponseSuccessV3 version 3 api success response.
	ApiResponseSuccessV3 struct {
		TotalCount int         `json:"total_count"`
		Total      int         `json:"total"`
		Pagination *Pagination `json:"pagination"`
		Data       interface{} `json:"data"`
	}

	/*
		Database related
	*/

	// Pagination represents the pagination metadata.
	Pagination struct {
		TotalRecords int `json:"total"`
		Limit        int `json:"limit"`
		Offset       int `json:"offset"`
	}
	// PaginationResult represents the paginated output with pagination metadata.
	PaginationResult struct {
		PaginationInfo Pagination  `json:"pagination"`
		Records        interface{} `json:"records"`
	}

	// AppDependencies Application dependency object.
	AppDependencies struct {
		// DecisionDB database.Database
		// CisDB      database.Database
		Logger *logrus.Entry
		Config *viper.Viper
	}
)
