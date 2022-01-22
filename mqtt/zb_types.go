package mqtt

import (
	"strconv"
)

type ZBInt int64

// UnmarshalJSON is a custom unmarshaller for the base16 ints encoded as strings
// in the JSON payload sent over MQTT
func (z *ZBInt) UnmarshalJSON(data []byte) error {

	s := string(data)
	// Remove quotes
	s = s[1 : len(s)-1]

	i, err := strconv.ParseInt(s, 16, 32)
	if err != nil {
		return err
	}
	*z = ZBInt(i)
	return nil
}
