// Code generated by "stringer -output=rule_name_string.go -type=RuleName rule_name.go rule_name.og.go"; DO NOT EDIT.

package opt

import "strconv"

const _RuleName_name = "InvalidRuleNameNumManualRuleNamesEliminateEmptyAndEliminateEmptyOrEliminateSingletonAndOrSimplifyAndSimplifyOrSimplifyFiltersFoldNullAndOrNegateComparisonEliminateNotNegateAndNegateOrCommuteVarInequalityCommuteConstInequalityNormalizeCmpPlusConstNormalizeCmpMinusConstNormalizeCmpConstMinusNormalizeTupleEqualityFoldNullComparisonLeftFoldNullComparisonRightEnsureJoinFiltersAndEnsureJoinFiltersPushFilterIntoJoinLeftPushFilterIntoJoinRightPushLimitIntoProjectPushOffsetIntoProjectFoldPlusZeroFoldZeroPlusFoldMinusZeroFoldMultOneFoldOneMultFoldDivOneInvertMinusEliminateUnaryMinusEliminateProjectEliminateProjectProjectFilterUnusedProjectColsFilterUnusedScanColsFilterUnusedSelectColsFilterUnusedLimitColsFilterUnusedOffsetColsFilterUnusedJoinLeftColsFilterUnusedJoinRightColsFilterUnusedAggColsFilterUnusedGroupByColsFilterUnusedValueColsCommuteVarCommuteConstEliminateCoalesceSimplifyCoalesceEliminateCastFoldNullCastFoldNullUnaryFoldNullBinaryLeftFoldNullBinaryRightFoldNullInNonEmptyFoldNullInEmptyFoldNullNotInEmptyNormalizeInConstFoldInNullEnsureSelectFiltersAndEnsureSelectFiltersEliminateSelectMergeSelectsPushSelectIntoProjectPushSelectIntoJoinLeftPushSelectIntoJoinRightMergeSelectInnerJoinPushSelectIntoGroupByGenerateIndexScansNumRuleNames"

var _RuleName_index = [...]uint16{0, 15, 33, 50, 66, 89, 100, 110, 125, 138, 154, 166, 175, 183, 203, 225, 246, 268, 290, 312, 334, 357, 377, 394, 416, 439, 459, 480, 492, 504, 517, 528, 539, 549, 560, 579, 595, 618, 641, 661, 683, 704, 726, 750, 775, 794, 817, 838, 848, 860, 877, 893, 906, 918, 931, 949, 968, 986, 1001, 1019, 1035, 1045, 1067, 1086, 1101, 1113, 1134, 1156, 1179, 1199, 1220, 1238, 1250}

func (i RuleName) String() string {
	if i >= RuleName(len(_RuleName_index)-1) {
		return "RuleName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RuleName_name[_RuleName_index[i]:_RuleName_index[i+1]]
}