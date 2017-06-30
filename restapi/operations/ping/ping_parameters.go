package ping

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPingParams creates a new PingParams object
// with the default values initialized.
func NewPingParams() PingParams {
	var ()
	return PingParams{}
}

// PingParams contains all the bound params for the ping operation
// typically these are obtained from a http.Request
//
// swagger:parameters ping
type PingParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Login
	  Required: true
	  In: query
	*/
	Login string
	/*Password
	  Required: true
	  In: query
	*/
	Password string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PingParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLogin, qhkLogin, _ := qs.GetOK("login")
	if err := o.bindLogin(qLogin, qhkLogin, route.Formats); err != nil {
		res = append(res, err)
	}

	qPassword, qhkPassword, _ := qs.GetOK("password")
	if err := o.bindPassword(qPassword, qhkPassword, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PingParams) bindLogin(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("login", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("login", "query", raw); err != nil {
		return err
	}

	o.Login = raw

	return nil
}

func (o *PingParams) bindPassword(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("password", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("password", "query", raw); err != nil {
		return err
	}

	o.Password = raw

	return nil
}
