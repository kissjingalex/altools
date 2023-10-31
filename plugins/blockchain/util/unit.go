package util

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
	"regexp"
)

const (
	ExponentWei        = 0  // 1 Wei        = 10e0  = 1 Wei
	ExponentKwei       = 3  // 1 Kwei       = 10e3  = 1,000 Wei
	ExponentMwei       = 6  // 1 Mwei       = 10e6  = 1,000,000 Wei
	ExponentGwei       = 9  // 1 Gwei       = 10e9  = 1,000,000,000 Wei
	ExponentMicroether = 12 // 1 Microether = 10e12 = 1,000,000,000,000 Wei
	ExponentMilliether = 15 // 1 Milliether = 10e15 = 1,000,000,000,000,000 Wei
	ExponentEther      = 18 // 1 Ether      = 10e18 = 1,000,000,000,000,000,000 Wei
)

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

func ToEther(wei *big.Int) *big.Float {
	floatWei := new(big.Float)
	floatWei.SetString(wei.String())
	// Quo means Div
	ethValue := new(big.Float).Quo(floatWei, big.NewFloat(math.Pow10(ExponentEther)))
	return ethValue
}

func ToEtherDecimal(wei *big.Int) decimal.Decimal {
	decimalWei := decimal.NewFromBigInt(wei, 0)
	denominator := Pow10(ExponentEther)
	ethValue := decimalWei.Div(denominator)
	return ethValue
}

func Pow10(exponent int64) decimal.Decimal {
	base := decimal.NewFromInt(10)
	return base.Pow(decimal.NewFromInt(exponent))
}

// ToDecimal wei to decimals
func ToDecimal(wei interface{}, exponent int) decimal.Decimal {
	value := new(big.Int)
	switch v := wei.(type) {
	case string:
		value.SetString(v, 10)
	case *big.Int:
		value = v
	}

	divisor := Pow10(int64(exponent))
	dividend, _ := decimal.NewFromString(value.String())
	quotient := dividend.Div(divisor)
	return quotient
}

// ToWei decimals to wei
func ToWei(value interface{}, exponent int) *big.Int {
	multiplicand := decimal.NewFromFloat(0)
	switch v := value.(type) {
	case string:
		multiplicand, _ = decimal.NewFromString(v)
	case float64:
		multiplicand = decimal.NewFromFloat(v)
	case int64:
		multiplicand = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		multiplicand = v
	case *decimal.Decimal:
		multiplicand = *v
	}

	multiplier := Pow10(int64(exponent))
	product := multiplicand.Mul(multiplier)

	wei := new(big.Int)
	wei.SetString(product.String(), 10)
	return wei
}
