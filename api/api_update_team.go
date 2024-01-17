package api

import (
  "TeamManagementSystem/utils"
  "github.com/gofiber/fiber/v2"

  "TeamManagementSystem/dto"
    "github.com/go-playground/validator/v10"
  "TeamManagementSystem/dao"
  )

// @Summary      PUT Team input: Team
// @Description  PUT Team API
// @Tags         PUT Team - Team
// @Accept       json
// @Produce      json
// @Param        data body dto.Model_Team false "string collection" 
// @Success      200  {array}   dto.Model_Team "Status OK"
// @Success      202  {array}   dto.Model_Team "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateTeam [PUT]

    func UpdateTeamApi(c *fiber.Ctx) error {



    inputObj := dto.Team{}
    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateTeam(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}