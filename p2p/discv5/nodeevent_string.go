// Code generated by "stringer -type=nodeEvent"; DO NOT EDIT.

package discv5

import "strconv"

const (
	_nodeEvent_name_0 = "invalidEventpingPacketpongPacketfindnodePacketneighborsPacketfindnodeHashPackettopicRegisterPackettopicQueryPackettopicNodesPacket"
	_nodeEvent_name_1 = "pongTimeoutpingTimeoutneighboursTimeout"
)

var (
	_nodeEvent_index_0 = [...]uint8{0, 12, 22, 32, 46, 61, 79, 98, 114, 130}
	_nodeEvent_index_1 = [...]uint8{0, 11, 22, 39}
)

func (i nodeEvent) String() string { log.DebugLog()
	switch {
	case 0 <= i && i <= 8:
		return _nodeEvent_name_0[_nodeEvent_index_0[i]:_nodeEvent_index_0[i+1]]
	case 265 <= i && i <= 267:
		i -= 265
		return _nodeEvent_name_1[_nodeEvent_index_1[i]:_nodeEvent_index_1[i+1]]
	default:
		return "nodeEvent(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
