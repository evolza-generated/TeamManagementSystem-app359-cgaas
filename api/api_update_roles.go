package api

import (
  "TeamManagementSystem/utils"
  "github.com/gofiber/fiber/v2"

  "TeamManagementSystem/dto"
    "github.com/go-playground/validator/v10"
  "TeamManagementSystem/dao"
  )

// @Summary      PUT Roles input: Roles
// @Description  PUT Roles API
// @Tags         PUT Roles - Roles
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Roles false "string collection" 
// @Success      200  {array}   dto.Model_Roles "Status OK"
// @Success      202  {array}   dto.Model_Roles "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateRoles [PUT]

    func UpdateRolesApi(c *fiber.Ctx) error {



    inputObj := dto.Roles{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateRoles(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}