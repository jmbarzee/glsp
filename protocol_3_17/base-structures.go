package protocol

import (
	protocol316 "github.com/tliron/glsp/protocol_3_16"
)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#baseTypes

/**
 * The LSP any type
 *
 * @since 3.17.0
 */
type LSPAny = any

/**
 * LSP object definition.
 *
 * @since 3.17.0
 */
type LSPObject = map[string]LSPAny

/**
 * LSP arrays.
 *
 * @since 3.17.0
 */
type LSPArray = []LSPAny

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#errorCodes

const (
	/**
	 * This is the start range of LSP reserved error codes.
	 * It doesn't denote a real error code.
	 *
	 * @since 3.16.0
	 */
	LSPReservedErrorRangeStart = protocol316.Integer(-32899)

	/**
	 * A request failed but it was syntactically correct, e.g the
	 * method name was known and the parameters were valid. The error
	 * message should contain human readable information about why
	 * the request failed.
	 *
	 * @since 3.17.0
	 */
	RequestFailed = protocol316.Integer(-32803)

	/**
	 * The server cancelled the request. This error code should
	 * only be used for requests that explicitly support being
	 * server cancellable.
	 *
	 * @since 3.17.0
	 */
	ServerCancelled = protocol316.Integer(-32802)

	/**
	 * The server detected that the content of a document got
	 * modified outside normal conditions. A server should
	 * NOT send this error code if it detects a content change
	 * in it unprocessed messages. The result even computed
	 * on an older state might still be useful for the client.
	 *
	 * If a client decides that a result is not of any use anymore
	 * the client should cancel the request.
	 */
	ContentModified = protocol316.Integer(-32801)

	/**
	 * The client has canceled a request and a server as detected
	 * the cancel.
	 */
	RequestCancelled = protocol316.Integer(-32800)

	/**
	 * This is the end range of LSP reserved error codes.
	 * It doesn't denote a real error code.
	 *
	 * @since 3.16.0
	 */
	LSPReservedErrorRangeEnd = protocol316.Integer(-32800)
)

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#positionEncodingKind

/**
 * A type indicating how positions are encoded,
 * specifically what column offsets mean.
 *
 * @since 3.17.0
 */
type PositionEncodingKind string

const (
	/**
	 * Character offsets count UTF-8 code units (e.g bytes).
	 */
	PositionEncodingKindUTF8 PositionEncodingKind = "utf-8"

	/**
	 * Character offsets count UTF-16 code units.
	 *
	 * This is the default and must always be supported
	 * by servers
	 */
	PositionEncodingKindUTF16 PositionEncodingKind = "utf-16"

	/**
	 * Character offsets count UTF-32 code units.
	 *
	 * Implementation note: these are the same as Unicode code points,
	 * so this `PositionEncodingKind` may also be used for an
	 * encoding-agnostic representation of character offsets.
	 */
	PositionEncodingKindUTF32 PositionEncodingKind = "utf-32"
)
