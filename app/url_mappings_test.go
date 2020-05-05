package app

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestRouter(t *testing.T) {
	assert.NotNil(t, router)

	mapUrls()

	assert.NotNil(t, router)

	var routes []gin.RouteInfo
	for _, r := range router.Routes() {
		if strings.Contains(r.Path, "/volunteer") {
			routes = append(routes, r)
		}
	}

	assert.EqualValues(t, 9, len(routes))

	assert.EqualValues(t, http.MethodGet, routes[0].Method)
	assert.EqualValues(t, "/volunteer/get", routes[0].Path)

	assert.EqualValues(t, http.MethodGet, routes[1].Method)
	assert.EqualValues(t, "/volunteer/get/:id", routes[1].Path)

	assert.EqualValues(t, http.MethodGet, routes[2].Method)
	assert.EqualValues(t, "/volunteer/all", routes[2].Path)

	assert.EqualValues(t, http.MethodPost, routes[3].Method)
	assert.EqualValues(t, "/volunteer/create", routes[3].Path)

	assert.EqualValues(t, http.MethodPost, routes[4].Method)
	assert.EqualValues(t, "/volunteer/import", routes[4].Path)

	assert.EqualValues(t, http.MethodPost, routes[5].Method)
	assert.EqualValues(t, "/volunteer/auth", routes[5].Path)

	assert.EqualValues(t, http.MethodPut, routes[6].Method)
	assert.EqualValues(t, "/volunteer/update/:id", routes[6].Path)

	assert.EqualValues(t, http.MethodPut, routes[7].Method)
	assert.EqualValues(t, "/volunteer/auth/update", routes[7].Path)

	assert.EqualValues(t, http.MethodDelete, routes[8].Method)
	assert.EqualValues(t, "/volunteer/delete/:id", routes[8].Path)
}
