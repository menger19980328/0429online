package go_server_application_Tenant

import (
	"context"
	"time"

	"github.com/byted-apaas/faas-sdk-go/faas"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/byted-apaas/server-sdk-go/application"
	// "github.com/byted-apaas/server-sdk-go/common/structs"
)

// /*Params 函数入参定义
//   - 结构体名称不支持自定义, 必须为 Params
//   - 结构体属性支持自定义, 和 meta.json 中的 input 参数一一对应
//     */
type Params struct{}

// /*Result 函数出参定义
//   - 结构体名称不支持自定义, 必须为 Result
//   - 结构体属性支持自定义, 和 meta.json 中的 output 参数一一对应
//     */
type Result struct {
	// Tenant *TenantInfo `json:"tenant"`
}

// type TenantInfo struct {
// 	ID int64 `json:"id"`
// }

/*Handler 函数入口
 * 入口函数签名不支持自定义, 必须为 func(context.Context, *Params) (*Result, error)
 * @param ctx    请求上下文参数, 使用 server-sdk-go 需要链路透该参数
 * @param params 请求参数
 */
func Handler(ctx context.Context, params *Params) (*Result, error) {
	// 日志功能
	application.GetLogger(ctx).Infof("%s 函数开始执行111, logid: %s", time.Now().Format("2006-01-02 15:04:05.999"), faas.Tool.GetLogID(ctx))

	// 获取租户信息
	tenant, err := application.Tenant.GetTenantInfo(ctx)
	if err != nil {
		panic(err)
	}
	application.GetLogger(ctx).Infof("tenant: %s", cUtils.ToString(tenant))
	return &Result{
		// Tenant:tenant,
	}, nil
}
