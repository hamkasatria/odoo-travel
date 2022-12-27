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

type ResponseOnlyMeta struct {
	Meta Meta `json:"meta"`
}

type Meta struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func JsonResponse(ctx *fiber.Ctx, status int, data interface{}) {
	var ress interface{}
	if status == fiber.StatusOK || status == fiber.StatusCreated {
		meta := Meta{
			Status:  status,
			Message: "Success",
		}

		if data == nil {
			ress = ResponseOnlyMeta{
				Meta: meta,
			}
		} else {
			ress = Response{
				Meta: meta,
				Data: data,
			}
		}

	} else {
		ress = ResponseOnlyMeta{
			Meta: Meta{
				Status:  status,
				Message: fmt.Sprintf("%v", data),
			},
		}

	}
	ctx.Status(status).JSON(ress)

}
