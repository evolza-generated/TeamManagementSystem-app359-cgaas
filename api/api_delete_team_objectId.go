package api

import (
  "TeamManagementSystem/utils"
  "github.com/gofiber/fiber/v2"

  "TeamManagementSystem/dao"
  )

// @Summary      DELETE Team input: Team
// @Description  DELETE Team API
// @Tags         DELETE Team - Team
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Team "Status OK"
// @Success      202  {array}   dto.Model_Team "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteTeam [DELETE]

    func DeleteTeamApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    err := dao.DB_DeleteTeam(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}