package stacks

type ChainID uint32

const (
	ChainIDTestnet ChainID = 0x80000000
	ChainIDMainnet ChainID = 0x00000001
)

type PayloadType byte

const (
	PayloadTypeTokenTransfer          PayloadType = 0x00
	PayloadTypeSmartContract          PayloadType = 0x01
	PayloadTypeContractCall           PayloadType = 0x02
	PayloadTypePoisonMicroblock       PayloadType = 0x03
	PayloadTypeCoinbase               PayloadType = 0x04
	PayloadTypeCoinbaseToAltRecipient PayloadType = 0x05
	PayloadTypeVersionedSmartContract PayloadType = 0x06
	PayloadTypeTenureChange           PayloadType = 0x07
	PayloadTypeNakamotoCoinbase       PayloadType = 0x08
)

type AddressType byte

const (
	AddressTypeStandard AddressType = 0x05
	AddressTypeContract AddressType = 0x06
)

type AnchorMode uint8

const (
	AnchorModeOnChainOnly  AnchorMode = 0x01
	AnchorModeOffChainOnly AnchorMode = 0x02
	AnchorModeAny          AnchorMode = 0x03
)

type TransactionVersion uint8

const (
	TransactionVersionMainnet TransactionVersion = 0x00
	TransactionVersionTestnet TransactionVersion = 0x80
)

type PostConditionMode uint8

const (
	PostConditionModeAllow PostConditionMode = 0x01
	PostConditionModeDeny  PostConditionMode = 0x02
)

type PostConditionType uint8

const (
	PostConditionTypeSTX         PostConditionType = 0x00
	PostConditionTypeFungible    PostConditionType = 0x01
	PostConditionTypeNonFungible PostConditionType = 0x02
)

type PostConditionPrincipalType uint8

const (
	PostConditionPrincipalTypeOrigin   PostConditionPrincipalType = 0x01
	PostConditionPrincipalTypeStandard PostConditionPrincipalType = 0x02
	PostConditionPrincipalTypeContract PostConditionPrincipalType = 0x03
)

type FungibleConditionCode uint8

const (
	FungibleConditionCodeSentEq FungibleConditionCode = 0x01
	FungibleConditionCodeSentGt FungibleConditionCode = 0x02
	FungibleConditionCodeSentGe FungibleConditionCode = 0x03
	FungibleConditionCodeSentLt FungibleConditionCode = 0x04
	FungibleConditionCodeSentLe FungibleConditionCode = 0x05
)

type NonFungibleConditionCode uint8

const (
	NonFungibleConditionCodeSent    NonFungibleConditionCode = 0x10
	NonFungibleConditionCodeNotSent NonFungibleConditionCode = 0x11
)

type AuthType uint8

const (
	AuthTypeStandard  AuthType = 0x04
	AuthTypeSponsored AuthType = 0x05
)

type AddressHashMode uint8

const (
	AddressHashModeSerializeP2PKH  AddressHashMode = 0x00
	AddressHashModeSerializeP2WPKH AddressHashMode = 0x02
)

type AddressVersion uint8

const (
	AddressVersionMainnetSingleSig AddressVersion = 22
	AddressVersionTestnetSingleSig AddressVersion = 26
	AddressVersionMainnetMultiSig  AddressVersion = 20
	AddressVersionTestnetMultiSig  AddressVersion = 21
)

type ClarityVersion uint8

const (
	ClarityVersionUnspecified ClarityVersion = 0
	ClarityVersion1           ClarityVersion = 1
	ClarityVersion2           ClarityVersion = 2
	ClarityVersion3           ClarityVersion = 3
)

type PubKeyEncoding uint8

const (
	PubKeyEncodingCompressed   PubKeyEncoding = 0x00
	PubKeyEncodingUncompressed PubKeyEncoding = 0x01
)

const (
	MaxStringLengthBytes           = 128
	ClarityIntSize                 = 128
	ClarityIntByteSize             = 16
	RecoverableECDSASigLengthBytes = 65
	CompressedPubkeyLengthBytes    = 32
	UncompressedPubkeyLengthBytes  = 64
	MemoMaxLengthBytes             = 34
	AddressHashLength              = 20
)
