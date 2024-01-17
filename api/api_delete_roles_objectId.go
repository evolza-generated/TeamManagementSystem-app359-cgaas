package api

import (
  "TeamManagementSystem/utils"
  "github.com/gofiber/fiber/v2"

  "TeamManagementSystem/dao"
  )

// @Summary      DELETE Roles input: Roles
// @Description  DELETE Roles API
// @Tags         DELETE Roles - Roles
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Roles "Status OK"
// @Success      202  {array}   dto.Model_Roles "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteRoles [DELETE]

    func DeleteRolesApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    err := dao.DB_DeleteRoles(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}