package parcels


import (
	"database/sql/driver"
)


// Value is to make it so parcels.internalBytesParcel fits the database/sql/driver.Valuer interface.
func (parcel *internalBytesParcel) Value() (driver.Value, error) {
	s := parcel.String()
	return s, nil
}


// Value is to make it so parcels.internalStringParcel fits the database/sql/driver.Valuer interface.
func (parcel *internalStringParcel) Value() (driver.Value, error) {
	s := parcel.String()
	return s, nil
}
