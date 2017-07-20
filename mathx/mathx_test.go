package mathx

import (
	"math"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("Round",
	func(f float64, x float64) {
		if math.IsNaN(x) {
			Expect(math.IsNaN(Round(f))).To(BeTrue())
		} else {
			Expect(Round(f)).To(Equal(x))
		}
	},
	Entry("0.0", 0.0, 0.0),
	Entry("-0.0", -0.0, -0.0),
	Entry("0.5", 0.5, 1.0),
	Entry("0.5+epsilon", 0.5000000000000001, 1.0),
	Entry("0.5-epsilon", 0.49999999999999994, 0.0),
	Entry("-0.5", -0.5, -1.0),
	Entry("-0.5+epsilon", -0.49999999999999994, -0.0),
	Entry("-0.5-epsilon", -0.5000000000000001, -1.0),
	Entry("denormal", 1.390671161567e-309, 0.0),
	Entry("1 bit fraction", 2.2517998136852485e+15, 2.251799813685249e+15),
	Entry("large integer", 4.503599627370497e+15, 4.503599627370497e+15),
	Entry("-inf", math.Inf(-1), math.Inf(-1)),
	Entry("+inf", math.Inf(1), math.Inf(1)),
	Entry("NaN", math.NaN(), math.NaN()),
)

var _ = DescribeTable("RoundN",
	func(f float64, n int, x float64) {
		if math.IsNaN(x) {
			Expect(math.IsNaN(RoundN(f, n))).To(BeTrue())
		} else {
			Expect(RoundN(f, n)).To(Equal(x))
		}
	},
	Entry("0.0 (0)", 0.0, 0, 0.0),
	Entry("-0.0 (0)", -0.0, 0, -0.0),

	Entry("0.0 (2)", 0.0, 2, 0.0),
	Entry("-0.0 (2)", -0.0, 2, -0.0),

	Entry("1.2345 (2)", 1.2345, 2, 1.23),
	Entry("-1.2345 (2)", -1.2345, 2, -1.23),
	Entry("4.5678 (2)", 4.5678, 2, 4.57),
	Entry("-4.5678 (2)", -4.5678, 2, -4.57),

	Entry("-inf", math.Inf(-1), 2, math.Inf(-1)),
	Entry("+inf", math.Inf(1), 2, math.Inf(1)),
	Entry("NaN", math.NaN(), 2, math.NaN()),
)

// ----------------------------------------------------------------------------

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "go-x/mathx")
}
