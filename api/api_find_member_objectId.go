package api

import (
  "TeamManagementSystem/utils"
  "github.com/gofiber/fiber/v2"

  "TeamManagementSystem/dao"
  )

// @Summary      GET Member input: Member
// @Description  GET Member API
// @Tags         GET Member - Member
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Member "Status OK"
// @Success      202  {array}   dto.Model_Member "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindMember [GET]

    func FindMemberApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    returnValue, err := dao.DB_FindMemberbyObjectId(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}