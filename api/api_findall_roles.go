package api

import (
  "TeamManagementSystem/utils"
  "github.com/gofiber/fiber/v2"

  "TeamManagementSystem/dao"
  )

// @Summary      GET Roles input: Roles
// @Description  GET Roles API
// @Tags         GET Roles - Roles
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Roles "Status OK"
// @Success      202  {array}   dto.Model_Roles "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindallRoles [GET]

    func FindallRolesApi(c *fiber.Ctx) error {


returnValue, err := dao.DB_FindallRoles()
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}