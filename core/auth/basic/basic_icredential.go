// Copyright 2020 Huawei Technologies Co.,Ltd.
//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package basic

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/iam"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/signer"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
	"strings"
)

const (
	ProjectIdInHeader     = "X-Project-Id"
	SecurityTokenInHeader = "X-Security-Token"
	ContentTypeInHeader   = "Content-Type"
)

type Credentials struct {
	IamEndpoint   string
	AK            string
	SK            string
	ProjectId     string
	SecurityToken string
}

func (s Credentials) ProcessAuthParams(client *impl.DefaultHttpClient, region string) auth.ICredential {
	if s.ProjectId != "" {
		return s
	}

	req, err := s.ProcessAuthRequest(client, iam.GetKeystoneListProjectsRequest(s.IamEndpoint, region))
	if err != nil {
		panic(fmt.Sprintf("failed to get project id, %s", err.Error()))
	}

	id, err := iam.KeystoneListProjects(client, req)
	if err != nil {
		panic(fmt.Sprintf("failed to get project id, %s", err.Error()))
	}

	s.ProjectId = id
	return s
}

func (s Credentials) ProcessAuthRequest(client *impl.DefaultHttpClient, req *request.DefaultHttpRequest) (*request.DefaultHttpRequest, error) {
	reqBuilder := req.Builder()

	if s.ProjectId != "" {
		reqBuilder = reqBuilder.
			AddAutoFilledPathParam("project_id", s.ProjectId).
			AddHeaderParam(ProjectIdInHeader, s.ProjectId)
	}

	if s.SecurityToken != "" {
		reqBuilder.AddHeaderParam(SecurityTokenInHeader, s.SecurityToken)
	}

	if _, ok := req.GetHeaderParams()[ContentTypeInHeader]; ok {
		if !strings.Contains(req.GetHeaderParams()[ContentTypeInHeader], "application/json") {
			reqBuilder.AddHeaderParam("X-Sdk-Content-Sha256", "UNSIGNED-PAYLOAD")
		}
	}

	headerParams, err := signer.Sign(reqBuilder.Build(), s.AK, s.SK)
	if err != nil {
		return nil, err
	}

	for key, value := range headerParams {
		req.AddHeaderParam(key, value)
	}
	return req, nil
}

type CredentialsBuilder struct {
	Credentials Credentials
}

func NewCredentialsBuilder() *CredentialsBuilder {
	return &CredentialsBuilder{Credentials: Credentials{
		IamEndpoint: iam.DefaultIamEndpoint,
	}}
}

func (builder *CredentialsBuilder) WithIamEndpointOverride(endpoint string) *CredentialsBuilder {
	builder.Credentials.IamEndpoint = endpoint
	return builder
}

func (builder *CredentialsBuilder) WithAk(ak string) *CredentialsBuilder {
	builder.Credentials.AK = ak
	return builder
}

func (builder *CredentialsBuilder) WithSk(sk string) *CredentialsBuilder {
	builder.Credentials.SK = sk
	return builder
}

func (builder *CredentialsBuilder) WithProjectId(projectId string) *CredentialsBuilder {
	builder.Credentials.ProjectId = projectId
	return builder
}

func (builder *CredentialsBuilder) WithSecurityToken(token string) *CredentialsBuilder {
	builder.Credentials.SecurityToken = token
	return builder
}

func (builder *CredentialsBuilder) Build() Credentials {
	return builder.Credentials
}
