package util

import (
	"database/sql"
	"encoding/json"
	"time"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx/types"
)

// std
func ConvertTimeToInt64(x time.Time) int64 {
	return x.Unix()
}
func ConvertInt64ToTime(x int64) time.Time {
	return time.Unix(x, 0)
}
func ConvertNullInt64ToInt64(x sql.NullInt64) int64 {
	if x.Valid {
		return x.Int64
	}
	return 0
}
func ConvertNullStringToString(x sql.NullString) string {
	if x.Valid {
		return x.String
	}
	return ""
}
func ConvertUint8ToUInt64(x uint8) uint64 {
	return uint64(x)
}
func ConvertUint8ToUInt32(x uint8) uint32 {
	return uint32(x)
}

// protobuf
func ConvertInterfaceToStruct(x interface{}) (*structpb.Struct, error) {
	ret := new(structpb.Struct)
	buf, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	if err := ret.UnmarshalJSON(buf); err != nil {
		return nil, err
	}
	return ret, nil
}

// uuid
func ConvertUUIDToString(x uuid.UUID) string {
	return x.String()
}

func ConvertNullUUIDToString(x uuid.NullUUID) string {
	if x.Valid {
		return ConvertUUIDToString(x.UUID)
	}
	return ""
}

// sqlx
func ConvertStructToJSONText(x *structpb.Struct) (types.JSONText, error) {
	buf, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
