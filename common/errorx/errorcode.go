package errorx

const (
	OK                 uint32 = 200
	Canceled           uint32 = 408
	Unknown            uint32 = 500
	InvalidArgument    uint32 = 400
	DeadlineExceeded   uint32 = 504
	NotFound           uint32 = 404
	AlreadyExists      uint32 = 409
	PermissionDenied   uint32 = 403
	ResourceExhausted  uint32 = 429
	FailedPrecondition uint32 = 400
	OutOfRange         uint32 = 400
	Aborted            uint32 = 409
	Unimplemented      uint32 = 501
	Internal           uint32 = 500
	Unavailable        uint32 = 503
	DataLoss           uint32 = 500
	Unauthenticated    uint32 = 401

	ServerCommonError uint32 = 500
	RequestParamError uint32 = 400
)

var codeToErr = map[uint32]uint32{
	0:  OK,
	1:  Canceled,
	2:  Unknown,
	3:  InvalidArgument,
	4:  DeadlineExceeded,
	5:  NotFound,
	6:  AlreadyExists,
	7:  PermissionDenied,
	8:  ResourceExhausted,
	9:  FailedPrecondition,
	10: Aborted,
	11: OutOfRange,
	12: Unimplemented,
	13: Internal,
	14: Unavailable,
	15: DataLoss,
	16: Unauthenticated,
}
