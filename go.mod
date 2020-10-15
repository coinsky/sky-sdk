module github.com/coinsky/sky-sdk

go 1.13

require (
	github.com/DataDog/zstd v1.4.0 // indirect
	github.com/Shopify/sarama v1.23.1
	github.com/coinexchain/cosmos-utils v0.0.0-20200109031554-f15ba3b1d6a7
	github.com/coinexchain/shorthanzi v0.1.0
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/emirpasic/gods v1.12.0
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/gorilla/mux v1.7.3
	github.com/pelletier/go-toml v1.6.0
	github.com/pierrec/lz4 v2.0.5+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.6 // indirect
	github.com/spf13/cobra v0.0.6
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.5.1
	github.com/tendermint/tendermint v0.33.8
	github.com/tendermint/tm-db v0.4.0
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413
	gopkg.in/jcmturner/goidentity.v3 v3.0.0 // indirect
)

replace github.com/cosmos/cosmos-sdk => github.com/coinsky/cosmos-sdk v0.0.0-20200922151007-53059ac63a8a

replace github.com/tendermint/tendermint => github.com/coinsky/tendermint v0.0.0-20200922144924-fc894ae41af2
