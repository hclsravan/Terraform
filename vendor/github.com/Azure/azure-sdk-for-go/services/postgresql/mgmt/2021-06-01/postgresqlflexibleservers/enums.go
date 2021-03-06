package postgresqlflexibleservers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ConfigurationDataType enumerates the values for configuration data type.
type ConfigurationDataType string

const (
	// ConfigurationDataTypeBoolean ...
	ConfigurationDataTypeBoolean ConfigurationDataType = "Boolean"
	// ConfigurationDataTypeEnumeration ...
	ConfigurationDataTypeEnumeration ConfigurationDataType = "Enumeration"
	// ConfigurationDataTypeInteger ...
	ConfigurationDataTypeInteger ConfigurationDataType = "Integer"
	// ConfigurationDataTypeNumeric ...
	ConfigurationDataTypeNumeric ConfigurationDataType = "Numeric"
)

// PossibleConfigurationDataTypeValues returns an array of possible values for the ConfigurationDataType const type.
func PossibleConfigurationDataTypeValues() []ConfigurationDataType {
	return []ConfigurationDataType{ConfigurationDataTypeBoolean, ConfigurationDataTypeEnumeration, ConfigurationDataTypeInteger, ConfigurationDataTypeNumeric}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// CreatedByTypeApplication ...
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey ...
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity ...
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser ...
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{CreatedByTypeApplication, CreatedByTypeKey, CreatedByTypeManagedIdentity, CreatedByTypeUser}
}

// CreateMode enumerates the values for create mode.
type CreateMode string

const (
	// CreateModeCreate ...
	CreateModeCreate CreateMode = "Create"
	// CreateModeDefault ...
	CreateModeDefault CreateMode = "Default"
	// CreateModePointInTimeRestore ...
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	// CreateModeUpdate ...
	CreateModeUpdate CreateMode = "Update"
)

// PossibleCreateModeValues returns an array of possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{CreateModeCreate, CreateModeDefault, CreateModePointInTimeRestore, CreateModeUpdate}
}

// CreateModeForUpdate enumerates the values for create mode for update.
type CreateModeForUpdate string

const (
	// CreateModeForUpdateDefault ...
	CreateModeForUpdateDefault CreateModeForUpdate = "Default"
	// CreateModeForUpdateUpdate ...
	CreateModeForUpdateUpdate CreateModeForUpdate = "Update"
)

// PossibleCreateModeForUpdateValues returns an array of possible values for the CreateModeForUpdate const type.
func PossibleCreateModeForUpdateValues() []CreateModeForUpdate {
	return []CreateModeForUpdate{CreateModeForUpdateDefault, CreateModeForUpdateUpdate}
}

// GeoRedundantBackupEnum enumerates the values for geo redundant backup enum.
type GeoRedundantBackupEnum string

const (
	// GeoRedundantBackupEnumDisabled ...
	GeoRedundantBackupEnumDisabled GeoRedundantBackupEnum = "Disabled"
	// GeoRedundantBackupEnumEnabled ...
	GeoRedundantBackupEnumEnabled GeoRedundantBackupEnum = "Enabled"
)

// PossibleGeoRedundantBackupEnumValues returns an array of possible values for the GeoRedundantBackupEnum const type.
func PossibleGeoRedundantBackupEnumValues() []GeoRedundantBackupEnum {
	return []GeoRedundantBackupEnum{GeoRedundantBackupEnumDisabled, GeoRedundantBackupEnumEnabled}
}

// HighAvailabilityMode enumerates the values for high availability mode.
type HighAvailabilityMode string

const (
	// HighAvailabilityModeDisabled ...
	HighAvailabilityModeDisabled HighAvailabilityMode = "Disabled"
	// HighAvailabilityModeZoneRedundant ...
	HighAvailabilityModeZoneRedundant HighAvailabilityMode = "ZoneRedundant"
)

// PossibleHighAvailabilityModeValues returns an array of possible values for the HighAvailabilityMode const type.
func PossibleHighAvailabilityModeValues() []HighAvailabilityMode {
	return []HighAvailabilityMode{HighAvailabilityModeDisabled, HighAvailabilityModeZoneRedundant}
}

// OperationOrigin enumerates the values for operation origin.
type OperationOrigin string

const (
	// OperationOriginNotSpecified ...
	OperationOriginNotSpecified OperationOrigin = "NotSpecified"
	// OperationOriginSystem ...
	OperationOriginSystem OperationOrigin = "system"
	// OperationOriginUser ...
	OperationOriginUser OperationOrigin = "user"
)

// PossibleOperationOriginValues returns an array of possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{OperationOriginNotSpecified, OperationOriginSystem, OperationOriginUser}
}

// ServerHAState enumerates the values for server ha state.
type ServerHAState string

const (
	// ServerHAStateCreatingStandby ...
	ServerHAStateCreatingStandby ServerHAState = "CreatingStandby"
	// ServerHAStateFailingOver ...
	ServerHAStateFailingOver ServerHAState = "FailingOver"
	// ServerHAStateHealthy ...
	ServerHAStateHealthy ServerHAState = "Healthy"
	// ServerHAStateNotEnabled ...
	ServerHAStateNotEnabled ServerHAState = "NotEnabled"
	// ServerHAStateRemovingStandby ...
	ServerHAStateRemovingStandby ServerHAState = "RemovingStandby"
	// ServerHAStateReplicatingData ...
	ServerHAStateReplicatingData ServerHAState = "ReplicatingData"
)

// PossibleServerHAStateValues returns an array of possible values for the ServerHAState const type.
func PossibleServerHAStateValues() []ServerHAState {
	return []ServerHAState{ServerHAStateCreatingStandby, ServerHAStateFailingOver, ServerHAStateHealthy, ServerHAStateNotEnabled, ServerHAStateRemovingStandby, ServerHAStateReplicatingData}
}

// ServerPublicNetworkAccessState enumerates the values for server public network access state.
type ServerPublicNetworkAccessState string

const (
	// ServerPublicNetworkAccessStateDisabled ...
	ServerPublicNetworkAccessStateDisabled ServerPublicNetworkAccessState = "Disabled"
	// ServerPublicNetworkAccessStateEnabled ...
	ServerPublicNetworkAccessStateEnabled ServerPublicNetworkAccessState = "Enabled"
)

// PossibleServerPublicNetworkAccessStateValues returns an array of possible values for the ServerPublicNetworkAccessState const type.
func PossibleServerPublicNetworkAccessStateValues() []ServerPublicNetworkAccessState {
	return []ServerPublicNetworkAccessState{ServerPublicNetworkAccessStateDisabled, ServerPublicNetworkAccessStateEnabled}
}

// ServerState enumerates the values for server state.
type ServerState string

const (
	// ServerStateDisabled ...
	ServerStateDisabled ServerState = "Disabled"
	// ServerStateDropping ...
	ServerStateDropping ServerState = "Dropping"
	// ServerStateReady ...
	ServerStateReady ServerState = "Ready"
	// ServerStateStarting ...
	ServerStateStarting ServerState = "Starting"
	// ServerStateStopped ...
	ServerStateStopped ServerState = "Stopped"
	// ServerStateStopping ...
	ServerStateStopping ServerState = "Stopping"
	// ServerStateUpdating ...
	ServerStateUpdating ServerState = "Updating"
)

// PossibleServerStateValues returns an array of possible values for the ServerState const type.
func PossibleServerStateValues() []ServerState {
	return []ServerState{ServerStateDisabled, ServerStateDropping, ServerStateReady, ServerStateStarting, ServerStateStopped, ServerStateStopping, ServerStateUpdating}
}

// ServerVersion enumerates the values for server version.
type ServerVersion string

const (
	// ServerVersionOneOne ...
	ServerVersionOneOne ServerVersion = "11"
	// ServerVersionOneThree ...
	ServerVersionOneThree ServerVersion = "13"
	// ServerVersionOneTwo ...
	ServerVersionOneTwo ServerVersion = "12"
)

// PossibleServerVersionValues returns an array of possible values for the ServerVersion const type.
func PossibleServerVersionValues() []ServerVersion {
	return []ServerVersion{ServerVersionOneOne, ServerVersionOneThree, ServerVersionOneTwo}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// SkuTierBurstable ...
	SkuTierBurstable SkuTier = "Burstable"
	// SkuTierGeneralPurpose ...
	SkuTierGeneralPurpose SkuTier = "GeneralPurpose"
	// SkuTierMemoryOptimized ...
	SkuTierMemoryOptimized SkuTier = "MemoryOptimized"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{SkuTierBurstable, SkuTierGeneralPurpose, SkuTierMemoryOptimized}
}
