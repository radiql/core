
package core

type CreateMsg struct {
  Interactor CompositeInteractorSpec
  
  Version VersionSpec
  Build   BuildSpec
  Package PackageSpec
  Deploy  DeploymentSpec
  Meta    Metadata
}

type UpdateMsg struct {
  Interactor CompositeInteractorSpec
  
  Version VersionSpec
  Build   BuildSpec
  Package PackageSpec
  Deploy  DeploymentSpec
  Meta    Metadata
}

type DeprecateMsg struct {
  TypeRef InteractorTypeRef
  Meta    Metadata
}

type ArchiveMsg struct {
  TypeRef InteractorTypeRef
  Meta    Metadata
}

type BuildMsg struct {
  TypeRef InteractorTypeRef
  Meta    Metadata
}

type DeployMsg struct {
  TypeRef InteractorTypeRef
  Meta    Metadata
}

type RedeployMsg struct {
  TypeRef InteractorTypeRef
  Meta    Metadata
}

type UndeployMsg struct {
  TypeRef InteractorTypeRef
  Meta    Metadata
}  


type Metadata map[string]string


type InteractorTypeRef struct {
  Location  string
  Name      string
  Version   VersionSpec
}

type VersionSpec struct {
  MajorVersion  string
  MinorVersion  string
  Build         string
  Label         string
}


// InteractorSpec describes an Interactor made up of one or more other Interactors.
type CompositeInteractorSpec struct {
  TypeRef      InteractorTypeRef 
  
  Interactors InteractorSpecs
  InPorts     InPortSpecs
  OutPorts    OutPortSpecs
  Connections ConnectionSpecs
}

// InPortSpec defines an input port. The data type accepted by the port is given by the implementation.
type InPortSpec struct {
  TargetInteractor  string
  TargetPort        string
}

// OutPortSpec defines an output port. The data type accepted by the port is given by the implementation.
type OutPortSpec struct {  
  SourceInteractor  string
  SourcePort        string
}

type ConnectionSpec struct {
  Name        string
  
  Source      OutPortSpec
  Target      InPortSpec
  
  BufferSize  uint16
}


InteractorSpecs map[string]InteractorSpec

InPortSpecs map[string]InPortSpec

OutPortSpecs map[string]OutPortSpec

ConnectionSpecs []ConnectionSpec

