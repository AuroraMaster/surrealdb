// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"github.com/abcum/fibre"
	"github.com/abcum/surreal/kvs"
	"github.com/abcum/surreal/sql"
)

func errors(err error, c *fibre.Context) {

	var code int
	var text string
	var info map[string]interface{}

	switch e := err.(type) {
	default:
		code = 400
	case *kvs.DBError:
		code, text = 503, e.Error()
	case *kvs.TXError:
		code, text = 500, e.Error()
	case *kvs.KVError:
		code, text = 409, e.Error()
	case *kvs.CKError:
		code, text = 403, e.Error()
	case *sql.ParseError:
		code, text = 400, e.Error()
	case *fibre.HTTPError:
		code = e.Code()
	}

	if _, ok := errs[code]; !ok {
		code = 500
	}

	info = errs[code]
	if text != "" {
		info["information"] = text
	}

	c.Send(code, info)

}

var errs = map[int]map[string]interface{}{

	200: map[string]interface{}{
		"code":          200,
		"details":       "Information",
		"documentation": "https://docs.surreal.io/",
		"information":   "Visit the documentation for details on accessing the api.",
	},

	400: map[string]interface{}{
		"code":          400,
		"details":       "Request problems detected",
		"documentation": "https://docs.surreal.io/",
		"information":   "There is a problem with your request. Ensure that the request is valid.",
	},

	401: map[string]interface{}{
		"code":          401,
		"details":       "Authentication failed",
		"documentation": "https://docs.surreal.io/",
		"information":   "Your authentication details are invalid. Reauthenticate using a valid token.",
	},

	403: map[string]interface{}{
		"code":          403,
		"details":       "Request resource forbidden",
		"documentation": "https://docs.surreal.io/",
		"information":   "Your request was forbidden. Perhaps you don't have the necessary permissions to access this resource.",
	},

	404: map[string]interface{}{
		"code":          404,
		"details":       "Request resource not found",
		"documentation": "https://docs.surreal.io/",
		"information":   "The requested resource does not exist. Check that you have entered the url correctly.",
	},

	405: map[string]interface{}{
		"code":          405,
		"details":       "This method is not allowed",
		"documentation": "https://docs.surreal.io/",
		"information":   "The requested http method is not allowed for this resource. Refer to the documentation for allowed methods.",
	},

	409: map[string]interface{}{
		"code":          409,
		"details":       "Request conflict detected",
		"documentation": "https://docs.surreal.io/",
		"information":   "The request could not be processed because of a conflict in the request.",
	},

	413: map[string]interface{}{
		"code":          413,
		"details":       "Request content length too large",
		"documentation": "https://docs.surreal.io/",
		"information":   "All requests to the database must not exceed the predefined content length.",
	},

	415: map[string]interface{}{
		"code":          415,
		"details":       "Unsupported content type requested",
		"documentation": "https://docs.surreal.io/",
		"information":   "The request needs to adhere to certain constraints. Check your request settings and try again.",
	},

	422: map[string]interface{}{
		"code":          422,
		"details":       "Request problems detected",
		"documentation": "https://docs.surreal.io/",
		"information":   "There is a problem with your request. The request appears to contain invalid data.",
	},

	426: map[string]interface{}{
		"code":          426,
		"details":       "Upgrade required",
		"documentation": "https://docs.surreal.io/",
		"information":   "There is a problem with your request. The request is expected to upgrade to a websocket connection.",
	},

	500: map[string]interface{}{
		"code":          500,
		"details":       "There was a problem with our servers, and we have been notified",
		"documentation": "https://docs.surreal.io/",
	},
}
