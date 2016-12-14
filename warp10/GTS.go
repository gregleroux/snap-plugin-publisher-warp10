package warp10

//TS	Timestamp of the reading, in microseconds since the Unix Epoch
//NAME	Class name of the reading as a [URL encoded](http://en.wikipedia.org/wiki/Percent-encoding) UTF-8 character string. The encoding of character `{` (Unicode *LEFT CURLY BRACKET*, *0x007B*) is MANDATORY.
//LABELS	Comma separated list of labels, using the syntax `key=value` where both `key` and `value` are URL encoded UTF-8 character strings. If a key or value contains `,` (Unicode COMMA, 0x002C),`}` (Unicode RIGHT CURLY BRACKET, 0x007D) or `=` (Unicode EQUALS SIGN, 0x003D), those characters MUST be encoded.
//VALUE	The value of the reading. It can be of one of four types: `LONG`, `DOUBLE`, `BOOLEAN`, `STRING`

type GTS struct {
	Timestamp int64         `json:"timestamp"`
	Metric    string        `json:"metric"`
	Tags      string        `json:"tags"`
	Value     interface{}   `json:"value"`
}