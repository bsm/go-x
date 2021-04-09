package hashx

import (
	"fmt"
	"testing"

	. "github.com/bsm/ginkgo/extensions/table"
	. "github.com/bsm/gomega"
)

var _ = DescribeTable("XX32",
	func(b []byte, x uint32) {
		Expect(XX32(b, 0)).To(Equal(x))
	},
	Entry("nil", ([]byte)(nil), uint32(0)),
	Entry("blank", []byte{}, uint32(0)),
	Entry("hello", []byte("hello"), uint32(4211111929)),
	Entry("longer value", []byte("//tabs.ultimate-guitar.com/t/three_days_grace/painkiller_tab.htm"), uint32(784213349)),
	Entry("utf8", []byte("日本国"), uint32(187444576)),
)

var _ = DescribeTable("XX64",
	func(b []byte, x uint64) {
		Expect(XX64(b, 0)).To(Equal(x))
	},
	Entry("nil", ([]byte)(nil), uint64(0)),
	Entry("blank", []byte{}, uint64(0)),
	Entry("hello", []byte("hello"), uint64(2794345569481354659)),
	Entry("utf8", []byte("日本国"), uint64(8331024367628775902)),
)

// --------------------------------------------------------------------

func BenchmarkXX32(b *testing.B) {
	p := []byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	for i := 0; i < b.N; i++ {
		XX32(p, 0)
	}
}

func BenchmarkXX64(b *testing.B) {
	p := []byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	for i := 0; i < b.N; i++ {
		XX64(p, 0)
	}
}

// --------------------------------------------------------------------

func ExampleXX32() {
	p := []byte("One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin.")
	fmt.Println(XX32(p, 0))
	// Output: 119031159
}

func ExampleXX64() {
	p := []byte("One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin.")
	fmt.Println(XX64(p, 0))
	// Output: 2399806787706828701
}
