package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3Blades/go-sdk/models"
)

// AuthJwtTokenVerifyCreateReader is a Reader for the AuthJwtTokenVerifyCreate structure.
type AuthJwtTokenVerifyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthJwtTokenVerifyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAuthJwtTokenVerifyCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAuthJwtTokenVerifyCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAuthJwtTokenVerifyCreateCreated creates a AuthJwtTokenVerifyCreateCreated with default headers values
func NewAuthJwtTokenVerifyCreateCreated() *AuthJwtTokenVerifyCreateCreated {
	return &AuthJwtTokenVerifyCreateCreated{}
}

/*AuthJwtTokenVerifyCreateCreated handles this case with default header values.

VerifyJSONWebToken created
*/
type AuthJwtTokenVerifyCreateCreated struct {
	Payload *models.VerifyJSONWebToken
}

func (o *AuthJwtTokenVerifyCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/auth/jwt-token-verify/][%d] authJwtTokenVerifyCreateCreated  %+v", 201, o.Payload)
}

func (o *AuthJwtTokenVerifyCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VerifyJSONWebToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthJwtTokenVerifyCreateBadRequest creates a AuthJwtTokenVerifyCreateBadRequest with default headers values
func NewAuthJwtTokenVerifyCreateBadRequest() *AuthJwtTokenVerifyCreateBadRequest {
	return &AuthJwtTokenVerifyCreateBadRequest{}
}

/*AuthJwtTokenVerifyCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type AuthJwtTokenVerifyCreateBadRequest struct {
}

func (o *AuthJwtTokenVerifyCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/auth/jwt-token-verify/][%d] authJwtTokenVerifyCreateBadRequest ", 400)
}

func (o *AuthJwtTokenVerifyCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*AuthJwtTokenVerifyCreateBody auth jwt token verify create body
swagger:model AuthJwtTokenVerifyCreateBody
*/
type AuthJwtTokenVerifyCreateBody struct {

	// token
	// Required: true
	Token *string `json:"token"`
}