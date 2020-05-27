// +build go1.9

// Copyright 2020 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package locks

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/locks"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type LockLevel = original.LockLevel

const (
	CanNotDelete LockLevel = original.CanNotDelete
	NotSpecified LockLevel = original.NotSpecified
	ReadOnly     LockLevel = original.ReadOnly
)

type AuthorizationOperationsClient = original.AuthorizationOperationsClient
type BaseClient = original.BaseClient
type ManagementLockListResult = original.ManagementLockListResult
type ManagementLockListResultIterator = original.ManagementLockListResultIterator
type ManagementLockListResultPage = original.ManagementLockListResultPage
type ManagementLockObject = original.ManagementLockObject
type ManagementLockOwner = original.ManagementLockOwner
type ManagementLockProperties = original.ManagementLockProperties
type ManagementLocksClient = original.ManagementLocksClient
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAuthorizationOperationsClient(subscriptionID string) AuthorizationOperationsClient {
	return original.NewAuthorizationOperationsClient(subscriptionID)
}
func NewAuthorizationOperationsClientWithBaseURI(baseURI string, subscriptionID string) AuthorizationOperationsClient {
	return original.NewAuthorizationOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagementLockListResultIterator(page ManagementLockListResultPage) ManagementLockListResultIterator {
	return original.NewManagementLockListResultIterator(page)
}
func NewManagementLockListResultPage(getNextPage func(context.Context, ManagementLockListResult) (ManagementLockListResult, error)) ManagementLockListResultPage {
	return original.NewManagementLockListResultPage(getNextPage)
}
func NewManagementLocksClient(subscriptionID string) ManagementLocksClient {
	return original.NewManagementLocksClient(subscriptionID)
}
func NewManagementLocksClientWithBaseURI(baseURI string, subscriptionID string) ManagementLocksClient {
	return original.NewManagementLocksClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleLockLevelValues() []LockLevel {
	return original.PossibleLockLevelValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
