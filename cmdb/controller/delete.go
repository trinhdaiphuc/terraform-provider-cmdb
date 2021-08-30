package controller

import (
	"github.com/whatvn/denny"
	"terraform-provider-cmdb/cmdb/model"
)

func DeleteConfig(ctx *denny.Context) {
	name, ok := ctx.GetQuery("name")
	if !ok {
		ctx.Status(400)
		return
	}

	model.DeleteAllocatedConfig(name)
	ctx.Status(200)
}
