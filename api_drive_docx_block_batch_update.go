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

// BatchUpdateDocxBlock 批量更新文档块内容。每次最多 200 个块，每个 block_id 只能出现一次。
//
// doc: https://open.feishu.cn/document/server-docs/docs/docs/docx-v1/document-block/batch_update
func (r *DriveService) BatchUpdateDocxBlock(ctx context.Context, request *BatchUpdateDocxBlockReq, options ...MethodOptionFunc) (*BatchUpdateDocxBlockResp, *Response, error) {
	if r.cli.mock.mockDriveBatchUpdateDocxBlock != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#BatchUpdateDocxBlock mock enable")
		return r.cli.mock.mockDriveBatchUpdateDocxBlock(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "BatchUpdateDocxBlock",
		Method:                "PATCH",
		URL:                   r.cli.openBaseURL + "/open-apis/docx/v1/documents/:document_id/blocks/batch_update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(batchUpdateDocxBlockResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveBatchUpdateDocxBlock mock DriveBatchUpdateDocxBlock method
func (r *Mock) MockDriveBatchUpdateDocxBlock(f func(ctx context.Context, request *BatchUpdateDocxBlockReq, options ...MethodOptionFunc) (*BatchUpdateDocxBlockResp, *Response, error)) {
	r.mockDriveBatchUpdateDocxBlock = f
}

// UnMockDriveBatchUpdateDocxBlock un-mock DriveBatchUpdateDocxBlock method
func (r *Mock) UnMockDriveBatchUpdateDocxBlock() {
	r.mockDriveBatchUpdateDocxBlock = nil
}

// BatchUpdateDocxBlockReq ...
type BatchUpdateDocxBlockReq struct {
	DocumentID         string                         `path:"document_id" json:"-"`           // 文档唯一标识
	DocumentRevisionID *int64                         `query:"document_revision_id" json:"-"` // 文档版本, -1 表示最新版本
	ClientToken        *string                        `query:"client_token" json:"-"`         // 幂等标识
	UserIDType         *IDType                        `query:"user_id_type" json:"-"`         // 用户 ID 类型
	Requests           []*BatchUpdateDocxBlockRequest `json:"requests,omitempty"`             // 批量更新请求列表, 最大 200 条
}

// BatchUpdateDocxBlockRequest 单个块更新请求
type BatchUpdateDocxBlockRequest struct {
	BlockID            string                               `json:"block_id"`                        // 块的唯一标识
	UpdateTextElements *UpdateDocxBlockReqUpdateTextElements `json:"update_text_elements,omitempty"`  // 更新文本元素请求
	UpdateText         *UpdateDocxBlockReqUpdateText         `json:"update_text,omitempty"`           // 更新文本元素及样式请求
	ReplaceImage       *UpdateDocxBlockReqReplaceImage       `json:"replace_image,omitempty"`         // 替换图片请求
	ReplaceFile        *UpdateDocxBlockReqReplaceFile        `json:"replace_file,omitempty"`          // 替换附件请求
}

// BatchUpdateDocxBlockResp ...
type BatchUpdateDocxBlockResp struct {
	DocumentRevisionID int64  `json:"document_revision_id,omitempty"` // 更新后文档版本号
	ClientToken        string `json:"client_token,omitempty"`         // 幂等标识
}

// batchUpdateDocxBlockResp ...
type batchUpdateDocxBlockResp struct {
	Code  int64                     `json:"code,omitempty"`
	Msg   string                    `json:"msg,omitempty"`
	Data  *BatchUpdateDocxBlockResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
