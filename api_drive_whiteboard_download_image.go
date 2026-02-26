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
	"io"
)

// DownloadWhiteboardImage 下载画板为图片
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/board-v1/whiteboard/download_as_image
func (r *DriveService) DownloadWhiteboardImage(ctx context.Context, request *DownloadWhiteboardImageReq, options ...MethodOptionFunc) (*DownloadWhiteboardImageResp, *Response, error) {
	if r.cli.mock.mockDriveDownloadWhiteboardImage != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Drive#DownloadWhiteboardImage mock enable")
		return r.cli.mock.mockDriveDownloadWhiteboardImage(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "DownloadWhiteboardImage",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/board/v1/whiteboards/:whiteboard_id/download_as_image",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(downloadWhiteboardImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveDownloadWhiteboardImage mock DriveDownloadWhiteboardImage method
func (r *Mock) MockDriveDownloadWhiteboardImage(f func(ctx context.Context, request *DownloadWhiteboardImageReq, options ...MethodOptionFunc) (*DownloadWhiteboardImageResp, *Response, error)) {
	r.mockDriveDownloadWhiteboardImage = f
}

// UnMockDriveDownloadWhiteboardImage un-mock DriveDownloadWhiteboardImage method
func (r *Mock) UnMockDriveDownloadWhiteboardImage() {
	r.mockDriveDownloadWhiteboardImage = nil
}

// DownloadWhiteboardImageReq ...
type DownloadWhiteboardImageReq struct {
	WhiteboardID string `path:"whiteboard_id" json:"-"` // 画板唯一标识
}

// downloadWhiteboardImageResp ...
type downloadWhiteboardImageResp struct {
	Code int64                        `json:"code,omitempty"`
	Msg  string                       `json:"msg,omitempty"`
	Data *DownloadWhiteboardImageResp `json:"data,omitempty"`
}

func (r *downloadWhiteboardImageResp) SetReader(file io.Reader) {
	if r.Data == nil {
		r.Data = &DownloadWhiteboardImageResp{}
	}
	r.Data.File = file
}

func (r *downloadWhiteboardImageResp) SetFilename(filename string) {
	if r.Data == nil {
		r.Data = &DownloadWhiteboardImageResp{}
	}
	r.Data.Filename = filename
}

// DownloadWhiteboardImageResp ...
type DownloadWhiteboardImageResp struct {
	File     io.Reader `json:"file,omitempty"`
	Filename string    `json:"filename,omitempty"`
}
