package protocol

// DO NOT EDIT
//
// This file was generated by ./schema.sh


// EncodeLeader encodes a Leader request.
func EncodeLeader(request *Message) {
	request.putUint64(0)

	request.putHeader(RequestLeader)
}

// EncodeClient encodes a Client request.
func EncodeClient(request *Message, id uint64) {
	request.putUint64(id)

	request.putHeader(RequestClient)
}

// EncodeHeartbeat encodes a Heartbeat request.
func EncodeHeartbeat(request *Message, timestamp uint64) {
	request.putUint64(timestamp)

	request.putHeader(RequestHeartbeat)
}

// EncodeOpen encodes a Open request.
func EncodeOpen(request *Message, name string, flags uint64, vfs string) {
	request.putString(name)
	request.putUint64(flags)
	request.putString(vfs)

	request.putHeader(RequestOpen)
}

// EncodePrepare encodes a Prepare request.
func EncodePrepare(request *Message, db uint64, sql string) {
	request.putUint64(db)
	request.putString(sql)

	request.putHeader(RequestPrepare)
}

// EncodeExec encodes a Exec request.
func EncodeExec(request *Message, db uint32, stmt uint32, values NamedValues) {
	request.putUint32(db)
	request.putUint32(stmt)
	request.putNamedValues(values)

	request.putHeader(RequestExec)
}

// EncodeQuery encodes a Query request.
func EncodeQuery(request *Message, db uint32, stmt uint32, values NamedValues) {
	request.putUint32(db)
	request.putUint32(stmt)
	request.putNamedValues(values)

	request.putHeader(RequestQuery)
}

// EncodeFinalize encodes a Finalize request.
func EncodeFinalize(request *Message, db uint32, stmt uint32) {
	request.putUint32(db)
	request.putUint32(stmt)

	request.putHeader(RequestFinalize)
}

// EncodeExecSQL encodes a ExecSQL request.
func EncodeExecSQL(request *Message, db uint64, sql string, values NamedValues) {
	request.putUint64(db)
	request.putString(sql)
	request.putNamedValues(values)

	request.putHeader(RequestExecSQL)
}

// EncodeQuerySQL encodes a QuerySQL request.
func EncodeQuerySQL(request *Message, db uint64, sql string, values NamedValues) {
	request.putUint64(db)
	request.putString(sql)
	request.putNamedValues(values)

	request.putHeader(RequestQuerySQL)
}

// EncodeInterrupt encodes a Interrupt request.
func EncodeInterrupt(request *Message, db uint64) {
	request.putUint64(db)

	request.putHeader(RequestInterrupt)
}

// EncodeJoin encodes a Join request.
func EncodeJoin(request *Message, id uint64, address string) {
	request.putUint64(id)
	request.putString(address)

	request.putHeader(RequestJoin)
}

// EncodePromote encodes a Promote request.
func EncodePromote(request *Message, id uint64) {
	request.putUint64(id)

	request.putHeader(RequestPromote)
}

// EncodeRemove encodes a Remove request.
func EncodeRemove(request *Message, id uint64) {
	request.putUint64(id)

	request.putHeader(RequestRemove)
}

// EncodeDump encodes a Dump request.
func EncodeDump(request *Message, name string) {
	request.putString(name)

	request.putHeader(RequestDump)
}

// EncodeCluster encodes a Cluster request.
func EncodeCluster(request *Message) {
	request.putUint64(0)

	request.putHeader(RequestCluster)
}
