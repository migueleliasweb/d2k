// Code generated by go-swagger; DO NOT EDIT.

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// ImageListURL generates an URL for the image list operation
type ImageListURL struct {
	All     *bool
	Digests *bool
	Filters *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ImageListURL) WithBasePath(bp string) *ImageListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *ImageListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *ImageListURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/images/json"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1.41"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var allQ string
	if o.All != nil {
		allQ = swag.FormatBool(*o.All)
	}
	if allQ != "" {
		qs.Set("all", allQ)
	}

	var digestsQ string
	if o.Digests != nil {
		digestsQ = swag.FormatBool(*o.Digests)
	}
	if digestsQ != "" {
		qs.Set("digests", digestsQ)
	}

	var filtersQ string
	if o.Filters != nil {
		filtersQ = *o.Filters
	}
	if filtersQ != "" {
		qs.Set("filters", filtersQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *ImageListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *ImageListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *ImageListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on ImageListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on ImageListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *ImageListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
