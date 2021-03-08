package http

import (
	"fmt"
	"net/http"
)


type RouterHandler interface {
	API() (*fury.Application, error)
	RouteURLs(app *fury.Application)
}

func Response(w http.ResponseWriter, method string, body interface{}, err error) error {
	if err != nil {
		return err
	}
	if method == http.MethodPost {
		return web.RespondJSON(w, body, http.StatusCreated)
	}
	return web.RespondJSON(w, body, http.StatusOK)
}

// API constructs an http.Handler with all application routes defined.
func API(routesHandler RouterHandler) (*fury.Application, error) {
	app, err := fury.NewWebApplication(fury.WithLogLevel(log.DebugLevel), fury.WithErrorHandler(middleware.ErrorHandler))
	if err == nil {
		routesHandler.RouteURLs(app)
	}

	return app, err
}


type MutantService interface {
	IsMutant(dnaRequest model.IsMutantRequestBody) bool
	GetMutantStats() (*model.Stats, *errors.ApiErrorImpl)
}

type MutantHandler struct {
	service interfaces.MutantService
}

func NewMutantController(service interfaces.MutantService) MutantHandler {
	return MutantHandler{service: service}
}

func (i MutantHandler) IsMutantHandler(ctx *gin.Context) {
	var json IsMutantRequestBody
	if err := ctx.BindJSON(&json); err != nil {
		apiErr := errors.BadRequestError(err)
		ctx.JSON(apiErr.Code, apiErr)
	}
	if valid, message := json.IsValid(); !valid {
		apiErr := errors.BadRequestError(fmt.Errorf(message))
		ctx.JSON(apiErr.Code, apiErr)
	}
	is := i.service.IsMutant(json)
	if is {
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(http.StatusForbidden)
	}
}

func (i MutantHandler) GetStatsHandler(ctx *gin.Context) {
	stats, apiErr := i.service.GetMutantStats()
	if apiErr != nil {
		ctx.JSON(apiErr.Code, apiErr)
	}
	ctx.JSON(http.StatusOK, stats)
}

import "fmt"

const InvalidNitrogenBaseFoundMessage = "invalid nitrogen base found: %v"
const InvalidInputMatrixToShortMessage = "invalid input, the matrix is to short, has to be 4x4 or bigger"
const InvalidInputNotAnNxNMatrixMessage = "invalid input, it isn't a NxN matrix, this could cause an Internal Error"

type IsMutantRequestBody struct {
	Dna []string `form:"dna" json:"dna" binding:"required"`
}

var validDna = map[string]bool{
	"A": true,
	"T": true,
	"C": true,
	"G": true,
}

func (i IsMutantRequestBody) IsValid() (bool, string) {
	input := i.Dna
	size := len(input)
	if size < 4 {
		return false, InvalidInputMatrixToShortMessage
	}
	for _, v := range input {
		if size != len(v) {
			return false, InvalidInputNotAnNxNMatrixMessage
		}
	}
	for _, v := range input {
		for _, w := range v {
			word := string(w)
			if !validDna[word] {
				return false, fmt.Sprintf(InvalidNitrogenBaseFoundMessage, word)
			}
		}
	}
	return true, ""
}
