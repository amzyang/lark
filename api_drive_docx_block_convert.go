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

// ConvertDocxBlock 将 Markdown/HTML 等内容转换为文档块
//
// doc: https://open.feishu.cn/document/server-docs/docs/docs/docx-v1/document-block/convert
func (r *DriveService) ConvertDocxBlock(ctx context.Context, request *ConvertDocxBlockReq, options ...MethodOptionFunc) (*ConvertDocxBlockResp, *Response, error) {
	if r.cli.mock.mockDriveConvertDocxBlock != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#ConvertDocxBlock mock enable")
		return r.cli.mock.mockDriveConvertDocxBlock(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "ConvertDocxBlock",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/docx/v1/documents/blocks/convert",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(convertDocxBlockResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveConvertDocxBlock mock DriveConvertDocxBlock method
func (r *Mock) MockDriveConvertDocxBlock(f func(ctx context.Context, request *ConvertDocxBlockReq, options ...MethodOptionFunc) (*ConvertDocxBlockResp, *Response, error)) {
	r.mockDriveConvertDocxBlock = f
}

// UnMockDriveConvertDocxBlock un-mock DriveConvertDocxBlock method
func (r *Mock) UnMockDriveConvertDocxBlock() {
	r.mockDriveConvertDocxBlock = nil
}

// ConvertDocxBlockReq ...
type ConvertDocxBlockReq struct {
	ContentType string `json:"content_type"` // 内容类型, 如 "markdown"
	Content     string `json:"content"`      // 内容
}

// ConvertDocxBlockResp ...
type ConvertDocxBlockResp struct {
	FirstLevelBlockIDs []string    `json:"first_level_block_ids,omitempty"` // 一级块 ID 列表
	Blocks             []*DocxBlock `json:"blocks,omitempty"`               // 所有块
}

// convertDocxBlockResp ...
type convertDocxBlockResp struct {
	Code  int64                 `json:"code,omitempty"`
	Msg   string                `json:"msg,omitempty"`
	Data  *ConvertDocxBlockResp `json:"data,omitempty"`
	Error *ErrorDetail          `json:"error,omitempty"`
}
