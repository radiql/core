package contacts


type TelephoneNumber string 

type EmailAddress string

type PostalAddress struct {
  BuildingName  string
  Street        string
  AddressLine1  string
  AddressLine2  string
  PostTown      string
  Postcode      string
  Country       string
}

type GPSCoordinates struct {
  Lat   string
  Long  string
}

