package core

// BuildSpec defines how the software is to be built
type BuildSpec struct {
}

// PackageSpec defines how the software is to be packaged
type PackageSpec struct {
}

// DeploymentSpec defines how software is to be deployed
type DeploymentSpec struct {
}

// Metadata is a map of key value pairs the can be used to hold metada.
type Metadata map[string]string

// CreateMsg is an instruction to create a CompositeInteractor based
// upon the CompositeInteractorSpec. The CompositeInteractorSpec
// specifies the functionality. The Build, Package, Deploy fields
/// specify how that functionality should be implemented.
type CreateMsg struct {
	Interactor CompositeInteractorSpec

	Version VersionSpec
	Build   BuildSpec
	Package PackageSpec
	Deploy  DeploymentSpec
	Meta    Metadata
}

// UpdateMsg is an instruction to re-create a CompositeInteractor based
// upon the CompositeInteractorSpec. The CompositeInteractorSpec
// specifies the functionality. The Build, Package, Deploy fields
// specify how that functionality should be reimplemented.
type UpdateMsg struct {
	Interactor CompositeInteractorSpec

	Version VersionSpec
	Build   BuildSpec
	Package PackageSpec
	Deploy  DeploymentSpec
	Meta    Metadata
}

// DeprecateMsg is an instruction to deprecate the particular version of
// the Interactor specified by the InteractorTypeRef
type DeprecateMsg struct {
	TypeRef InteractorTypeRef
	Meta    Metadata
}

// ArchiveMsg is an instruction to archive the particular version of
// the Interactor specified by the InteractorTypeRef
type ArchiveMsg struct {
	TypeRef InteractorTypeRef
	Meta    Metadata
}

// BuildMsg is an instruction to (re-) build the particular version
// of the Interactor specified by the InteractorTypeRef
type BuildMsg struct {
	TypeRef InteractorTypeRef
	Meta    Metadata
}

// DeployMsg is an instruction to (re-)deploy the particular version
// of the Interactor specified by the InteractorTypeRef
type DeployMsg struct {
	TypeRef InteractorTypeRef
	Meta    Metadata
}

// UndeployMsg is an instruction to undeploy the particular version
// of the Interactor specified by the InteractorTypeRef
type UndeployMsg struct {
	TypeRef InteractorTypeRef
	Meta    Metadata
}

// HotSwapMsg is an instruction to substitute all instances of the
// RemoveInteractor with the InsertInteractor.
type HotSwapMsg struct {
	Container        InteractorTypeRef
	RemoveInteractor InteractorTypeRef
	InsertInteractor InteractorTypeRef

	Meta Metadata
}
