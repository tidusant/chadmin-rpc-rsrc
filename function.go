package main

import (
	"github.com/tidusant/c3m-common/c3mcommon"
	"github.com/tidusant/c3m-common/log"
	rpch "github.com/tidusant/chadmin-repo/cuahang"
	"github.com/tidusant/chadmin-repo/models"
	//	"c3m/common/inflect"
	//	"c3m/log"
)

type ViewData struct {
	Key   string
	Value map[string]string
}

func LoadAll(usex models.UserSession) models.RequestResult {
	request := "gettemplateresource|" + usex.Session
	resp := c3mcommon.RequestBuildService(request, "POST", usex.Shop.Theme)

	if resp.Status != "1" {
		return resp
	}

	return c3mcommon.ReturnJsonMessage("1", "", "success", string(resp.Data))

}
func Save(usex models.UserSession) models.RequestResult {

	str := `{"Code":"` + usex.Shop.Theme + `","Resources":` + usex.Params + `}`
	log.Debugf("save resource:%s", str)
	request := "savetemplateresource|" + usex.Session
	resp := c3mcommon.RequestBuildService(request, "POST", str)

	if resp.Status != "1" {
		return resp
	}

	rpch.CreateCommonDataBuild(usex)

	return c3mcommon.ReturnJsonMessage("1", "", "success", string(resp.Data))

	// var items []models.Resource
	// log.Debugf("Unmarshal %s", usex.Params)
	// err := json.Unmarshal([]byte(usex.Params), &items)
	// if !c3mcommon.CheckError("json parse resource", err) {
	// 	return c3mcommon.ReturnJsonMessage("0", "json parse fail", "", "")
	// }

	// //update
	// for _, saveitem := range items {
	// 	olditem := rpch.GetResourceByKey(usex.Shop.Theme, usex.Shop.ID.Hex(), saveitem.Key)
	// 	if olditem.Key == "" {
	// 		continue
	// 	}
	// 	olditem.Value = saveitem.Value
	// 	rpch.SaveResource(olditem)
	// }
	// //check olditem

	// var bs models.BuildScript
	// //build script
	// bs.ShopId = usex.Shop.ID.Hex()
	// bs.TemplateCode = usex.Shop.Theme
	// bs.Object = "common"
	// rpb.CreateBuild(bs)
	// return c3mcommon.ReturnJsonMessage("1", "", "success", "")
}
