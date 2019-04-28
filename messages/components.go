package messages

type Create struct {
  Interactor CompositeInteractorSpec
  
  Version VersionSpec
  Build   BuildSpec
  Package PackageSpec
  Deploy  DeploymentSpec
}

type Update struct {
  Interactor CompositeInteractorSpec
  
  Version VersionSpec
  Build   BuildSpec
  Package PackageSpec
  Deploy  DeploymentSpec
}

type Deprecate struct {
}

type Archive struct {
}



// InteractorSpec describes an Interactor made up of one or more other Interactors.
type CompositeInteractorSpec struct {
  Location    string
  TypeName    string
  
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

