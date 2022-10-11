package maintenance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// Application ...
	Application CreatedByType = "Application"
	// Key ...
	Key CreatedByType = "Key"
	// ManagedIdentity ...
	ManagedIdentity CreatedByType = "ManagedIdentity"
	// User ...
	User CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{Application, Key, ManagedIdentity, User}
}

// ImpactType enumerates the values for impact type.
type ImpactType string

const (
	// Freeze Pending updates can freeze network or disk io operation on resource.
	Freeze ImpactType = "Freeze"
	// None Pending updates has no impact on resource.
	None ImpactType = "None"
	// Redeploy Pending updates can redeploy resource.
	Redeploy ImpactType = "Redeploy"
	// Restart Pending updates can cause resource to restart.
	Restart ImpactType = "Restart"
)

// PossibleImpactTypeValues returns an array of possible values for the ImpactType const type.
func PossibleImpactTypeValues() []ImpactType {
	return []ImpactType{Freeze, None, Redeploy, Restart}
}

// RebootOptions enumerates the values for reboot options.
type RebootOptions string

const (
	// Always ...
	Always RebootOptions = "Always"
	// IfRequired ...
	IfRequired RebootOptions = "IfRequired"
	// Never ...
	Never RebootOptions = "Never"
)

// PossibleRebootOptionsValues returns an array of possible values for the RebootOptions const type.
func PossibleRebootOptionsValues() []RebootOptions {
	return []RebootOptions{Always, IfRequired, Never}
}

// Scope enumerates the values for scope.
type Scope string

const (
	// Extension This maintenance scope controls extension installation on VM/VMSS
	Extension Scope = "Extension"
	// Host This maintenance scope controls installation of azure platform updates i.e. services on physical
	// nodes hosting customer VMs.
	Host Scope = "Host"
	// InGuestPatch This maintenance scope controls installation of windows and linux packages on VM/VMSS
	InGuestPatch Scope = "InGuestPatch"
	// OSImage This maintenance scope controls os image installation on VM/VMSS
	OSImage Scope = "OSImage"
	// SQLDB This maintenance scope controls installation of SQL server platform updates.
	SQLDB Scope = "SQLDB"
	// SQLManagedInstance This maintenance scope controls installation of SQL managed instance platform update.
	SQLManagedInstance Scope = "SQLManagedInstance"
)

// PossibleScopeValues returns an array of possible values for the Scope const type.
func PossibleScopeValues() []Scope {
	return []Scope{Extension, Host, InGuestPatch, OSImage, SQLDB, SQLManagedInstance}
}

// TaskScope enumerates the values for task scope.
type TaskScope string

const (
	// TaskScopeGlobal ...
	TaskScopeGlobal TaskScope = "Global"
	// TaskScopeResource ...
	TaskScopeResource TaskScope = "Resource"
)

// PossibleTaskScopeValues returns an array of possible values for the TaskScope const type.
func PossibleTaskScopeValues() []TaskScope {
	return []TaskScope{TaskScopeGlobal, TaskScopeResource}
}

// UpdateStatus enumerates the values for update status.
type UpdateStatus string

const (
	// Completed All updates are successfully applied.
	Completed UpdateStatus = "Completed"
	// InProgress Updates installation are in progress.
	InProgress UpdateStatus = "InProgress"
	// Pending There are pending updates to be installed.
	Pending UpdateStatus = "Pending"
	// RetryLater Updates installation failed and should be retried later.
	RetryLater UpdateStatus = "RetryLater"
	// RetryNow Updates installation failed but are ready to retry again.
	RetryNow UpdateStatus = "RetryNow"
)

// PossibleUpdateStatusValues returns an array of possible values for the UpdateStatus const type.
func PossibleUpdateStatusValues() []UpdateStatus {
	return []UpdateStatus{Completed, InProgress, Pending, RetryLater, RetryNow}
}

// Visibility enumerates the values for visibility.
type Visibility string

const (
	// Custom Only visible to users with permissions.
	Custom Visibility = "Custom"
	// Public Visible to all users.
	Public Visibility = "Public"
)

// PossibleVisibilityValues returns an array of possible values for the Visibility const type.
func PossibleVisibilityValues() []Visibility {
	return []Visibility{Custom, Public}
}
