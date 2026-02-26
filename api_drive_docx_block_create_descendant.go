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

// CreateDocxDescendant 使用 descendant API 创建嵌套块结构（Table, Callout, QuoteContainer, 嵌套列表等）
// 每次最多 1000 个后代块。
//
// doc: https://open.feishu.cn/document/server-docs/docs/docs/docx-v1/document-block/create-descendant
func (r *DriveService) CreateDocxDescendant(ctx context.Context, request *CreateDocxDescendantReq, options ...MethodOptionFunc) (*CreateDocxDescendantResp, *Response, error) {
	if r.cli.mock.mockDriveCreateDocxDescendant != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#CreateDocxDescendant mock enable")
		return r.cli.mock.mockDriveCreateDocxDescendant(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CreateDocxDescendant",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/docx/v1/documents/:document_id/blocks/:block_id/descendant",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createDocxDescendantResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCreateDocxDescendant mock DriveCreateDocxDescendant method
func (r *Mock) MockDriveCreateDocxDescendant(f func(ctx context.Context, request *CreateDocxDescendantReq, options ...MethodOptionFunc) (*CreateDocxDescendantResp, *Response, error)) {
	r.mockDriveCreateDocxDescendant = f
}

// UnMockDriveCreateDocxDescendant un-mock DriveCreateDocxDescendant method
func (r *Mock) UnMockDriveCreateDocxDescendant() {
	r.mockDriveCreateDocxDescendant = nil
}

// CreateDocxDescendantReq ...
type CreateDocxDescendantReq struct {
	DocumentID         string       `path:"document_id" json:"-"`           // 文档唯一标识
	BlockID            string       `path:"block_id" json:"-"`              // 父块唯一标识
	DocumentRevisionID *int64       `query:"document_revision_id" json:"-"` // 文档版本, -1 表示最新版本
	ClientToken        *string      `query:"client_token" json:"-"`         // 幂等标识
	ChildrenID         []string     `json:"children_id"`                    // 顶层临时 block_id 列表
	Descendants        []*DocxBlock `json:"descendants"`                    // 所有后代块（扁平数组）
	Index              int          `json:"index"`                          // 插入位置, -1 表示末尾
}

// CreateDocxDescendantResp ...
type CreateDocxDescendantResp struct {
	Children           []*CreateDocxDescendantRespChild `json:"children,omitempty"`              // 创建的子块信息
	BlockIDRelations   []*BlockIDRelation               `json:"block_id_relations,omitempty"`    // 临时 ID → 实际 ID 映射
	DocumentRevisionID int64                            `json:"document_revision_id,omitempty"`  // 更新后文档版本号
}

// CreateDocxDescendantRespChild 是 descendant API 返回的单个子块信息
type CreateDocxDescendantRespChild struct {
	BlockID   string         `json:"block_id,omitempty"`
	BlockType DocxBlockType  `json:"block_type,omitempty"`
	Board     *DocxBlockBoard `json:"board,omitempty"`
}

// BlockIDRelation 临时 ID → 实际 ID 的映射
type BlockIDRelation struct {
	TemporaryBlockID string `json:"temporary_block_id,omitempty"`
	BlockID          string `json:"block_id,omitempty"`
}

// createDocxDescendantResp ...
type createDocxDescendantResp struct {
	Code  int64                     `json:"code,omitempty"`
	Msg   string                    `json:"msg,omitempty"`
	Data  *CreateDocxDescendantResp `json:"data,omitempty"`
	Error *ErrorDetail              `json:"error,omitempty"`
}
