package hashx

import (
	"fmt"
	"testing"

	. "github.com/bsm/ginkgo/extensions/table"
	. "github.com/bsm/gomega"
)

var _ = DescribeTable("CRC16",
	func(b []byte, x uint16) {
		Expect(CRC16(b)).To(Equal(x))
	},
	Entry("nil", ([]byte)(nil), uint16(0)),
	Entry("blank", []byte{}, uint16(0)),
	Entry("hello", []byte("hello"), uint16(50018)),
	Entry("utf8", []byte("日本国"), uint16(20736)),
)

var _ = DescribeTable("CRC16String",
	func(s string, x uint16) {
		Expect(CRC16String(s)).To(Equal(x))
	},
	Entry("blank", "", uint16(0)),
	Entry("hello", "hello", uint16(50018)),
	Entry("utf8", "日本国", uint16(20736)),
)

// --------------------------------------------------------------------

func BenchmarkCRC16(b *testing.B) {
	p := []byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	for i := 0; i < b.N; i++ {
		CRC16(p)
	}
}

func BenchmarkCRC16String(b *testing.B) {
	s := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	for i := 0; i < b.N; i++ {
		CRC16String(s)
	}
}

// --------------------------------------------------------------------

func ExampleCRC16() {
	p := []byte("One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin.")
	fmt.Println(CRC16(p))
	// Output: 13678
}

func ExampleCRC16String() {
	s := "One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin."
	fmt.Println(CRC16String(s))
	// Output: 13678
}
