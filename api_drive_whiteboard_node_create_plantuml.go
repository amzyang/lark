/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// CreateWhiteboardPlantUML 向画板中添加 PlantUML 节点
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/board-v1/whiteboard-node/plantuml
func (r *DriveService) CreateWhiteboardPlantUML(ctx context.Context, request *CreateWhiteboardPlantUMLReq, options ...MethodOptionFunc) (*CreateWhiteboardPlantUMLResp, *Response, error) {
	if r.cli.mock.mockDriveCreateWhiteboardPlantUML != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#CreateWhiteboardPlantUML mock enable")
		return r.cli.mock.mockDriveCreateWhiteboardPlantUML(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateWhiteboardPlantUML",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/board/v1/whiteboards/:whiteboard_id/nodes/plantuml",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createWhiteboardPlantUMLResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateWhiteboardPlantUML mock DriveCreateWhiteboardPlantUML method
func (r *Mock) MockDriveCreateWhiteboardPlantUML(f func(ctx context.Context, request *CreateWhiteboardPlantUMLReq, options ...MethodOptionFunc) (*CreateWhiteboardPlantUMLResp, *Response, error)) {
	r.mockDriveCreateWhiteboardPlantUML = f
}

// UnMockDriveCreateWhiteboardPlantUML un-mock DriveCreateWhiteboardPlantUML method
func (r *Mock) UnMockDriveCreateWhiteboardPlantUML() {
	r.mockDriveCreateWhiteboardPlantUML = nil
}

// CreateWhiteboardPlantUMLReq ...
type CreateWhiteboardPlantUMLReq struct {
	WhiteboardID string `path:"whiteboard_id" json:"-"`  // 画板唯一标识
	PlantUMLCode string `json:"plant_uml_code"`          // PlantUML 源码
	StyleType    int64  `json:"style_type"`              // 样式类型, 2: 经典样式（可二次编辑）
	SyntaxType   int64  `json:"syntax_type"`             // 语法类型, 1: PlantUML
}

// CreateWhiteboardPlantUMLResp ...
type CreateWhiteboardPlantUMLResp struct{}

// createWhiteboardPlantUMLResp ...
type createWhiteboardPlantUMLResp struct {
	Code  int64                         `json:"code,omitempty"`
	Msg   string                        `json:"msg,omitempty"`
	Data  *CreateWhiteboardPlantUMLResp `json:"data,omitempty"`
	Error *ErrorDetail                  `json:"error,omitempty"`
}
