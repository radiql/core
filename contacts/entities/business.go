// This file describes data entity data structures associated with 
// business contacts

package contacts


type BusinessContact struct {
  Name          string
  CompanyNumber string
  
  Address   PostalAddress
}

type Employee struct {
  ID        string
  
  Forenames Name
  Surname   Name
  
  Role      JobDescription
  Email     EmailAddress
  WorkTel   TelephoneNumber
  HomeTel   TelephoneNumber
  MobileTel TelephoneNumber
}


