// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

/*
Register all entry-points in pdh.dll.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.

Note that all dll exported named entry-points are listed,
including those that are undocumented by the vendor.
*/
package pdh

import "github.com/tHinqa/outside"

func init() {
	outside.AddEPs("pdh.dll", false, EntryPoints)
	outside.AddEPs("pdh.dll", true, UnicodeEntryPoints)
}

//TODO(t): Check Ws with no A counterparts and vv

var EntryPoints = outside.EPs{
	"PdhAdd009CounterA",
	"PdhAddCounterA",
	"PdhBindInputDataSourceA",
	"PdhBrowseCountersA",
	"PdhBrowseCountersHA",
	"PdhCalculateCounterFromRawValue",
	"PdhCloseLog",
	"PdhCloseQuery",
	"PdhCollectQueryData",
	"PdhCollectQueryDataEx",
	"PdhComputeCounterStatistics",
	"PdhConnectMachineA",
	"PdhCreateSQLTablesA",
	"PdhEnumLogSetNamesA",
	"PdhEnumMachinesA",
	"PdhEnumMachinesHA",
	"PdhEnumObjectItemsA",
	"PdhEnumObjectItemsHA",
	"PdhEnumObjectsA",
	"PdhEnumObjectsHA",
	"PdhExpandCounterPathA",
	"PdhExpandWildCardPathA",
	"PdhExpandWildCardPathHA",
	"PdhFormatFromRawValue",
	"PdhGetCounterInfoA",
	"PdhGetCounterTimeBase",
	"PdhGetDataSourceTimeRangeA",
	"PdhGetDataSourceTimeRangeH",
	"PdhGetDefaultPerfCounterA",
	"PdhGetDefaultPerfCounterHA",
	"PdhGetDefaultPerfObjectA",
	"PdhGetDefaultPerfObjectHA",
	"PdhGetDllVersion",
	"PdhGetFormattedCounterArrayA",
	"PdhGetFormattedCounterValue",
	"PdhGetLogFileSize",
	"PdhGetLogFileTypeA",
	"PdhGetLogSetGUID",
	"PdhGetRawCounterArrayA",
	"PdhGetRawCounterValue",
	"PdhIsRealTimeQuery",
	"PdhListLogFileHeaderA",
	"PdhLogServiceCommandA",
	"PdhLogServiceControlA",
	"PdhLookupPerfIndexByNameA",
	"PdhLookupPerfNameByIndexA",
	"PdhMakeCounterPathA",
	"PdhOpenLogA",
	"PdhOpenQuery",
	"PdhOpenQueryA",
	"PdhOpenQueryH",
	"PdhParseCounterPathA",
	"PdhParseInstanceNameA",
	"PdhPlaAddItemA",
	"PdhPlaCreateA",
	"PdhPlaDeleteA",
	"PdhPlaEnumCollectionsA",
	"PdhPlaGetInfoA",
	"PdhPlaGetLogFileNameA",
	"PdhPlaGetScheduleA",
	"PdhPlaRemoveAllItemsA",
	"PdhPlaScheduleA",
	"PdhPlaSetInfoA",
	"PdhPlaSetItemListA",
	"PdhPlaSetRunAsA",
	"PdhPlaStartA",
	"PdhPlaStopA",
	"PdhPlaValidateInfoA",
	"PdhReadRawLogRecord",
	"PdhRelogA",
	"PdhRemoveCounter",
	"PdhSelectDataSourceA",
	"PdhSetCounterScaleFactor",
	"PdhSetDefaultRealTimeDataSource",
	"PdhSetLogSetRunID",
	"PdhSetQueryTimeRange",
	"PdhTranslate009CounterA",
	"PdhTranslateLocaleCounterA",
	"PdhUpdateLogA",
	"PdhUpdateLogFileCatalog",
	"PdhValidatePathA",
	"PdhVbAddCounter",
	"PdhVbCreateCounterPathList",
	"PdhVbGetCounterPathElements",
	"PdhVbGetCounterPathFromList",
	"PdhVbGetDoubleCounterValue",
	"PdhVbGetLogFileSize",
	"PdhVbGetOneCounterPath",
	"PdhVbIsGoodStatus",
	"PdhVbOpenLog",
	"PdhVbOpenQuery",
	"PdhVbUpdateLog",
	"PdhVerifySQLDBA",
	"PdhiPla2003SP1Installed",
	"PdhiPlaFormatBlanksA",
	"PdhiPlaGetVersion",
	"PdhiPlaRunAs",
	"PdhiPlaSetRunAs",
	"PlaTimeInfoToMilliSeconds",
}

var UnicodeEntryPoints = outside.EPs{
	"PdhAdd009CounterW",
	"PdhAddCounterW",
	"PdhBindInputDataSourceW",
	"PdhBrowseCountersHW",
	"PdhBrowseCountersW",
	"PdhConnectMachineW",
	"PdhCreateSQLTablesW",
	"PdhEnumLogSetNamesW",
	"PdhEnumMachinesHW",
	"PdhEnumMachinesW",
	"PdhEnumObjectItemsHW",
	"PdhEnumObjectItemsW",
	"PdhEnumObjectsHW",
	"PdhEnumObjectsW",
	"PdhExpandCounterPathW",
	"PdhExpandWildCardPathHW",
	"PdhExpandWildCardPathW",
	"PdhGetCounterInfoW",
	"PdhGetDataSourceTimeRangeW",
	"PdhGetDefaultPerfCounterHW",
	"PdhGetDefaultPerfCounterW",
	"PdhGetDefaultPerfObjectHW",
	"PdhGetDefaultPerfObjectW",
	"PdhGetFormattedCounterArrayW",
	"PdhGetLogFileTypeW",
	"PdhGetRawCounterArrayW",
	"PdhListLogFileHeaderW",
	"PdhLogServiceCommandW",
	"PdhLogServiceControlW",
	"PdhLookupPerfIndexByNameW",
	"PdhLookupPerfNameByIndexW",
	"PdhMakeCounterPathW",
	"PdhOpenLogW",
	"PdhOpenQueryW",
	"PdhParseCounterPathW",
	"PdhParseInstanceNameW",
	"PdhPlaAddItemW",
	"PdhPlaCreateW",
	"PdhPlaDeleteW",
	"PdhPlaEnumCollectionsW",
	"PdhPlaGetInfoW",
	"PdhPlaGetLogFileNameW",
	"PdhPlaGetScheduleW",
	"PdhPlaRemoveAllItemsW",
	"PdhPlaScheduleW",
	"PdhPlaSetInfoW",
	"PdhPlaSetItemListW",
	"PdhPlaSetRunAsW",
	"PdhPlaStartW",
	"PdhPlaStopW",
	"PdhPlaValidateInfoW",
	"PdhRelogW",
	"PdhSelectDataSourceW",
	"PdhTranslate009CounterW",
	"PdhTranslateLocaleCounterW",
	"PdhUpdateLogW",
	"PdhValidatePathW",
	"PdhVerifySQLDBW",
	"PdhiPlaFormatBlanksW",
}
