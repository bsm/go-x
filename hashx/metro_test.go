package hashx

import (
	"fmt"
	"testing"

	. "github.com/bsm/ginkgo/extensions/table"
	. "github.com/bsm/gomega"
)

var _ = DescribeTable("Metro64",
	func(b []byte, x uint64) {
		Expect(Metro64(b, 0)).To(Equal(x))
	},
	Entry("nil", ([]byte)(nil), uint64(8097384203561113213)),
	Entry("blank", []byte{}, uint64(8097384203561113213)),
	Entry("hello", []byte("hello"), uint64(4571129541730210044)),
	Entry("utf8", []byte("日本国"), uint64(4962688877382791512)),
)

var _ = DescribeTable("Metro64String",
	func(s string, x uint64) {
		Expect(Metro64String(s, 0)).To(Equal(x))
	},
	Entry("blank", "", uint64(8097384203561113213)),
	Entry("hello", "hello", uint64(4571129541730210044)),
	Entry("utf8", "日本国", uint64(4962688877382791512)),
)

// --------------------------------------------------------------------

func BenchmarkMetro64(b *testing.B) {
	p := []byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	for i := 0; i < b.N; i++ {
		Metro64(p, 0)
	}
}

func BenchmarkMetro64String(b *testing.B) {
	p := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	for i := 0; i < b.N; i++ {
		Metro64String(p, 0)
	}
}

// --------------------------------------------------------------------

func ExampleMetro64() {
	p := []byte("One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin.")
	fmt.Println(Metro64(p, 0))
	// Output: 17957037419850818313
}

func ExampleMetro64String() {
	s := "One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin."
	fmt.Println(Metro64String(s, 0))
	// Output: 17957037419850818313
}
