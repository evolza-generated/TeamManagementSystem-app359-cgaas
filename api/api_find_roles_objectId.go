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
// @Router      /FindRoles [GET]

    func FindRolesApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    returnValue, err := dao.DB_FindRolesbyObjectId(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}