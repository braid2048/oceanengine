package project

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// BudgetGroupDelete  批量删除预算组
func BudgetGroupDelete(clt *core.SDKClient, accessToken string, req *project.BudgetGroupDeleteRequest) (*project.BudgetGroupDeleteResult, error) {
	var resp project.BudgetGroupDeleteResponse
	if err := clt.Post("v3.0/budget_group/delete/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
