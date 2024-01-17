package api

import (
  "TeamManagementSystem/utils"
  "github.com/gofiber/fiber/v2"

  "TeamManagementSystem/dao"
  )

// @Summary      DELETE Member input: Member
// @Description  DELETE Member API
// @Tags         DELETE Member - Member
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Model_Member "Status OK"
// @Success      202  {array}   dto.Model_Member "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteMember [DELETE]

    func DeleteMemberApi(c *fiber.Ctx) error {


objectId := c.Query("objectId")
    err := dao.DB_DeleteMember(objectId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return utils.SendSuccessResponse(c)
    
}