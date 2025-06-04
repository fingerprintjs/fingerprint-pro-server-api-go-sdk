// Deprecated Aliases for Backward Compatibility
//
// This file contains deprecated constant aliases to provide backward compatibility
// with previous versions of the SDK. These aliases allow existing codebases to
// migrate gradually to the new naming convention without breaking changes.
//
// IMPORTANT: This file is intended as a temporary solution and should be **removed**
// in the next major release.

package sdk

// BotdBotResult

// Deprecated: Use BotdBotResult_NOT_DETECTED instead.
const NOT_DETECTED BotdBotResult = BotdBotResult_NOT_DETECTED

// Deprecated: Use BotdBotResult_GOOD instead.
const GOOD BotdBotResult = BotdBotResult_GOOD

// Deprecated: Use BotdBotResult_BAD instead.
const BAD BotdBotResult = BotdBotResult_BAD

// ErrorCode

// Deprecated: Use ErrorCode_REQUEST_CANNOT_BE_PARSED instead.
const REQUEST_CANNOT_BE_PARSED ErrorCode = ErrorCode_REQUEST_CANNOT_BE_PARSED

// Deprecated: Use ErrorCode_TOKEN_REQUIRED instead.
const TOKEN_REQUIRED ErrorCode = ErrorCode_TOKEN_REQUIRED

// Deprecated: Use ErrorCode_TOKEN_NOT_FOUND instead.
const TOKEN_NOT_FOUND ErrorCode = ErrorCode_TOKEN_NOT_FOUND

// Deprecated: Use ErrorCode_SUBSCRIPTION_NOT_ACTIVE instead.
const SUBSCRIPTION_NOT_ACTIVE ErrorCode = ErrorCode_SUBSCRIPTION_NOT_ACTIVE

// Deprecated: Use ErrorCode_WRONG_REGION instead.
const WRONG_REGION ErrorCode = ErrorCode_WRONG_REGION

// Deprecated: Use ErrorCode_FEATURE_NOT_ENABLED instead.
const FEATURE_NOT_ENABLED ErrorCode = ErrorCode_FEATURE_NOT_ENABLED

// Deprecated: Use ErrorCode_REQUEST_NOT_FOUND instead.
const REQUEST_NOT_FOUND ErrorCode = ErrorCode_REQUEST_NOT_FOUND

// Deprecated: Use ErrorCode_VISITOR_NOT_FOUND instead.
const VISITOR_NOT_FOUND ErrorCode = ErrorCode_VISITOR_NOT_FOUND

// Deprecated: Use ErrorCode_TOO_MANY_REQUESTS instead.
const TOO_MANY_REQUESTS ErrorCode = ErrorCode_TOO_MANY_REQUESTS

// Deprecated: Use ErrorCode_TOOMANYREQUESTS429 instead.
const TOOMANYREQUESTS429 ErrorCode = ErrorCode_TOOMANYREQUESTS429

// Deprecated: Use ErrorCode_STATE_NOT_READY instead.
const STATE_NOT_READY ErrorCode = ErrorCode_STATE_NOT_READY

// Deprecated: Use ErrorCode_FAILED instead.
const FAILED ErrorCode = ErrorCode_FAILED

// VpnConfidence

// Deprecated: Use VpnConfidence_LOW instead.
const LOW VpnConfidence = VpnConfidence_LOW

// Deprecated: Use VpnConfidence_MEDIUM instead.
const MEDIUM VpnConfidence = VpnConfidence_MEDIUM

// Deprecated: Use VpnConfidence_HIGH instead.
const HIGH VpnConfidence = VpnConfidence_HIGH
