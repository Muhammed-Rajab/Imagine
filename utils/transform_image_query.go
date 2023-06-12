package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	_ "image/jpeg"

	_ "image/png"
)

/*
A struct to store the query parameters from
url, validate and parse them to favourable dtypes.
*/
type TransformImageQuery struct {
	Rotate string `form:"rotate,omitempty"`
	Scale  string `form:"scale,omitempty"`
	Resize string `form:"resize,omitempty"`
}

func (q *TransformImageQuery) Validate() (bool, error) {

	// Validate Rotate
	if q.Rotate != "" {
		if _, err := strconv.ParseFloat(q.Rotate, 32); err != nil {
			return false, errors.New("invalid query value for rotate: the value should be a valid float between 0.0 and 360.0")
		}
	}

	// Validate Scale
	if q.Scale != "" {
		if scale, err := strconv.ParseFloat(q.Scale, 32); err != nil || scale > 5.0 || scale < 1.0 {
			return false, errors.New("invalid query value for scale: the value should be a valid float between 1.0 and 5.0")
		}
	}

	// Validate Resize
	if q.Resize != "" {

		splittedResize := strings.Split(q.Resize, "x")
		w, err1 := strconv.ParseUint(splittedResize[0], 10, 32)
		h, err2 := strconv.ParseUint(splittedResize[1], 10, 32)

		sizeIs2 := len(splittedResize) == 2
		resIsValid := w <= 5000 && h <= 5000

		if !sizeIs2 || err1 != nil || err2 != nil || !resIsValid {
			return false, fmt.Errorf("invalid query value for resize: the value %s is not valid. the maximum resize resolution is 5000x5000 and the minimum resize resolution is 0x0", q.Resize)
		}

	}

	return true, nil
}

func (q *TransformImageQuery) Parse() *TransformImageParsedQuery {

	parsed := new(TransformImageParsedQuery)

	// Parse rotate
	if q.Rotate != "" {
		parsed.Rotate, _ = strconv.ParseFloat(q.Rotate, 32)
	} else {
		parsed.Rotate = 0
	}

	// Parse scale
	if q.Scale != "" {
		parsed.Scale, _ = strconv.ParseFloat(q.Scale, 32)
	} else {
		parsed.Scale = 1
	}

	// Parse Resize
	if q.Resize != "" {
		splittedResize := strings.Split(q.Resize, "x")
		W, _ := strconv.ParseUint(splittedResize[0], 10, 32)
		H, _ := strconv.ParseUint(splittedResize[1], 10, 32)
		parsed.Resize[0] = int(W)
		parsed.Resize[1] = int(H)
	} else {
		parsed.Resize = [2]int{-1, -1}
	}

	return parsed
}

/*
A struct to store the parsed values from TransformImageQuery.
*/
type TransformImageParsedQuery struct {
	Rotate float64
	Scale  float64
	Resize [2]int
}
