// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

/*
Register all entry-points in ws2_32.dll.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.

Note that all dll exported named entry-points are listed,
including those that are undocumented by the vendor.
*/
package ws2_32

import "github.com/tHinqa/outside"

func init() {
	outside.AddEPs("ws2_32.dll", false, EntryPoints)
	outside.AddEPs("ws2_32.dll", true, UnicodeEntryPoints)
}

//TODO(t): Check Ws with no A counterparts and vv

var EntryPoints = outside.EPs{
	"WEP",
	"WPUCompleteOverlappedRequest",
	"WSAAccept",
	"WSAAddressToStringA",
	"WSAAsyncGetHostByAddr",
	"WSAAsyncGetHostByName",
	"WSAAsyncGetProtoByName",
	"WSAAsyncGetProtoByNumber",
	"WSAAsyncGetServByName",
	"WSAAsyncGetServByPort",
	"WSAAsyncSelect",
	"WSACancelAsyncRequest",
	"WSACancelBlockingCall",
	"WSACleanup",
	"WSACloseEvent",
	"WSAConnect",
	"WSACreateEvent",
	"WSADuplicateSocketA",
	"WSAEnumNameSpaceProvidersA",
	"WSAEnumNetworkEvents",
	"WSAEnumProtocolsA",
	"WSAEventSelect",
	"WSAGetLastError",
	"WSAGetOverlappedResult",
	"WSAGetQOSByName",
	"WSAGetServiceClassInfoA",
	"WSAGetServiceClassNameByClassIdA",
	"WSAHtonl",
	"WSAHtons",
	"WSAInstallServiceClassA",
	"WSAIoctl",
	"WSAIsBlocking",
	"WSAJoinLeaf",
	"WSALookupServiceBeginA",
	"WSALookupServiceEnd",
	"WSALookupServiceNextA",
	"WSANSPIoctl",
	"WSANtohl",
	"WSANtohs",
	"WSAProviderConfigChange",
	"WSARecv",
	"WSARecvDisconnect",
	"WSARecvFrom",
	"WSARemoveServiceClass",
	"WSAResetEvent",
	"WSASend",
	"WSASendDisconnect",
	"WSASendTo",
	"WSASetBlockingHook",
	"WSASetEvent",
	"WSASetLastError",
	"WSASetServiceA",
	"WSASocketA",
	"WSAStartup",
	"WSAStringToAddressA",
	"WSAUnhookBlockingHook",
	"WSAWaitForMultipleEvents",
	"WSApSetPostRoutine",
	"WSCDeinstallProvider",
	"WSCEnableNSProvider",
	"WSCEnumProtocols",
	"WSCGetProviderPath",
	"WSCInstallNameSpace",
	"WSCInstallProvider",
	"WSCUnInstallNameSpace",
	"WSCUpdateProvider",
	"WSCWriteNameSpaceOrder",
	"WSCWriteProviderOrder",
	"__WSAFDIsSet",
	"accept",
	"bind",
	"closesocket",
	"connect",
	"freeaddrinfo",
	"getaddrinfo",
	"gethostbyaddr",
	"gethostbyname",
	"gethostname",
	"getnameinfo",
	"getpeername",
	"getprotobyname",
	"getprotobynumber",
	"getservbyname",
	"getservbyport",
	"getsockname",
	"getsockopt",
	"htonl",
	"htons",
	"inet_addr",
	"inet_ntoa",
	"ioctlsocket",
	"listen",
	"ntohl",
	"ntohs",
	"recv",
	"recvfrom",
	"select",
	"send",
	"sendto",
	"setsockopt",
	"shutdown",
	"socket",
}

var UnicodeEntryPoints = outside.EPs{
	"FreeAddrInfoW",
	"GetAddrInfoW",
	"GetNameInfoW",
	"WSAAddressToStringW",
	"WSADuplicateSocketW",
	"WSAEnumNameSpaceProvidersW",
	"WSAEnumProtocolsW",
	"WSAGetServiceClassInfoW",
	"WSAGetServiceClassNameByClassIdW",
	"WSAInstallServiceClassW",
	"WSALookupServiceBeginW",
	"WSALookupServiceNextW",
	"WSASetServiceW",
	"WSASocketW",
	"WSAStringToAddressW",
}
