package dlcbf

import (
	"testing"
)

func TestDlcbf(t *testing.T) {
	dlcbf, _ := NewDlcbfForCapacity(1000000)


	count := dlcbf.GetCount()
	if float64(count)*100/235886 < 1 {
		t.Error("Expected error < 1 percent, got", float64(count)*100/235886)
	}

	count = dlcbf.GetCount()
	if count != 0 {
		t.Error("Expected count == 0, got", count)
	}

}


func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}