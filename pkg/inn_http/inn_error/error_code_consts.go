package inn_error

// User errors
const (
	NO_ERROR                                           = 0
	BAD_SYNTAX                                         = 1010
	OPERATION_NOT_ALLOWED                              = 1020
	INVALID_SENDER                                     = 1021
	INVALID_RECEIVER                                   = 1022
	REJECTED_BY_ORACLE                                 = 1030
	INVALID_SIGNATURE                                  = 1040
	INVALID_SIGNATURE_WRONG_SIZE                       = 1041
	INVALID_SIGNATURE_FAILED_SIGN                      = 1042
	URI_ERROR                                          = 1050
	URI_ERROR_UNEXPECTED_URI                           = 1051
	URI_ERROR_MALFORMED_URI                            = 1052
	SHARD_NAME_NOT_AVAILABLE                           = 1060
	SHARD_INDEX_NOT_AVAILABLE                          = 1061
	INTEGRITY_CHECK_FAILED                             = 1070
	SHOPPING_CART_ERROR                                = 1080
	ITEM_IS_NO_LONGER_AVAILABLE                        = 1081
	ITEM_IS_ALREADY_IN_YOUR_CART                       = 1082
	INVALID_COIN_ID                                    = 1100
	COIN_ID_UNDERFLOW                                  = 1101
	COIN_ID_OVERFLOW                                   = 1102
	ASSET_NOT_FOUND                                    = 1103
	ASSET_MISSING_REQUIRED_PROPERTY                    = 1104
	INVALID_CLIENT_ID                                  = 1110
	DUPLICATED_CLIENT_ID                               = 1111
	CLIENT_ID_TOO_SHORT                                = 1112
	INVALID_AMOUNT                                     = 1120
	INSUFFICIENT_FUNDS                                 = 1121
	INVALID_COMMENT                                    = 1130
	INVALID_AUTHORIZATION                              = 1140
	INSUFFICIENT_ASSET_PERMISSIONS                     = 1141
	INVALID_PASSWORD                                   = 1142
	INVALID_EMAIL_ADDRESS                              = 1143
	SECOND_FACTOR_REQUIRED                             = 1144
	SECOND_FACTOR_INCORRECT                            = 1145
	IDENTITY_ERROR                                     = 1150
	INVALID_PROPERTIES_FOR_IDENTITY                    = 1151
	DUPLICATE_IDENTITY                                 = 1152
	INSUFFICIENT_PERMISSION_TO_VIEW_IDENTITY           = 1153
	INVALID_TEMPLATE_NAME                              = 1160
	INVALID_CLIENT_API_KEY                             = 1170
	UNSUPPORTED_CLIENT                                 = 1171
	INVALID_TIME                                       = 1180
	TRANSACTION_IS_IN_THE_FUTURE                       = 1181
	TRANSACTION_EXPIRED                                = 1182
	TRANSACTION_ACCEPTED                               = 1200
	TRANSACTION_REJECTED_INSUFFICIENT_FUNDS            = 1210
	TRANSACTION_REJECTED_TRANSFER_AMOUNTS_NOT_BALANCED = 1211
)

// Network Errors
const (
	CONNECTION_REFUSED  = 2010
	OPERATION_TIMED_OUT = 2020
)

// IO Errors
const (
	DATABASE_LOGIN_FAILED = 3010
	USER_NOT_FOUND        = 3020
	SENDING_EMAIL_FAILED  = 3030
)

// Session Errors
const (
	SESSION_EXPIRED    = 4010
	UNAUTHORIZED       = 4020
	EMAIL_VERIFICATION = 4030
)

// Internal Errors
const (
	INTERNAL_CONFIGURATION_ERROR = 5001
)
