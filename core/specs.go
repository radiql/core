package core

// InteractorTypeRef defines the location, name and version of an Interactor type.
type InteractorTypeRef struct {
	TypeLocation string
	TypeName     string
	TypeVersion  VersionSpec
}

// VersionSpec defines the a version using Semantic Versioning (SemVer) with an
// optional label.
type VersionSpec struct {
	MajorVersion string
	MinorVersion string
	BuildNumber  string
	Label        string
}

// InteractorSpec defines the specification of an Interactor to be included in a
// network for Interactors.
type InteractorSpec struct {
	TypeRef InteractorTypeRef

	// Port mappings allow ports to be renamed. This is optional
	PortMappings map[string]string
}

// CompositeInteractorSpec describes an Interactor made up of one or more other Interactors.
type CompositeInteractorSpec struct {
	TypeRef InteractorTypeRef

	Interactors InteractorSpecs
	InPorts     InPortSpecs
	OutPorts    OutPortSpecs
	Connections ConnectionSpecs
}

// InPortSpec defines an input port. The data type accepted by the port is given by the implementation.
type InPortSpec struct {
	// The interactor unique name
	TargetInteractor string
	// The target port unique name
	TargetPort string
}

// OutPortSpec defines an output port. The data type accepted by the port is given by the implementation.
type OutPortSpec struct {
	// The interactor unique name
	SourceInteractor string
	// The source port unique name
	SourcePort string
}

// ConnectionSpec defines the connection between an InPort and an OutPort.
// This connection can be buffered or unbuffered.
type ConnectionSpec struct {
	Name string

	Source OutPortSpec
	Target InPortSpec

	BufferSize uint16
}

// InteractorSpecs defines a named map of Interactor specifications.
type InteractorSpecs map[string]InteractorSpec

// InPortSpecs defines a map of input port names to input port specifications
type InPortSpecs map[string]InPortSpec

// OutPortSpecs defines a map of output port names to output port specifications
type OutPortSpecs map[string]OutPortSpec

// ConnectionSpecs defines a list of connections specifications
type ConnectionSpecs []ConnectionSpec

// DummyStruct is stupid - please remove
type DummyStruct struct {
	SillyName string
}
