package pkg

import (
	"bytes"
	"encoding/gob"
)

func EmojiIntToString(i int) string {
	switch i {
	case 1:
		return "1️⃣"
	case 2:
		return "2️⃣"
	case 3:
		return "3️⃣"
	case 4:
		return "4️⃣"
	case 5:
		return "5️⃣"
	case 6:
		return "6️⃣"
	case 7:
		return "7️⃣"
	case 8:
		return "8️⃣"
	case 9:
		return "9️⃣"
	default:
		return "0️⃣"
	}
}

func GetRealSizeOf(v *[]byte) (int, error) {
	b := new(bytes.Buffer)
	if err := gob.NewEncoder(b).Encode(*v); err != nil {
		return 0, err
	}
	return b.Len(), nil
}
