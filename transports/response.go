package transports

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type GetListTravels struct {
	Count int         `json:"count"`
	List  interface{} `json:"list"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func JsonResponse(ctx *fiber.Ctx, status int, data interface{}) {
	var ress Response
	if status == fiber.StatusOK || status == fiber.StatusCreated {
		ress = Response{
			Meta: Meta{
				Status:  status,
				Message: "Success",
			},
			Data: data,
		}

	} else {
		ress = Response{
			Meta: Meta{
				Status:  status,
				Message: fmt.Sprintf("%v", data),
			},
			Data: nil,
		}
	}
	ctx.Status(status).JSON(ress)

}
