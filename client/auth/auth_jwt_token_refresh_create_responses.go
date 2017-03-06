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

// AuthJwtTokenRefreshCreateReader is a Reader for the AuthJwtTokenRefreshCreate structure.
type AuthJwtTokenRefreshCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthJwtTokenRefreshCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAuthJwtTokenRefreshCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAuthJwtTokenRefreshCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAuthJwtTokenRefreshCreateCreated creates a AuthJwtTokenRefreshCreateCreated with default headers values
func NewAuthJwtTokenRefreshCreateCreated() *AuthJwtTokenRefreshCreateCreated {
	return &AuthJwtTokenRefreshCreateCreated{}
}

/*AuthJwtTokenRefreshCreateCreated handles this case with default header values.

RefreshJSONWebToken created
*/
type AuthJwtTokenRefreshCreateCreated struct {
	Payload *models.RefreshJSONWebToken
}

func (o *AuthJwtTokenRefreshCreateCreated) Error() string {
	return fmt.Sprintf("[POST /api/v0/auth/jwt-token-refresh/][%d] authJwtTokenRefreshCreateCreated  %+v", 201, o.Payload)
}

func (o *AuthJwtTokenRefreshCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RefreshJSONWebToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthJwtTokenRefreshCreateBadRequest creates a AuthJwtTokenRefreshCreateBadRequest with default headers values
func NewAuthJwtTokenRefreshCreateBadRequest() *AuthJwtTokenRefreshCreateBadRequest {
	return &AuthJwtTokenRefreshCreateBadRequest{}
}

/*AuthJwtTokenRefreshCreateBadRequest handles this case with default header values.

Invalid data supplied
*/
type AuthJwtTokenRefreshCreateBadRequest struct {
}

func (o *AuthJwtTokenRefreshCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v0/auth/jwt-token-refresh/][%d] authJwtTokenRefreshCreateBadRequest ", 400)
}

func (o *AuthJwtTokenRefreshCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*AuthJwtTokenRefreshCreateBody auth jwt token refresh create body
swagger:model AuthJwtTokenRefreshCreateBody
*/
type AuthJwtTokenRefreshCreateBody struct {

	// token
	// Required: true
	Token *string `json:"token"`
}
