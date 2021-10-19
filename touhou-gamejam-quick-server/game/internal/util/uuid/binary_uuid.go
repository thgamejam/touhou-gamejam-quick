package uuid

import (
    "database/sql/driver"
    "github.com/google/uuid"
)

//BinaryUUID -> new datatype
type BinaryUUID uuid.UUID

func New() BinaryUUID {
    return BinaryUUID(uuid.New())
}

// StringToBinaryUUID -> parse string to BinaryUUID
func StringToBinaryUUID(s string) (BinaryUUID, error) {
    id, err := uuid.Parse(s)
    return BinaryUUID(id), err
}

//String -> String Representation of Binary16
func (b BinaryUUID) String() string {
    return uuid.UUID(b).String()
}

//GormDataType -> sets type to binary(16)
func (b BinaryUUID) GormDataType() string {
    return "binary(16)"
}

func (b BinaryUUID) MarshalJSON() ([]byte, error) {
    s := uuid.UUID(b)
    str := "\"" + s.String() + "\""
    return []byte(str), nil
}

func (b *BinaryUUID) UnmarshalJSON(by []byte) error {
    s, err := uuid.ParseBytes(by)
    *b = BinaryUUID(s)
    return err
}

// Scan --> tells GORM how to receive from the database
func (b *BinaryUUID) Scan(value interface{}) error {

    bytes, _ := value.([]byte)
    parseByte, err := uuid.FromBytes(bytes)
    *b = BinaryUUID(parseByte)
    return err
}

// Value -> tells GORM how to save into the database
func (b BinaryUUID) Value() (driver.Value, error) {
    return uuid.UUID(b).MarshalBinary()
}