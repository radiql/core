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
