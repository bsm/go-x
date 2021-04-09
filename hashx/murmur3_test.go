package hashx

import (
	"fmt"
	"testing"

	. "github.com/bsm/ginkgo/extensions/table"
	. "github.com/bsm/gomega"
)

var _ = DescribeTable("MM32",
	func(b []byte, x uint32) {
		Expect(MM32(b)).To(Equal(x))
	},
	Entry("nil", ([]byte)(nil), uint32(0)),
	Entry("blank", []byte{}, uint32(0)),
	Entry("hello", []byte("hello"), uint32(613153351)),
	Entry("longer value", []byte("//tabs.ultimate-guitar.com/t/three_days_grace/painkiller_tab.htm"), uint32(2829525266)),
	Entry("utf8", []byte("日本国"), uint32(2473588326)),
)

var _ = DescribeTable("MM32String",
	func(s string, x uint32) {
		Expect(MM32String(s)).To(Equal(x))
	},
	Entry("blank", "", uint32(0)),
	Entry("hello", "hello", uint32(613153351)),
	Entry("longer value", "//tabs.ultimate-guitar.com/t/three_days_grace/painkiller_tab.htm", uint32(2829525266)),
	Entry("utf8", "日本国", uint32(2473588326)),
)

var _ = DescribeTable("MM64",
	func(b []byte, x uint64) {
		Expect(MM64(b)).To(Equal(x))
	},
	Entry("nil", ([]byte)(nil), uint64(0)),
	Entry("blank", []byte{}, uint64(0)),
	Entry("hello", []byte("hello"), uint64(14688674573012802306)),
	Entry("utf8", []byte("日本国"), uint64(15494675815960301006)),
)

var _ = DescribeTable("MM64String",
	func(s string, x uint64) {
		Expect(MM64String(s)).To(Equal(x))
	},
	Entry("blank", "", uint64(0)),
	Entry("hello", "hello", uint64(14688674573012802306)),
	Entry("utf8", "日本国", uint64(15494675815960301006)),
)

var _ = DescribeTable("MM128",
	func(b []byte, x1, x2 uint64) {
		n1, n2 := MM128(b)
		Expect(n1).To(Equal(x1))
		Expect(n2).To(Equal(x2))
	},
	Entry("nil", ([]byte)(nil), uint64(0), uint64(0)),
	Entry("blank", []byte{}, uint64(0), uint64(0)),
	Entry("hello", []byte("hello"), uint64(14688674573012802306), uint64(6565844092913065241)),
	Entry("longer value", []byte("//tabs.ultimate-guitar.com/t/three_days_grace/painkiller_tab.htm"), uint64(7923461173300967811), uint64(2180639483051343085)),
	Entry("utf8", []byte("日本国"), uint64(15494675815960301006), uint64(6456422527984082000)),
)

var _ = DescribeTable("MM128String",
	func(s string, x1, x2 uint64) {
		n1, n2 := MM128String(s)
		Expect(n1).To(Equal(x1))
		Expect(n2).To(Equal(x2))
	},
	Entry("blank", "", uint64(0), uint64(0)),
	Entry("hello", "hello", uint64(14688674573012802306), uint64(6565844092913065241)),
	Entry("longer value", "//tabs.ultimate-guitar.com/t/three_days_grace/painkiller_tab.htm", uint64(7923461173300967811), uint64(2180639483051343085)),
	Entry("utf8", "日本国", uint64(15494675815960301006), uint64(6456422527984082000)),
)

// --------------------------------------------------------------------

func BenchmarkMM32(b *testing.B) {
	p := []byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	for i := 0; i < b.N; i++ {
		MM32(p)
	}
}

func BenchmarkMM64(b *testing.B) {
	p := []byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	for i := 0; i < b.N; i++ {
		MM64(p)
	}
}

func BenchmarkMM128(b *testing.B) {
	p := []byte("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	for i := 0; i < b.N; i++ {
		MM128(p)
	}
}

func BenchmarkMM32String(b *testing.B) {
	s := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	for i := 0; i < b.N; i++ {
		MM32String(s)
	}
}

func BenchmarkMM64String(b *testing.B) {
	s := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	for i := 0; i < b.N; i++ {
		MM64String(s)
	}
}

func BenchmarkMM128String(b *testing.B) {
	s := "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	for i := 0; i < b.N; i++ {
		MM128String(s)
	}
}

// --------------------------------------------------------------------

func ExampleMM32() {
	p := []byte("One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin.")
	fmt.Println(MM32(p))
	// Output: 409351655
}

func ExampleMM64() {
	p := []byte("One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin.")
	fmt.Println(MM64(p))
	// Output: 7376845339986439510
}

func ExampleMM128() {
	p := []byte("One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin.")
	fmt.Println(MM128(p))
	// Output: 7376845339986439510 5610010571125513187
}

func ExampleMM32String() {
	s := "One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin."
	fmt.Println(MM32String(s))
	// Output: 409351655
}

func ExampleMM64String() {
	s := "One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin."
	fmt.Println(MM64String(s))
	// Output: 7376845339986439510
}

func ExampleMM128String() {
	s := "One morning, when Gregor Samsa woke from troubled dreams, he found himself transformed in his bed into a horrible vermin."
	fmt.Println(MM128String(s))
	// Output: 7376845339986439510 5610010571125513187
}
