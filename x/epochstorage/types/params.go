package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyUnstakeHoldBlocks = []byte("UnstakeHoldBlocks")
	// TODO: Determine the default value
	DefaultUnstakeHoldBlocks uint64 = 210
)

var (
	KeyEpochBlocks = []byte("EpochBlocks")
	// TODO: Determine the default value
	DefaultEpochBlocks uint64 = 20
)

var (
	KeyEpochsToSave = []byte("EpochsToSave")
	// TODO: Determine the default value
	DefaultEpochsToSave uint64 = 10
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	unstakeHoldBlocks uint64,
	epochBlocks uint64,
	epochsToSave uint64,
) Params {
	return Params{
		UnstakeHoldBlocks: unstakeHoldBlocks,
		EpochBlocks:       epochBlocks,
		EpochsToSave:      epochsToSave,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultUnstakeHoldBlocks,
		DefaultEpochBlocks,
		DefaultEpochsToSave,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyUnstakeHoldBlocks, &p.UnstakeHoldBlocks, validateUnstakeHoldBlocks),
		paramtypes.NewParamSetPair(KeyEpochBlocks, &p.EpochBlocks, validateEpochBlocks),
		paramtypes.NewParamSetPair(KeyEpochsToSave, &p.EpochsToSave, validateEpochsToSave),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateUnstakeHoldBlocks(p.UnstakeHoldBlocks); err != nil {
		return err
	}

	if err := validateEpochBlocks(p.EpochBlocks); err != nil {
		return err
	}

	if err := validateEpochsToSave(p.EpochsToSave); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateUnstakeHoldBlocks validates the UnstakeHoldBlocks param
func validateUnstakeHoldBlocks(v interface{}) error {
	unstakeHoldBlocks, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = unstakeHoldBlocks

	return nil
}

// validateEpochBlocks validates the EpochBlocks param
func validateEpochBlocks(v interface{}) error {
	epochBlocks, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = epochBlocks

	return nil
}

// validateEpochsToSave validates the EpochsToSave param
func validateEpochsToSave(v interface{}) error {
	epochsToSave, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = epochsToSave

	return nil
}