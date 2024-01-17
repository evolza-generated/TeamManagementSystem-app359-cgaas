package api

import (
  "TeamManagementSystem/utils"
  "github.com/gofiber/fiber/v2"

  "TeamManagementSystem/dto"
    "github.com/go-playground/validator/v10"
  "TeamManagementSystem/dao"
  )

// @Summary      PUT Member input: Member
// @Description  PUT Member API
// @Tags         PUT Member - Member
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Member false "string collection" 
// @Success      200  {array}   dto.Model_Member "Status OK"
// @Success      202  {array}   dto.Model_Member "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateMember [PUT]

    func UpdateMemberApi(c *fiber.Ctx) error {



    inputObj := dto.Member{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateMember(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}