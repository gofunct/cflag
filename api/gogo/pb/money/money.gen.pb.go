// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/money/money.gen.proto

package pbmoney

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CurrencyCode int32

const (
	CurrencyCode_CURRENCY_CODE_NONE CurrencyCode = 0
	CurrencyCode_CURRENCY_CODE_AED  CurrencyCode = 1
	CurrencyCode_CURRENCY_CODE_AFN  CurrencyCode = 2
	CurrencyCode_CURRENCY_CODE_ALL  CurrencyCode = 3
	CurrencyCode_CURRENCY_CODE_AMD  CurrencyCode = 4
	CurrencyCode_CURRENCY_CODE_ANG  CurrencyCode = 5
	CurrencyCode_CURRENCY_CODE_AOA  CurrencyCode = 6
	CurrencyCode_CURRENCY_CODE_ARS  CurrencyCode = 7
	CurrencyCode_CURRENCY_CODE_AUD  CurrencyCode = 8
	CurrencyCode_CURRENCY_CODE_AWG  CurrencyCode = 9
	CurrencyCode_CURRENCY_CODE_AZN  CurrencyCode = 10
	CurrencyCode_CURRENCY_CODE_BAM  CurrencyCode = 11
	CurrencyCode_CURRENCY_CODE_BBD  CurrencyCode = 12
	CurrencyCode_CURRENCY_CODE_BDT  CurrencyCode = 13
	CurrencyCode_CURRENCY_CODE_BGN  CurrencyCode = 14
	CurrencyCode_CURRENCY_CODE_BHD  CurrencyCode = 15
	CurrencyCode_CURRENCY_CODE_BIF  CurrencyCode = 16
	CurrencyCode_CURRENCY_CODE_BMD  CurrencyCode = 17
	CurrencyCode_CURRENCY_CODE_BND  CurrencyCode = 18
	CurrencyCode_CURRENCY_CODE_BOB  CurrencyCode = 19
	CurrencyCode_CURRENCY_CODE_BRL  CurrencyCode = 20
	CurrencyCode_CURRENCY_CODE_BSD  CurrencyCode = 21
	CurrencyCode_CURRENCY_CODE_BWP  CurrencyCode = 22
	CurrencyCode_CURRENCY_CODE_BYN  CurrencyCode = 23
	CurrencyCode_CURRENCY_CODE_BZD  CurrencyCode = 24
	CurrencyCode_CURRENCY_CODE_CAD  CurrencyCode = 25
	CurrencyCode_CURRENCY_CODE_CDF  CurrencyCode = 26
	CurrencyCode_CURRENCY_CODE_CHF  CurrencyCode = 27
	CurrencyCode_CURRENCY_CODE_CLP  CurrencyCode = 28
	CurrencyCode_CURRENCY_CODE_CNY  CurrencyCode = 29
	CurrencyCode_CURRENCY_CODE_COP  CurrencyCode = 30
	CurrencyCode_CURRENCY_CODE_CRC  CurrencyCode = 31
	CurrencyCode_CURRENCY_CODE_CVE  CurrencyCode = 32
	CurrencyCode_CURRENCY_CODE_CZK  CurrencyCode = 33
	CurrencyCode_CURRENCY_CODE_DJF  CurrencyCode = 34
	CurrencyCode_CURRENCY_CODE_DKK  CurrencyCode = 35
	CurrencyCode_CURRENCY_CODE_DOP  CurrencyCode = 36
	CurrencyCode_CURRENCY_CODE_DZD  CurrencyCode = 37
	CurrencyCode_CURRENCY_CODE_EGP  CurrencyCode = 38
	CurrencyCode_CURRENCY_CODE_ERN  CurrencyCode = 39
	CurrencyCode_CURRENCY_CODE_ETB  CurrencyCode = 40
	CurrencyCode_CURRENCY_CODE_EUR  CurrencyCode = 41
	CurrencyCode_CURRENCY_CODE_FJD  CurrencyCode = 42
	CurrencyCode_CURRENCY_CODE_GBP  CurrencyCode = 43
	CurrencyCode_CURRENCY_CODE_GEL  CurrencyCode = 44
	CurrencyCode_CURRENCY_CODE_GHS  CurrencyCode = 45
	CurrencyCode_CURRENCY_CODE_GIP  CurrencyCode = 46
	CurrencyCode_CURRENCY_CODE_GMD  CurrencyCode = 47
	CurrencyCode_CURRENCY_CODE_GNF  CurrencyCode = 48
	CurrencyCode_CURRENCY_CODE_GTQ  CurrencyCode = 49
	CurrencyCode_CURRENCY_CODE_GYD  CurrencyCode = 50
	CurrencyCode_CURRENCY_CODE_HKD  CurrencyCode = 51
	CurrencyCode_CURRENCY_CODE_HNL  CurrencyCode = 52
	CurrencyCode_CURRENCY_CODE_HRK  CurrencyCode = 53
	CurrencyCode_CURRENCY_CODE_HUF  CurrencyCode = 54
	CurrencyCode_CURRENCY_CODE_IDR  CurrencyCode = 55
	CurrencyCode_CURRENCY_CODE_ILS  CurrencyCode = 56
	CurrencyCode_CURRENCY_CODE_INR  CurrencyCode = 57
	CurrencyCode_CURRENCY_CODE_IQD  CurrencyCode = 58
	CurrencyCode_CURRENCY_CODE_IRR  CurrencyCode = 59
	CurrencyCode_CURRENCY_CODE_ISK  CurrencyCode = 60
	CurrencyCode_CURRENCY_CODE_JMD  CurrencyCode = 61
	CurrencyCode_CURRENCY_CODE_JOD  CurrencyCode = 62
	CurrencyCode_CURRENCY_CODE_JPY  CurrencyCode = 63
	CurrencyCode_CURRENCY_CODE_KES  CurrencyCode = 64
	CurrencyCode_CURRENCY_CODE_KGS  CurrencyCode = 65
	CurrencyCode_CURRENCY_CODE_KHR  CurrencyCode = 66
	CurrencyCode_CURRENCY_CODE_KMF  CurrencyCode = 67
	CurrencyCode_CURRENCY_CODE_KPW  CurrencyCode = 68
	CurrencyCode_CURRENCY_CODE_KRW  CurrencyCode = 69
	CurrencyCode_CURRENCY_CODE_KWD  CurrencyCode = 70
	CurrencyCode_CURRENCY_CODE_KYD  CurrencyCode = 71
	CurrencyCode_CURRENCY_CODE_KZT  CurrencyCode = 72
	CurrencyCode_CURRENCY_CODE_LAK  CurrencyCode = 73
	CurrencyCode_CURRENCY_CODE_LBP  CurrencyCode = 74
	CurrencyCode_CURRENCY_CODE_LKR  CurrencyCode = 75
	CurrencyCode_CURRENCY_CODE_LRD  CurrencyCode = 76
	CurrencyCode_CURRENCY_CODE_LYD  CurrencyCode = 77
	CurrencyCode_CURRENCY_CODE_MAD  CurrencyCode = 78
	CurrencyCode_CURRENCY_CODE_MDL  CurrencyCode = 79
	CurrencyCode_CURRENCY_CODE_MGA  CurrencyCode = 80
	CurrencyCode_CURRENCY_CODE_MKD  CurrencyCode = 81
	CurrencyCode_CURRENCY_CODE_MMK  CurrencyCode = 82
	CurrencyCode_CURRENCY_CODE_MNT  CurrencyCode = 83
	CurrencyCode_CURRENCY_CODE_MOP  CurrencyCode = 84
	CurrencyCode_CURRENCY_CODE_MRO  CurrencyCode = 85
	CurrencyCode_CURRENCY_CODE_MUR  CurrencyCode = 86
	CurrencyCode_CURRENCY_CODE_MVR  CurrencyCode = 87
	CurrencyCode_CURRENCY_CODE_MWK  CurrencyCode = 88
	CurrencyCode_CURRENCY_CODE_MXN  CurrencyCode = 89
	CurrencyCode_CURRENCY_CODE_MYR  CurrencyCode = 90
	CurrencyCode_CURRENCY_CODE_MZN  CurrencyCode = 91
	CurrencyCode_CURRENCY_CODE_NGN  CurrencyCode = 92
	CurrencyCode_CURRENCY_CODE_NIO  CurrencyCode = 93
	CurrencyCode_CURRENCY_CODE_NOK  CurrencyCode = 94
	CurrencyCode_CURRENCY_CODE_NPR  CurrencyCode = 95
	CurrencyCode_CURRENCY_CODE_NZD  CurrencyCode = 96
	CurrencyCode_CURRENCY_CODE_OMR  CurrencyCode = 97
	CurrencyCode_CURRENCY_CODE_PEN  CurrencyCode = 98
	CurrencyCode_CURRENCY_CODE_PGK  CurrencyCode = 99
	CurrencyCode_CURRENCY_CODE_PHP  CurrencyCode = 100
	CurrencyCode_CURRENCY_CODE_PKR  CurrencyCode = 101
	CurrencyCode_CURRENCY_CODE_PLN  CurrencyCode = 102
	CurrencyCode_CURRENCY_CODE_PYG  CurrencyCode = 103
	CurrencyCode_CURRENCY_CODE_QAR  CurrencyCode = 104
	CurrencyCode_CURRENCY_CODE_RON  CurrencyCode = 105
	CurrencyCode_CURRENCY_CODE_RSD  CurrencyCode = 106
	CurrencyCode_CURRENCY_CODE_RUB  CurrencyCode = 107
	CurrencyCode_CURRENCY_CODE_RWF  CurrencyCode = 108
	CurrencyCode_CURRENCY_CODE_SAR  CurrencyCode = 109
	CurrencyCode_CURRENCY_CODE_SBD  CurrencyCode = 110
	CurrencyCode_CURRENCY_CODE_SCR  CurrencyCode = 111
	CurrencyCode_CURRENCY_CODE_SDG  CurrencyCode = 112
	CurrencyCode_CURRENCY_CODE_SEK  CurrencyCode = 113
	CurrencyCode_CURRENCY_CODE_SGD  CurrencyCode = 114
	CurrencyCode_CURRENCY_CODE_SHP  CurrencyCode = 115
	CurrencyCode_CURRENCY_CODE_SLL  CurrencyCode = 116
	CurrencyCode_CURRENCY_CODE_SOS  CurrencyCode = 117
	CurrencyCode_CURRENCY_CODE_SRD  CurrencyCode = 118
	CurrencyCode_CURRENCY_CODE_SSP  CurrencyCode = 119
	CurrencyCode_CURRENCY_CODE_STD  CurrencyCode = 120
	CurrencyCode_CURRENCY_CODE_SYP  CurrencyCode = 121
	CurrencyCode_CURRENCY_CODE_SZL  CurrencyCode = 122
	CurrencyCode_CURRENCY_CODE_THB  CurrencyCode = 123
	CurrencyCode_CURRENCY_CODE_TJS  CurrencyCode = 124
	CurrencyCode_CURRENCY_CODE_TMT  CurrencyCode = 125
	CurrencyCode_CURRENCY_CODE_TND  CurrencyCode = 126
	CurrencyCode_CURRENCY_CODE_TOP  CurrencyCode = 127
	CurrencyCode_CURRENCY_CODE_TRY  CurrencyCode = 128
	CurrencyCode_CURRENCY_CODE_TTD  CurrencyCode = 129
	CurrencyCode_CURRENCY_CODE_TZS  CurrencyCode = 130
	CurrencyCode_CURRENCY_CODE_UAH  CurrencyCode = 131
	CurrencyCode_CURRENCY_CODE_UGX  CurrencyCode = 132
	CurrencyCode_CURRENCY_CODE_USD  CurrencyCode = 133
	CurrencyCode_CURRENCY_CODE_UYU  CurrencyCode = 134
	CurrencyCode_CURRENCY_CODE_UZS  CurrencyCode = 135
	CurrencyCode_CURRENCY_CODE_VEF  CurrencyCode = 136
	CurrencyCode_CURRENCY_CODE_VND  CurrencyCode = 137
	CurrencyCode_CURRENCY_CODE_VUV  CurrencyCode = 138
	CurrencyCode_CURRENCY_CODE_WST  CurrencyCode = 139
	CurrencyCode_CURRENCY_CODE_XAF  CurrencyCode = 140
	CurrencyCode_CURRENCY_CODE_XCD  CurrencyCode = 141
	CurrencyCode_CURRENCY_CODE_XOF  CurrencyCode = 142
	CurrencyCode_CURRENCY_CODE_XPF  CurrencyCode = 143
	CurrencyCode_CURRENCY_CODE_YER  CurrencyCode = 144
	CurrencyCode_CURRENCY_CODE_ZAR  CurrencyCode = 145
	CurrencyCode_CURRENCY_CODE_ZMW  CurrencyCode = 146
	CurrencyCode_CURRENCY_CODE_ZWL  CurrencyCode = 147
)

var CurrencyCode_name = map[int32]string{
	0:   "CURRENCY_CODE_NONE",
	1:   "CURRENCY_CODE_AED",
	2:   "CURRENCY_CODE_AFN",
	3:   "CURRENCY_CODE_ALL",
	4:   "CURRENCY_CODE_AMD",
	5:   "CURRENCY_CODE_ANG",
	6:   "CURRENCY_CODE_AOA",
	7:   "CURRENCY_CODE_ARS",
	8:   "CURRENCY_CODE_AUD",
	9:   "CURRENCY_CODE_AWG",
	10:  "CURRENCY_CODE_AZN",
	11:  "CURRENCY_CODE_BAM",
	12:  "CURRENCY_CODE_BBD",
	13:  "CURRENCY_CODE_BDT",
	14:  "CURRENCY_CODE_BGN",
	15:  "CURRENCY_CODE_BHD",
	16:  "CURRENCY_CODE_BIF",
	17:  "CURRENCY_CODE_BMD",
	18:  "CURRENCY_CODE_BND",
	19:  "CURRENCY_CODE_BOB",
	20:  "CURRENCY_CODE_BRL",
	21:  "CURRENCY_CODE_BSD",
	22:  "CURRENCY_CODE_BWP",
	23:  "CURRENCY_CODE_BYN",
	24:  "CURRENCY_CODE_BZD",
	25:  "CURRENCY_CODE_CAD",
	26:  "CURRENCY_CODE_CDF",
	27:  "CURRENCY_CODE_CHF",
	28:  "CURRENCY_CODE_CLP",
	29:  "CURRENCY_CODE_CNY",
	30:  "CURRENCY_CODE_COP",
	31:  "CURRENCY_CODE_CRC",
	32:  "CURRENCY_CODE_CVE",
	33:  "CURRENCY_CODE_CZK",
	34:  "CURRENCY_CODE_DJF",
	35:  "CURRENCY_CODE_DKK",
	36:  "CURRENCY_CODE_DOP",
	37:  "CURRENCY_CODE_DZD",
	38:  "CURRENCY_CODE_EGP",
	39:  "CURRENCY_CODE_ERN",
	40:  "CURRENCY_CODE_ETB",
	41:  "CURRENCY_CODE_EUR",
	42:  "CURRENCY_CODE_FJD",
	43:  "CURRENCY_CODE_GBP",
	44:  "CURRENCY_CODE_GEL",
	45:  "CURRENCY_CODE_GHS",
	46:  "CURRENCY_CODE_GIP",
	47:  "CURRENCY_CODE_GMD",
	48:  "CURRENCY_CODE_GNF",
	49:  "CURRENCY_CODE_GTQ",
	50:  "CURRENCY_CODE_GYD",
	51:  "CURRENCY_CODE_HKD",
	52:  "CURRENCY_CODE_HNL",
	53:  "CURRENCY_CODE_HRK",
	54:  "CURRENCY_CODE_HUF",
	55:  "CURRENCY_CODE_IDR",
	56:  "CURRENCY_CODE_ILS",
	57:  "CURRENCY_CODE_INR",
	58:  "CURRENCY_CODE_IQD",
	59:  "CURRENCY_CODE_IRR",
	60:  "CURRENCY_CODE_ISK",
	61:  "CURRENCY_CODE_JMD",
	62:  "CURRENCY_CODE_JOD",
	63:  "CURRENCY_CODE_JPY",
	64:  "CURRENCY_CODE_KES",
	65:  "CURRENCY_CODE_KGS",
	66:  "CURRENCY_CODE_KHR",
	67:  "CURRENCY_CODE_KMF",
	68:  "CURRENCY_CODE_KPW",
	69:  "CURRENCY_CODE_KRW",
	70:  "CURRENCY_CODE_KWD",
	71:  "CURRENCY_CODE_KYD",
	72:  "CURRENCY_CODE_KZT",
	73:  "CURRENCY_CODE_LAK",
	74:  "CURRENCY_CODE_LBP",
	75:  "CURRENCY_CODE_LKR",
	76:  "CURRENCY_CODE_LRD",
	77:  "CURRENCY_CODE_LYD",
	78:  "CURRENCY_CODE_MAD",
	79:  "CURRENCY_CODE_MDL",
	80:  "CURRENCY_CODE_MGA",
	81:  "CURRENCY_CODE_MKD",
	82:  "CURRENCY_CODE_MMK",
	83:  "CURRENCY_CODE_MNT",
	84:  "CURRENCY_CODE_MOP",
	85:  "CURRENCY_CODE_MRO",
	86:  "CURRENCY_CODE_MUR",
	87:  "CURRENCY_CODE_MVR",
	88:  "CURRENCY_CODE_MWK",
	89:  "CURRENCY_CODE_MXN",
	90:  "CURRENCY_CODE_MYR",
	91:  "CURRENCY_CODE_MZN",
	92:  "CURRENCY_CODE_NGN",
	93:  "CURRENCY_CODE_NIO",
	94:  "CURRENCY_CODE_NOK",
	95:  "CURRENCY_CODE_NPR",
	96:  "CURRENCY_CODE_NZD",
	97:  "CURRENCY_CODE_OMR",
	98:  "CURRENCY_CODE_PEN",
	99:  "CURRENCY_CODE_PGK",
	100: "CURRENCY_CODE_PHP",
	101: "CURRENCY_CODE_PKR",
	102: "CURRENCY_CODE_PLN",
	103: "CURRENCY_CODE_PYG",
	104: "CURRENCY_CODE_QAR",
	105: "CURRENCY_CODE_RON",
	106: "CURRENCY_CODE_RSD",
	107: "CURRENCY_CODE_RUB",
	108: "CURRENCY_CODE_RWF",
	109: "CURRENCY_CODE_SAR",
	110: "CURRENCY_CODE_SBD",
	111: "CURRENCY_CODE_SCR",
	112: "CURRENCY_CODE_SDG",
	113: "CURRENCY_CODE_SEK",
	114: "CURRENCY_CODE_SGD",
	115: "CURRENCY_CODE_SHP",
	116: "CURRENCY_CODE_SLL",
	117: "CURRENCY_CODE_SOS",
	118: "CURRENCY_CODE_SRD",
	119: "CURRENCY_CODE_SSP",
	120: "CURRENCY_CODE_STD",
	121: "CURRENCY_CODE_SYP",
	122: "CURRENCY_CODE_SZL",
	123: "CURRENCY_CODE_THB",
	124: "CURRENCY_CODE_TJS",
	125: "CURRENCY_CODE_TMT",
	126: "CURRENCY_CODE_TND",
	127: "CURRENCY_CODE_TOP",
	128: "CURRENCY_CODE_TRY",
	129: "CURRENCY_CODE_TTD",
	130: "CURRENCY_CODE_TZS",
	131: "CURRENCY_CODE_UAH",
	132: "CURRENCY_CODE_UGX",
	133: "CURRENCY_CODE_USD",
	134: "CURRENCY_CODE_UYU",
	135: "CURRENCY_CODE_UZS",
	136: "CURRENCY_CODE_VEF",
	137: "CURRENCY_CODE_VND",
	138: "CURRENCY_CODE_VUV",
	139: "CURRENCY_CODE_WST",
	140: "CURRENCY_CODE_XAF",
	141: "CURRENCY_CODE_XCD",
	142: "CURRENCY_CODE_XOF",
	143: "CURRENCY_CODE_XPF",
	144: "CURRENCY_CODE_YER",
	145: "CURRENCY_CODE_ZAR",
	146: "CURRENCY_CODE_ZMW",
	147: "CURRENCY_CODE_ZWL",
}
var CurrencyCode_value = map[string]int32{
	"CURRENCY_CODE_NONE": 0,
	"CURRENCY_CODE_AED":  1,
	"CURRENCY_CODE_AFN":  2,
	"CURRENCY_CODE_ALL":  3,
	"CURRENCY_CODE_AMD":  4,
	"CURRENCY_CODE_ANG":  5,
	"CURRENCY_CODE_AOA":  6,
	"CURRENCY_CODE_ARS":  7,
	"CURRENCY_CODE_AUD":  8,
	"CURRENCY_CODE_AWG":  9,
	"CURRENCY_CODE_AZN":  10,
	"CURRENCY_CODE_BAM":  11,
	"CURRENCY_CODE_BBD":  12,
	"CURRENCY_CODE_BDT":  13,
	"CURRENCY_CODE_BGN":  14,
	"CURRENCY_CODE_BHD":  15,
	"CURRENCY_CODE_BIF":  16,
	"CURRENCY_CODE_BMD":  17,
	"CURRENCY_CODE_BND":  18,
	"CURRENCY_CODE_BOB":  19,
	"CURRENCY_CODE_BRL":  20,
	"CURRENCY_CODE_BSD":  21,
	"CURRENCY_CODE_BWP":  22,
	"CURRENCY_CODE_BYN":  23,
	"CURRENCY_CODE_BZD":  24,
	"CURRENCY_CODE_CAD":  25,
	"CURRENCY_CODE_CDF":  26,
	"CURRENCY_CODE_CHF":  27,
	"CURRENCY_CODE_CLP":  28,
	"CURRENCY_CODE_CNY":  29,
	"CURRENCY_CODE_COP":  30,
	"CURRENCY_CODE_CRC":  31,
	"CURRENCY_CODE_CVE":  32,
	"CURRENCY_CODE_CZK":  33,
	"CURRENCY_CODE_DJF":  34,
	"CURRENCY_CODE_DKK":  35,
	"CURRENCY_CODE_DOP":  36,
	"CURRENCY_CODE_DZD":  37,
	"CURRENCY_CODE_EGP":  38,
	"CURRENCY_CODE_ERN":  39,
	"CURRENCY_CODE_ETB":  40,
	"CURRENCY_CODE_EUR":  41,
	"CURRENCY_CODE_FJD":  42,
	"CURRENCY_CODE_GBP":  43,
	"CURRENCY_CODE_GEL":  44,
	"CURRENCY_CODE_GHS":  45,
	"CURRENCY_CODE_GIP":  46,
	"CURRENCY_CODE_GMD":  47,
	"CURRENCY_CODE_GNF":  48,
	"CURRENCY_CODE_GTQ":  49,
	"CURRENCY_CODE_GYD":  50,
	"CURRENCY_CODE_HKD":  51,
	"CURRENCY_CODE_HNL":  52,
	"CURRENCY_CODE_HRK":  53,
	"CURRENCY_CODE_HUF":  54,
	"CURRENCY_CODE_IDR":  55,
	"CURRENCY_CODE_ILS":  56,
	"CURRENCY_CODE_INR":  57,
	"CURRENCY_CODE_IQD":  58,
	"CURRENCY_CODE_IRR":  59,
	"CURRENCY_CODE_ISK":  60,
	"CURRENCY_CODE_JMD":  61,
	"CURRENCY_CODE_JOD":  62,
	"CURRENCY_CODE_JPY":  63,
	"CURRENCY_CODE_KES":  64,
	"CURRENCY_CODE_KGS":  65,
	"CURRENCY_CODE_KHR":  66,
	"CURRENCY_CODE_KMF":  67,
	"CURRENCY_CODE_KPW":  68,
	"CURRENCY_CODE_KRW":  69,
	"CURRENCY_CODE_KWD":  70,
	"CURRENCY_CODE_KYD":  71,
	"CURRENCY_CODE_KZT":  72,
	"CURRENCY_CODE_LAK":  73,
	"CURRENCY_CODE_LBP":  74,
	"CURRENCY_CODE_LKR":  75,
	"CURRENCY_CODE_LRD":  76,
	"CURRENCY_CODE_LYD":  77,
	"CURRENCY_CODE_MAD":  78,
	"CURRENCY_CODE_MDL":  79,
	"CURRENCY_CODE_MGA":  80,
	"CURRENCY_CODE_MKD":  81,
	"CURRENCY_CODE_MMK":  82,
	"CURRENCY_CODE_MNT":  83,
	"CURRENCY_CODE_MOP":  84,
	"CURRENCY_CODE_MRO":  85,
	"CURRENCY_CODE_MUR":  86,
	"CURRENCY_CODE_MVR":  87,
	"CURRENCY_CODE_MWK":  88,
	"CURRENCY_CODE_MXN":  89,
	"CURRENCY_CODE_MYR":  90,
	"CURRENCY_CODE_MZN":  91,
	"CURRENCY_CODE_NGN":  92,
	"CURRENCY_CODE_NIO":  93,
	"CURRENCY_CODE_NOK":  94,
	"CURRENCY_CODE_NPR":  95,
	"CURRENCY_CODE_NZD":  96,
	"CURRENCY_CODE_OMR":  97,
	"CURRENCY_CODE_PEN":  98,
	"CURRENCY_CODE_PGK":  99,
	"CURRENCY_CODE_PHP":  100,
	"CURRENCY_CODE_PKR":  101,
	"CURRENCY_CODE_PLN":  102,
	"CURRENCY_CODE_PYG":  103,
	"CURRENCY_CODE_QAR":  104,
	"CURRENCY_CODE_RON":  105,
	"CURRENCY_CODE_RSD":  106,
	"CURRENCY_CODE_RUB":  107,
	"CURRENCY_CODE_RWF":  108,
	"CURRENCY_CODE_SAR":  109,
	"CURRENCY_CODE_SBD":  110,
	"CURRENCY_CODE_SCR":  111,
	"CURRENCY_CODE_SDG":  112,
	"CURRENCY_CODE_SEK":  113,
	"CURRENCY_CODE_SGD":  114,
	"CURRENCY_CODE_SHP":  115,
	"CURRENCY_CODE_SLL":  116,
	"CURRENCY_CODE_SOS":  117,
	"CURRENCY_CODE_SRD":  118,
	"CURRENCY_CODE_SSP":  119,
	"CURRENCY_CODE_STD":  120,
	"CURRENCY_CODE_SYP":  121,
	"CURRENCY_CODE_SZL":  122,
	"CURRENCY_CODE_THB":  123,
	"CURRENCY_CODE_TJS":  124,
	"CURRENCY_CODE_TMT":  125,
	"CURRENCY_CODE_TND":  126,
	"CURRENCY_CODE_TOP":  127,
	"CURRENCY_CODE_TRY":  128,
	"CURRENCY_CODE_TTD":  129,
	"CURRENCY_CODE_TZS":  130,
	"CURRENCY_CODE_UAH":  131,
	"CURRENCY_CODE_UGX":  132,
	"CURRENCY_CODE_USD":  133,
	"CURRENCY_CODE_UYU":  134,
	"CURRENCY_CODE_UZS":  135,
	"CURRENCY_CODE_VEF":  136,
	"CURRENCY_CODE_VND":  137,
	"CURRENCY_CODE_VUV":  138,
	"CURRENCY_CODE_WST":  139,
	"CURRENCY_CODE_XAF":  140,
	"CURRENCY_CODE_XCD":  141,
	"CURRENCY_CODE_XOF":  142,
	"CURRENCY_CODE_XPF":  143,
	"CURRENCY_CODE_YER":  144,
	"CURRENCY_CODE_ZAR":  145,
	"CURRENCY_CODE_ZMW":  146,
	"CURRENCY_CODE_ZWL":  147,
}

func (x CurrencyCode) String() string {
	return proto.EnumName(CurrencyCode_name, int32(x))
}
func (CurrencyCode) EnumDescriptor() ([]byte, []int) { return fileDescriptorMoneyGen, []int{0} }

func init() {
	proto.RegisterEnum("pb.money.CurrencyCode", CurrencyCode_name, CurrencyCode_value)
}

func init() { proto.RegisterFile("pb/money/money.gen.proto", fileDescriptorMoneyGen) }

var fileDescriptorMoneyGen = []byte{
	// 773 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xd7, 0x57, 0x76, 0xdd, 0x54,
	0x18, 0x47, 0x71, 0x42, 0x49, 0x31, 0x01, 0x76, 0x2e, 0xc4, 0x84, 0xd0, 0x3b, 0x04, 0x70, 0x80,
	0xd0, 0xbb, 0x74, 0x3f, 0x49, 0xd7, 0x96, 0x74, 0x24, 0x1f, 0x49, 0x96, 0x25, 0x8a, 0xc1, 0x89,
	0x09, 0x2d, 0xb6, 0x31, 0x09, 0x60, 0x7a, 0xef, 0x9d, 0x81, 0x32, 0x05, 0xd6, 0x32, 0x4f, 0xac,
	0xf3, 0x7f, 0xd9, 0x0f, 0xbf, 0x19, 0xec, 0xb9, 0x63, 0xdb, 0xeb, 0x27, 0xcf, 0x6d, 0x6d, 0x6e,
	0xec, 0xfe, 0xd7, 0x85, 0xb3, 0x1b, 0x9b, 0x0b, 0xdb, 0x3b, 0x5b, 0xe7, 0xb7, 0x26, 0x07, 0xb7,
	0xd7, 0x17, 0xf6, 0xec, 0xc4, 0x3f, 0xc7, 0xe7, 0x0e, 0x4f, 0x2f, 0xec, 0xec, 0x6c, 0x6c, 0x9e,
	0xde, 0x9d, 0x6e, 0x9d, 0xd9, 0x98, 0xcc, 0xcf, 0x4d, 0xa6, 0x9d, 0xf7, 0x89, 0x9b, 0x0e, 0x6b,
	0xd3, 0xca, 0x92, 0x35, 0x57, 0xb9, 0x84, 0x8b, 0x26, 0x47, 0xe7, 0x8e, 0xfc, 0xdf, 0xa3, 0xc4,
	0xd8, 0x27, 0x38, 0x75, 0x5c, 0x2c, 0xb8, 0x28, 0xb8, 0x44, 0x70, 0x69, 0x5c, 0x2a, 0xd8, 0x65,
	0x5c, 0x26, 0xb8, 0x8a, 0xd8, 0x2f, 0xd8, 0x37, 0x1c, 0x10, 0xdc, 0x19, 0x07, 0x05, 0xf7, 0x19,
	0x87, 0x04, 0x8f, 0x8e, 0xb9, 0x90, 0xe3, 0xa8, 0xe4, 0x72, 0xc1, 0xb1, 0x71, 0x58, 0xb0, 0xb5,
	0x5c, 0x21, 0x38, 0x73, 0x5c, 0x29, 0x78, 0x66, 0x5c, 0x25, 0x78, 0x31, 0x05, 0xc1, 0xa5, 0x71,
	0x44, 0xb0, 0x33, 0x26, 0x82, 0xab, 0x98, 0xab, 0x05, 0xfb, 0x82, 0x6b, 0x04, 0x37, 0xc6, 0x51,
	0xc1, 0x7d, 0xcd, 0xbc, 0xe0, 0xc1, 0x71, 0xad, 0xe0, 0xd1, 0x38, 0x16, 0xf2, 0x34, 0x32, 0xae,
	0x13, 0x6c, 0x29, 0xc7, 0x05, 0xcf, 0x52, 0xae, 0x17, 0x5c, 0xd4, 0xdc, 0x20, 0xd8, 0x0d, 0xdc,
	0x28, 0xb8, 0xaa, 0xb9, 0x49, 0xb0, 0x9f, 0x72, 0xb3, 0xe0, 0x95, 0x84, 0x5b, 0x04, 0x8f, 0x39,
	0xb7, 0x86, 0x6c, 0x4b, 0x29, 0xb7, 0x09, 0xce, 0x73, 0x6e, 0x17, 0x5c, 0xd5, 0xdc, 0x21, 0x78,
	0x34, 0xee, 0x0c, 0x39, 0xc9, 0x6a, 0xee, 0x12, 0xec, 0x1d, 0x77, 0x0b, 0x6e, 0x63, 0xee, 0x11,
	0xdc, 0x79, 0xee, 0x0d, 0x39, 0x5d, 0x32, 0x4e, 0x84, 0x9c, 0xc5, 0x35, 0xf7, 0x09, 0x4e, 0x0a,
	0xee, 0x17, 0x3c, 0x6b, 0x78, 0x40, 0xf0, 0x62, 0xcd, 0x82, 0xe0, 0xd2, 0x38, 0x29, 0xd8, 0xa5,
	0x3c, 0x28, 0xb8, 0x5d, 0xe6, 0x21, 0xc1, 0x83, 0xf1, 0x70, 0xc8, 0xb3, 0xdc, 0x38, 0x25, 0xd8,
	0x15, 0x3c, 0x22, 0xd8, 0xe7, 0x3c, 0x2a, 0xb8, 0x4b, 0x79, 0x2c, 0xe4, 0x45, 0xf3, 0x3c, 0x2e,
	0xb8, 0x68, 0x78, 0x42, 0xb0, 0xf3, 0x3c, 0x29, 0x78, 0xd9, 0x78, 0x4a, 0xb0, 0xf7, 0x3c, 0x2d,
	0xb8, 0xc9, 0x79, 0x26, 0xe4, 0xa5, 0xd2, 0x78, 0x56, 0x70, 0x65, 0x3c, 0x27, 0xb8, 0x1e, 0x78,
	0x3e, 0xe4, 0x3c, 0x69, 0x78, 0x41, 0x70, 0xd6, 0x10, 0x09, 0x9e, 0x79, 0x62, 0xc1, 0x65, 0xca,
	0x54, 0x70, 0xdd, 0x63, 0x82, 0x7d, 0x4f, 0x22, 0xb8, 0x37, 0x52, 0xc1, 0x83, 0x91, 0x09, 0x1e,
	0x5b, 0x66, 0x21, 0x17, 0x51, 0xce, 0xa2, 0xe0, 0xb8, 0x66, 0x49, 0x70, 0xee, 0xc9, 0x05, 0x7b,
	0xa3, 0x10, 0x3c, 0x18, 0x65, 0xc8, 0x65, 0x64, 0x38, 0xc1, 0x56, 0x50, 0x09, 0xce, 0x22, 0x6a,
	0xc1, 0xb9, 0xb1, 0x2c, 0xb8, 0xcc, 0xf1, 0x82, 0x5d, 0x4b, 0x23, 0xb8, 0xaa, 0x69, 0x05, 0xfb,
	0x8a, 0x4e, 0x70, 0xe7, 0x59, 0x11, 0xbc, 0xe2, 0xe9, 0x05, 0xf7, 0x39, 0xab, 0x82, 0x57, 0x1d,
	0x83, 0xe0, 0xc1, 0x33, 0x0a, 0x1e, 0x1d, 0x2f, 0x86, 0xec, 0x32, 0xc7, 0x4b, 0x82, 0x17, 0x2b,
	0x5e, 0x16, 0x5c, 0xe5, 0xbc, 0x22, 0xb8, 0xf6, 0xac, 0x09, 0x1e, 0x8d, 0x57, 0x43, 0xae, 0x4a,
	0xcf, 0x6b, 0x21, 0xd7, 0x89, 0x63, 0x5d, 0x70, 0x96, 0x73, 0x5a, 0xf0, 0xac, 0xe6, 0x8c, 0xe0,
	0xdc, 0xb3, 0x21, 0xb8, 0x70, 0xbc, 0x2e, 0x78, 0xc8, 0x38, 0x1b, 0xf2, 0x72, 0xe4, 0x79, 0x23,
	0x64, 0x5f, 0x39, 0xde, 0x14, 0xdc, 0x18, 0x6f, 0x09, 0xee, 0x62, 0xde, 0x16, 0xdc, 0xa7, 0xbc,
	0x13, 0x72, 0x13, 0x79, 0xce, 0x09, 0x8e, 0x8d, 0x4d, 0xc1, 0x53, 0xcf, 0x96, 0x60, 0xcb, 0xd8,
	0x16, 0x9c, 0xe4, 0xbc, 0x2b, 0x38, 0x33, 0x76, 0x04, 0xcf, 0x6a, 0xde, 0x13, 0x5c, 0x14, 0x9c,
	0x17, 0x5c, 0x35, 0x5c, 0x10, 0xec, 0x8d, 0xf7, 0x05, 0x37, 0x35, 0x1f, 0x08, 0x6e, 0x8d, 0x0f,
	0x05, 0x0f, 0x35, 0xbb, 0x82, 0xc7, 0x82, 0x8f, 0x42, 0x6e, 0x67, 0x31, 0x1f, 0x0b, 0x5e, 0x6a,
	0xf8, 0x44, 0x70, 0xd9, 0xf2, 0xa9, 0x60, 0x67, 0x7c, 0x26, 0xb8, 0xaa, 0xf9, 0x7c, 0x32, 0x1f,
	0xb0, 0x1f, 0xf8, 0x62, 0x9f, 0xf0, 0xd6, 0xf8, 0x52, 0xf9, 0xd8, 0xf0, 0x95, 0xf0, 0x2e, 0x9a,
	0xf1, 0xb5, 0xf2, 0x6c, 0x95, 0x6f, 0x94, 0x37, 0xc6, 0xb7, 0xca, 0x87, 0x8e, 0xef, 0x94, 0x8f,
	0x0d, 0xdf, 0x0b, 0x5f, 0x49, 0x52, 0x7e, 0x50, 0xee, 0x8c, 0x1f, 0x95, 0x77, 0x2b, 0xfc, 0x24,
	0xbc, 0x6f, 0x5a, 0x7e, 0x16, 0xbe, 0x1a, 0xa5, 0xfc, 0xa2, 0x7c, 0x6a, 0xfc, 0xaa, 0xbc, 0x4a,
	0xf9, 0x4d, 0x79, 0x9d, 0xf2, 0xbb, 0xf0, 0x21, 0xf1, 0xfc, 0x21, 0x7c, 0x8c, 0x3c, 0x7f, 0x2a,
	0x2f, 0x7b, 0xfe, 0x52, 0xde, 0x17, 0xfc, 0xbd, 0x2f, 0x3e, 0x34, 0x1e, 0xd8, 0x5e, 0xdf, 0x9b,
	0xaf, 0xf5, 0xfd, 0x7b, 0x37, 0x76, 0xea, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0xbe, 0x2d,
	0xe7, 0xa9, 0x0d, 0x00, 0x00,
}