package hashx

import (
	"fmt"
	"testing"

	. "github.com/bsm/ginkgo/extensions/table"
	. "github.com/bsm/gomega"
)

var _ = DescribeTable("Fnv32a",
	func(b []byte, x uint32) {
		Expect(Fnv32a(b)).To(Equal(x))
	},
	Entry("nil", ([]byte)(nil), uint32(0)),
	Entry("blank", []byte{}, uint32(0)),
	Entry("hello", []byte("hello"), uint32(1335831723)),
	Entry("utf8", []byte("日本国"), uint32(2098335336)),
)

var _ = DescribeTable("Fnv64a",
	func(b []byte, x uint64) {
		Expect(Fnv64a(b)).To(Equal(x))
	},
	Entry("nil", ([]byte)(nil), uint64(0)),
	Entry("blank", []byte{}, uint64(0)),
	Entry("hello", []byte("hello"), uint64(11831194018420276491)),
	Entry("utf8", []byte("日本国"), uint64(14277310999654806792)),
)

var _ = DescribeTable("Fnv32aString",
	func(s string, x uint32) {
		Expect(Fnv32aString(s)).To(Equal(x))
	},
	Entry("blank", "", uint32(0)),
	Entry("hello", "hello", uint32(1335831723)),
	Entry("utf8", "日本国", uint32(2098335336)),
)

var _ = DescribeTable("Fnv64aString",
	func(s string, x uint64) {
		Expect(Fnv64aString(s)).To(Equal(x))
	},
	Entry("blank", "", uint64(0)),
	Entry("hello", "hello", uint64(11831194018420276491)),
	Entry("utf8", "日本国", uint64(14277310999654806792)),
)

// --------------------------------------------------------------------

func BenchmarkFnv32a(b *testing.B) {
	p := []byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	for i := 0; i < b.N; i++ {
		Fnv32a(p)
	}
}

func BenchmarkFnv64a(b *testing.B) {
	p := []byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	for i := 0; i < b.N; i++ {
		Fnv64a(p)
	}
}

func BenchmarkFnv32aString(b *testing.B) {
	s := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	for i := 0; i < b.N; i++ {
		Fnv32aString(s)
	}
}

func BenchmarkFnv64aString(b *testing.B) {
	s := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	for i := 0; i < b.N; i++ {
		Fnv64aString(s)
	}
}

// --------------------------------------------------------------------

func ExampleFnv32a() {
	p := []byte("One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin.")
	fmt.Println(Fnv32a(p))
	// Output: 1538131534
}

func ExampleFnv64a() {
	p := []byte("One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin.")
	fmt.Println(Fnv64a(p))
	// Output: 15849972611371769326
}

func ExampleFnv32aString() {
	s := "One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin."
	fmt.Println(Fnv32aString(s))
	// Output: 1538131534
}

func ExampleFnv64aString() {
	s := "One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin."
	fmt.Println(Fnv64aString(s))
	// Output: 15849972611371769326
}
