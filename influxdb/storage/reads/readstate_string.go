// Code generated by "stringer -type=readState -trimprefix=state"; DO NOT EDIT.

package reads

import "strconv"

const _readState_name = "ReadGroupReadSeriesReadPointsReadFloatPointsReadIntegerPointsReadUnsignedPointsReadBooleanPointsReadStringPointsReadErrDone"

var _readState_index = [...]uint8{0, 9, 19, 29, 44, 61, 79, 96, 112, 119, 123}

func (i readState) String() string {
	if i >= readState(len(_readState_index)-1) {
		return "readState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _readState_name[_readState_index[i]:_readState_index[i+1]]
}
