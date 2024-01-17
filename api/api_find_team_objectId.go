package api

import (
  "TeamManagementSystem/utils"
  "github.com/gofiber/fiber/v2"

  "TeamManagementSystem/dao"
  )

// @Summary      GET Team input: Team
// @Description  GET Team API
// @Tags         GET Team - Team
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Team "Status OK"
// @Success      202  {array}   dto.Model_Team "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindTeam [GET]

    func FindTeamApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    returnValue, err := dao.DB_FindTeambyObjectId(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}