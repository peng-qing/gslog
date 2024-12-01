package gslog

const (
	badFieldsKey = "!badFieldsKey"
	unknownFile  = "!unknownFile"
)

type LTextFlag int

const (
	LTextTime            LTextFlag = 1 << iota // 时间标记位
	LTextFile                                  // 文件路径标记位
	LTextFunction                              // 调用函数标记位
	LTextLogLevel                              // 日志级别 首字母大写 Trace/Debug/...
	LTextLogLevelUpCase                        // 日志级别 全大写 TRACE/DEBUG/...
	LTextLogLevelLowCase                       // 日志级别 全小写 debug/info/...

	DefaultLTextFlag = LTextTime | LTextFile | LTextLogLevel
	lCheckLogLevel   = LTextLogLevel | LTextLogLevelLowCase | LTextLogLevelUpCase
	lCheckShortFile  = LTextFile | LTextFunction
)

const (
	DefaultTimeLayout = "2006/01/02 15:04:05.000000"
)

const (
	defaultJsonTimeKey    = "time"
	defaultJsonSourceKey  = "source"
	defaultJsonLevelKey   = "level"
	defaultJsonMessageKey = "message"
	defaultJsonFieldsKey  = "fields"
)

const (
	serializeArrayBegin      = '['
	serializeCommaStep       = ','
	serializeArrayEnd        = ']'
	serializePrefixBegin     = '<'
	serializePrefixEnd       = '>'
	serializeRadixPointSplit = '.'
	serializeSpaceSplit      = ' '
	serializeColonSplit      = ':'
	serializeFieldStep       = '='
	serializeNewLine         = '\n'
	serializeJsonStart       = '{'
	serializeJsonEnd         = '}'
	serializeStringMarks     = '"'
)
