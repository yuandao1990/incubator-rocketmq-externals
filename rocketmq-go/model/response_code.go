package model

const (
	// success
	Success = 0
	// happened unknow error exception
	SystemError = 1
	// system busy
	SystemBusy = 2
	// unsupport request code
	RequestCodeNotSupported = 3
	// transaction failed, because of add db failed
	TransactionFailed = 4
	// Broker flush disk timeout
	FlushDiskTimeout = 10
	// Broker slave unavailable, just for sync double write
	SlaveNotAvailable = 11
	// Broker write slave timeout, just for sync double write
	FlushSlaveTimeout = 12
	// Broker illegal message
	MessageIllegal = 13
	// Broker, Namesrv not available，maybe service is closing or incorrect permission
	ServiceNotAvailable = 14
	// Broker, Namesrv unsupport version
	VersionNOtSupported = 15
	// Broker, Namesrv no permission for operation with send/receive or other
	NoPermission = 16
	// Broker, topic not exist
	TopicNotExist = 17
	// Broker, topic already exist
	TopicExistAlready = 18
	// Broker message not found when pull
	PullNotFound = 19
	// Broker retry immediately, maybe msg was filtered or incorrect notification TODO confirm annotation
	PullRetryImmediately = 20
	// Broker pull offset moved, because of too big or to small TODO confirm annotation
	PullOffsetMoved = 21
	// Broker query not found
	QueryNotFound = 22
	// Broker parse subscription failed
	SubscriptionParseFailed = 23
	// Broker subscription relationship not existed
	SubscriptionNotExist = 24
	// Broker subscription relationship not latest
	SubscriptionNotLatest = 25
	// Broker subscription group not exist
	SubscriptionGroupNotExist = 26
	// Producer transaction should commit
	TransactionShouldCommit = 200
	// Producer transaction should rollback
	TransactionShouldRollback = 201
	// Producer transaction status unknow
	TransactionStatusUnknow = 202
	// Producer ProducerGroup transaction error
	TransactionStatusGroupWrong = 203
	// unit message，need set buyerId
	NoBuyerID = 204

	// unit message，not current unit msg
	NotInCurrentUnit = 205

	// Consumer not online
	ConsumerNotOnline = 206

	// Consumer consume msg timeout
	ConsumeMsgTimeout = 207
)
