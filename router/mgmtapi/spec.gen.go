// Package mgmtapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package mgmtapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9xZ33PbNhL+VzBoH5op9cNOem305thOq5kk1lj29KH1eSBiRaIGARYAZet8+t9vFiAp",
	"/rLs5K7tXJ9skQD2w7ffLnbBRxrrLNcKlLN09kgN2FwrC/7HO8Yv4fcCrMNfsVYOlP+X5bkUMXNCq8lv",
	"Vit8ZuMUMob/fW1gTWf0q8l+6Ul4aydLxxRnhp8bow3d7XYR5WBjI3JcjM7QJjGlUXxbTvRw3p/hn9zo",
	"HIwTASMHKwzw20wokRXZrXu4FcqB2TBZvm4sfpUCKQeSahRZgbsHUMQZpmwmrBVaEb0m796fEdyz0ZLk",
	"LL4DZ4lLmSMuBYIQmNOGBPt2TK5SYcmGyQKIsITxDWK0wInTfkYOYCKS6nvYgPFPWOwKJvdAChwtLLE5",
	"xGItgJPVljh2J1Tix2fswSPX69IqH5WbGbmHUb0MU9wPD1j02v8wkGkHntnWRAMxiA3sQfhZYxpReGBZ",
	"LoHO6PF0mlkaUbfN8ad1RqiEes85iJHa26yQTuRSgBkmXRXZCgyCaTGZFdaRFfrElkxxiCUzQByyaSE4",
	"g1nC9b1CjoHURveY1zoQih6r5ghLYibjQjIXiCwhbis2W/QoSLQTfmhLBnuRbAOkPj2va2JwcAIGmQHF",
	"VhJ4n4y54mXgoOn7FFwKxgMXlpSzvAdjrdYiKQxwolWw7cGsWdy270wBNYSV1hKYQgiVq+vIKF39mVFR",
	"zuKHwgFdtbUOMmJTXUhObJHn2rjng6KUJcYGPhKBHWjJfe3TgYq35BsxhnHUxjoKWGrgr2rkTwJGJHEM",
	"uUO2KyRSx0yW23iR/BsU09kvB/PQE5Gyl8kBb91E1AnngbwTXJiwDJPkvTb3zHCU81kdEpVqaoUx1ZZN",
	"uQm9+g1ihzKZV2/7qXW15s+lc0zJu4jWJm7FgOaXp/OLT3sYRHBQDhOceT6Q/CzF5K1o4uzrl3FuwFrc",
	"cjWFdO1WqVAXrmOaHr09Hh/944fx8fh49vpoOp0OpTsFIklX2jxHSk3pp2qC14r0TrGpyJ9b4INQd5fN",
	"8f4c9OpxxbMnLA78eHXtJznm4CXWln5gV9Mttzb230QTeZlUpjr7HPRfQ9HBQ/O9h1TDQwfV+qnhi7Zq",
	"SyX0ZXJ9tpjMF6RQHIxk26Zk0GgHyxfoQ1h+GyqVg+qw/MT2qQ5zoxp+g6VqrxjqXU2D4rkWylW7kELd",
	"HY5ze1mWeH3q6mXDLweZfbHW0UZplBnDtl5+YiWFSm6/YN1lmHpg+d2eoGpHRArrkKWQzPEYLSGQBoQh",
	"crxPZo9Nj4/W6+l0Np0dHaGzc+ZQyHRG//nrr/zb0Te/sNF6Onp783gUvdnNXj0e79qPXv0bx31N9yjn",
	"y7PRyZLM6+w3pKFe6CMoVWSokdOLy3Ma0dOf5h/OaEQXJ5fnn67wn/PzS9TLHnw1ZHD5ZZUUqnWvFzSi",
	"Zxc/f2ov4h/3V9DJB9iA7KtHVo/bYfdBJ4n3iX8d1VY5rIrEZ4i1xse+IWgBKN8cPnfDsjcDTl0YvZKQ",
	"DbUMjokBpCckLTKGJQ/jvjSAh1wyFc7SsiiPQ70gLNFxXBgDan+w5MFgXWSkIPN1IXEGCrIsa6pRqM4E",
	"S2/GNyLkvlTf4+Dc6BiAj8nPRjgHeIKTc5VIYVM/q8aHdS+oRCgAYyNS2IJJuSVKO2ILgcUsjlCYVSFO",
	"lfAVjmN3kGrJwVi/Go728SL+Bbyd9U61UmVhgaU5c2zFLBAnMqxKCzeYBJV1TA0d0yfk+nJODKwhsBZo",
	"qqLBenJqlp9kNyIwTsZYjzPuix9G1oYlGajGYoZoQ2yxGuXMpXUDVrlnm8OYfGRb7DyKshhtOMhoXaZT",
	"YetJIpxMVhcmBhJr3jkgJuXASVxzNvKS/srpO1Aj1PIIHTfy7I0Ce2ttMubojBZGjGpmhmjF47Www7XP",
	"T1dXCxIGeGQkAYXdadlAYrdqRCIUsWCw9wzt0iEJt/b23fR1RMtinM6+e/s2omWRSmdH0+lQ1VamvL4C",
	"bKoNijPLmNn24sY75q8W/RKMj8drxTZMSLQ55JDwAHe4ZoVEH7KVLtxsJZm6o9FLtF8o8XsBctsNgiYf",
	"RCscENTnb2AeXIO3jeDAycliPiYXea4bnVUVSaxslcnl+9PR9z9Mv4+I8NlJgfC9p4FYZxkoHuausMOu",
	"gHrCka9QYzhNWMiRo9odXMcFBl+wo7QhidQr75Kwv7o7b7n5ZcHzGSHSORbKeKmkOHQ+1IXycENcdp+t",
	"64BCIXeKrLYOrN9YqMfK9rLsdw3kBiyWM9W+nY619Ak0LPHN4uz6VbvwlGwLxnMtbC3qxg0GszWkc/Sb",
	"AkdytpWacTIi8wX5CRgHQ0bk+qz60WL56M33x0Ox2qu0ni4L/5Lubl61c516PdR4f3gzV9LzN2vlBoh/",
	"sr/rdHQBSLOJK0vs+cESu8tjX2X/fff0v+6Z2pfVPcRQPW4L1o8mGVjLkucTVV33dqzvdmVp3D9FF/M6",
	"p4atXdb9ctUQ+QekOspOFnMa0Q0YG1aYjqfjI9ygzkGxXNAZfT2ejo9Dn5P6zU3CVRL+m4C/9A9X3kKr",
	"Oacz+iO40zAian80OJ5OO18L8Mya5JKJzneCLjG9bwHLIo7BWqyhLyrjCPtNMDGkkxrKpPHxwn9HCDUH",
	"ndGFEVVqvrr4+KFzZ7YWMlyUscSif/Bw1Ire4BqTyiFPMTIPHcv/Fx/vmBUxwa3hSYsc5CwB4suZuuww",
	"WvryEeXk+xNrD7DUbPdLrjpNobCuIeD9jFAZMQO9C/Dmxd0A8Y3c8wz9X/7xauAOZcBLfm963dvauOGq",
	"J+CUddC3nweranQHsMzVhklRf1Ebd1z/pBsarm1c3HnvSp1M6lb/qUCobwn+QG/UNv60SPkRsGtoX2f0",
	"IiCieTFAyrJDil//nebbP4WP6hKmaT+cQM4UsPtbeWn5Ei/5Kb4dxuePtDCSzmjqXD6bTB5Tbd1u9phr",
	"43YTlovJ5ggPUGYE9jyeIxzS7v98Qeofowa06bx+PX3z5hhZuKnh9CqHDZitSxG4L7pCR9bPIxFVLKtK",
	"6ep6tLvYqd8qnv0EHkJruNqWi5WZvLlUyczuZvefAAAA//9UPMQGEyAAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
