// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package client

// JSONError - Struct for standard error returns over REST API
type JSONError struct {
	Code    int    `json:"code,omitempty"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

// JSONResponse - Top level container of a REST API response
type JSONResponse struct {
	Data  interface{} `json:"data"`
	Error JSONError   `json:"error,omitempty"`
}

// EntityResponse - Response from the API for an Entity request
type EntityResponse struct {
	Data  Entity    `bson:"data" json:"data"`
	Error JSONError `json:"error,omitempty"`
}

// CollectionResponse - Response from the API for an Collection request
type CollectionResponse struct {
	Data  Collection `bson:"data" json:"data"`
	Error JSONError  `json:"error,omitempty"`
}

// ContainerResponse - Response from the API for an Container request
type ContainerResponse struct {
	Data  Container `bson:"data" json:"data"`
	Error JSONError `json:"error,omitempty"`
}

// ImageResponse - Response from the API for an Image request
type ImageResponse struct {
	Data  Image     `bson:"data" json:"data"`
	Error JSONError `json:"error,omitempty"`
}

// TagsResponse - Response from the API for a tags request
type TagsResponse struct {
	Data  TagMap    `bson:"data" json:"data"`
	Error JSONError `json:"error,omitempty"`
}
