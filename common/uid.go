package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcutil/base58"
)

/**
 * UID is method to generate an virtual unique identifier for whole system
 * it's structure contains 60 bits: LocalID - ObjectType - ShardID
 * 32 bits for LocalID, max (2^32) - 1
 * 10 bits for ObjectType
 * 18 bits for ShardID
 * */

type UID struct {
	LocalID    uint32
	ObjectType int
	ShardID    uint32
}

func NewUID(localId uint32, objectType int, shardId uint32) UID {
	return UID{
		LocalID:    localId,
		ObjectType: objectType,
		ShardID:    shardId,
	}
}

/** LocalID: 1, Object: 1, ShardID: 1 => 0001 0001 0001
 * 1 << 8 = 0001 0000 0000
 * 1 << 4 =         1 0000
 * 1 << 0 =              1
 * => 0001 0001 0001
 */
func (uid UID) String() string {
	val := uint64(uid.LocalID)<<28 | uint64(uid.ObjectType)<<18 | uint64(uid.ShardID)<<0
	return base58.Encode([]byte(fmt.Sprintf("%v", val)))
}

func (uid UID) GetLocalID() uint32 {
	return uid.LocalID
}

func (uid UID) GetShardID() uint32 {
	return uid.ShardID
}

func (uid UID) GetObjectType() int {
	return uid.ObjectType
}

func DecomposeUID(s string) (UID, error) {
	uid, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return UID{}, err
	}

	if (1 << 18) > uid {
		return UID{}, errors.New("wrong uid")
	}

	/**
	 * int(uid >> 18 & 0x3FF):
	 * 10111100 00000 >> 5 => 10111100
	 * 10111100 & 0x3F (111111) = 10111100 & 00111111 = 00111100
	 * */

	u := UID{
		LocalID:    uint32(uid >> 28),
		ObjectType: int(uid >> 18 & 0x3FF),
		ShardID:    uint32(uid >> 0 & 0x3FFFF),
	}

	return u, nil
}

func FromBase58(s string) (UID, error) {
	return DecomposeUID(string(base58.Decode(s)))
}

func (uid UID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", uid.String())), nil
}

func (uid UID) UnMarshalJSON(data []byte) error {
	decodeUID, err := FromBase58(strings.Replace(string(data), "\"", "", -1))

	if err != nil {
		return err
	}
	uid.LocalID = decodeUID.LocalID
	uid.ShardID = decodeUID.ShardID
	uid.ObjectType = decodeUID.ObjectType

	return nil
}

func (uid *UID) Value() (driver.Value, error) {
	if uid == nil {
		return nil, nil
	}
	return int64(uid.LocalID), nil
}

func (uid *UID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var i uint32

	switch t := value.(type) {
	case int:
		i = uint32(t)
	case int8:
		i = uint32(t) // standardizes across systems
	case int16:
		i = uint32(t) // standardizes across systems
	case int32:
		i = uint32(t) // standardizes across systems
	case int64:
		i = uint32(t) // standardizes across systems
	case uint8:
		i = uint32(t) // standardizes across systems
	case uint16:
		i = uint32(t) // standardizes across systems
	case uint32:
		i = t
	case uint64:
		i = uint32(t)
	case []byte:
		a, err := strconv.Atoi(string(t))
		if err != nil {
			return err
		}
		i = uint32(a)
	default:
		return errors.New("invalid Scan Source")
	}

	*uid = NewUID(i, 0, 1)

	return nil
}
