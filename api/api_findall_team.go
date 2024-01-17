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
// @Router      /FindallTeam [GET]

    func FindallTeamApi(c *fiber.Ctx) error {


returnValue, err := dao.DB_FindallTeam()
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}