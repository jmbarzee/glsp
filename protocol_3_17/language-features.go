package protocol

import (
	"github.com/tliron/glsp"
	protocol316 "github.com/tliron/glsp/protocol_3_16"
)

// ========================================================================================
// Type Hierarchy
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#typeHierarchy
// ========================================================================================

/**
 * @since 3.17.0
 */
type TypeHierarchyClientCapabilities struct {
	/**
	 * Whether implementation supports dynamic registration. If this is set to
	 * `true` the client supports the new `(TextDocumentRegistrationOptions &
	 * StaticRegistrationOptions)` return value for the corresponding server
	 * capability as well.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

/**
 * Type hierarchy options used during static registration.
 *
 * @since 3.17.0
 */
type TypeHierarchyOptions struct {
	protocol316.WorkDoneProgressOptions
}

/**
 * Type hierarchy options used during static or dynamic registration.
 *
 * @since 3.17.0
 */
type TypeHierarchyRegistrationOptions struct {
	protocol316.TextDocumentRegistrationOptions
	TypeHierarchyOptions
	protocol316.StaticRegistrationOptions
}

const MethodTextDocumentPrepareTypeHierarchy = protocol316.Method("textDocument/prepareTypeHierarchy")

/**
 * The parameters of a `textDocument/prepareTypeHierarchy` request.
 *
 * @since 3.17.0
 */
type TypeHierarchyPrepareParams struct {
	protocol316.TextDocumentPositionParams
	protocol316.WorkDoneProgressParams
}

type TextDocumentPrepareTypeHierarchyFunc func(context *glsp.Context, params *TypeHierarchyPrepareParams) ([]TypeHierarchyItem, error)

/**
 * @since 3.17.0
 */
type TypeHierarchyItem struct {
	/**
	 * The name of this item.
	 */
	Name string `json:"name"`

	/**
	 * The kind of this item.
	 */
	Kind protocol316.SymbolKind `json:"kind"`

	/**
	 * Tags for this item.
	 */
	Tags []protocol316.SymbolTag `json:"tags,omitempty"`

	/**
	 * More detail for this item, e.g. the signature of a function.
	 */
	Detail *string `json:"detail,omitempty"`

	/**
	 * The resource identifier of this item.
	 */
	URI protocol316.DocumentUri `json:"uri"`

	/**
	 * The range enclosing this symbol not including leading/trailing whitespace
	 * but everything else, e.g. comments and code.
	 */
	Range protocol316.Range `json:"range"`

	/**
	 * The range that should be selected and revealed when this symbol is being
	 * picked, e.g. the name of a function. Must be contained by the
	 * [`range`](#TypeHierarchyItem.range).
	 */
	SelectionRange protocol316.Range `json:"selectionRange"`

	/**
	 * A data entry field that is preserved between a type hierarchy prepare and
	 * supertypes or subtypes requests. It could also be used to identify the
	 * type hierarchy in the server, helping improve the performance on
	 * resolving supertypes and subtypes.
	 */
	Data LSPAny `json:"data,omitempty"`
}

const MethodTypeHierarchySupertypes = protocol316.Method("typeHierarchy/supertypes")

/**
 * The parameters of a `typeHierarchy/supertypes` request.
 *
 * @since 3.17.0
 */
type TypeHierarchySupertypesParams struct {
	protocol316.WorkDoneProgressParams
	protocol316.PartialResultParams

	Item TypeHierarchyItem `json:"item"`
}

type TypeHierarchySupertypesFunc func(context *glsp.Context, params *TypeHierarchySupertypesParams) ([]TypeHierarchyItem, error)

const MethodTypeHierarchySubtypes = protocol316.Method("typeHierarchy/subtypes")

/**
 * The parameters of a `typeHierarchy/subtypes` request.
 *
 * @since 3.17.0
 */
type TypeHierarchySubtypesParams struct {
	protocol316.WorkDoneProgressParams
	protocol316.PartialResultParams

	Item TypeHierarchyItem `json:"item"`
}

type TypeHierarchySubtypesFunc func(context *glsp.Context, params *TypeHierarchySubtypesParams) ([]TypeHierarchyItem, error)

// ========================================================================================
// Inlay Hint
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#inlayHint
// ========================================================================================

/**
 * Inlay hint client capabilities.
 *
 * @since 3.17.0
 */
type InlayHintClientCapabilities struct {
	/**
	 * Whether inlay hints support dynamic registration.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`

	/**
	 * Indicates which properties a client can resolve lazily on an inlay
	 * hint.
	 */
	ResolveSupport *struct {
		/**
		 * The properties that a client can resolve lazily.
		 */
		Properties []string `json:"properties"`
	} `json:"resolveSupport,omitempty"`
}

/**
 * Client workspace capabilities specific to inlay hints.
 *
 * @since 3.17.0
 */
type InlayHintWorkspaceClientCapabilities struct {
	/**
	 * Whether the client implementation supports a refresh request sent from
	 * the server to the client.
	 *
	 * Note that this event is global and will force the client to refresh all
	 * inlay hints currently shown. It should be used with absolute care and
	 * is useful for situation where a server for example detect a project wide
	 * change that requires such a calculation.
	 */
	RefreshSupport *bool `json:"refreshSupport,omitempty"`
}

/**
 * Inlay hint options used during static registration.
 *
 * @since 3.17.0
 */
type InlayHintOptions struct {
	protocol316.WorkDoneProgressOptions

	/**
	 * The server provides support to resolve additional
	 * information for an inlay hint item.
	 */
	ResolveProvider *bool `json:"resolveProvider,omitempty"`
}

/**
 * Inlay hint options used during static or dynamic registration.
 *
 * @since 3.17.0
 */
type InlayHintRegistrationOptions struct {
	InlayHintOptions
	protocol316.TextDocumentRegistrationOptions
	protocol316.StaticRegistrationOptions
}

const MethodTextDocumentInlayHint = protocol316.Method("textDocument/inlayHint")

/**
 * A parameter literal used in inlay hint requests.
 *
 * @since 3.17.0
 */
type InlayHintParams struct {
	protocol316.WorkDoneProgressParams

	/**
	 * The text document.
	 */
	TextDocument protocol316.TextDocumentIdentifier `json:"textDocument"`

	/**
	 * The visible document range for which inlay hints should be computed.
	 */
	Range protocol316.Range `json:"range"`
}

type TextDocumentInlayHintFunc func(context *glsp.Context, params *InlayHintParams) ([]InlayHint, error)

/**
 * Inlay hint information.
 *
 * @since 3.17.0
 */
type InlayHint struct {
	/**
	 * The position of this hint.
	 */
	Position protocol316.Position `json:"position"`

	/**
	 * The label of this hint. A human readable string or an array of
	 * InlayHintLabelPart label parts.
	 *
	 * *Note* that neither the string nor the label part can be empty.
	 */
	Label any `json:"label"` // string | []InlayHintLabelPart

	/**
	 * The kind of this hint. Can be omitted in which case the client
	 * should fall back to a reasonable default.
	 */
	Kind *InlayHintKind `json:"kind,omitempty"`

	/**
	 * Optional text edits that are performed when accepting this inlay hint.
	 *
	 * *Note* that edits are expected to change the document so that the inlay
	 * hint (or its nearest variant) is now part of the document and the inlay
	 * hint itself is now obsolete.
	 */
	TextEdits []protocol316.TextEdit `json:"textEdits,omitempty"`

	/**
	 * The tooltip text when you hover over this item.
	 */
	Tooltip any `json:"tooltip,omitempty"` // string | MarkupContent

	/**
	 * Render padding before the hint.
	 *
	 * Note: Padding should use the editor's background color, not the
	 * background color of the hint itself. That means padding can be used
	 * to visually align/separate an inlay hint.
	 */
	PaddingLeft *bool `json:"paddingLeft,omitempty"`

	/**
	 * Render padding after the hint.
	 *
	 * Note: Padding should use the editor's background color, not the
	 * background color of the hint itself. That means padding can be used
	 * to visually align/separate an inlay hint.
	 */
	PaddingRight *bool `json:"paddingRight,omitempty"`

	/**
	 * A data entry field that is preserved on an inlay hint between
	 * a `textDocument/inlayHint` and a `inlayHint/resolve` request.
	 */
	Data LSPAny `json:"data,omitempty"`
}

/**
 * An inlay hint label part allows for interactive and composite labels
 * of inlay hints.
 *
 * @since 3.17.0
 */
type InlayHintLabelPart struct {
	/**
	 * The value of this label part.
	 */
	Value string `json:"value"`

	/**
	 * The tooltip text when you hover over this label part. Depending on
	 * the client capability `inlayHint.resolveSupport` clients might resolve
	 * this property late using the resolve request.
	 */
	Tooltip any `json:"tooltip,omitempty"` // string | MarkupContent

	/**
	 * An optional source code location that represents this label part.
	 *
	 * The editor will use this location for the hover and for code navigation
	 * features: This part will become a clickable link that resolves to the
	 * definition of the symbol at the given location (not necessarily the
	 * location itself), it shows the hover that shows at the given location,
	 * and it shows a context menu with further code navigation commands.
	 *
	 * Depending on the client capability `inlayHint.resolveSupport` clients
	 * might resolve this property late using the resolve request.
	 */
	Location *protocol316.Location `json:"location,omitempty"`

	/**
	 * An optional command for this label part.
	 *
	 * Depending on the client capability `inlayHint.resolveSupport` clients
	 * might resolve this property late using the resolve request.
	 */
	Command *protocol316.Command `json:"command,omitempty"`
}

/**
 * Inlay hint kinds.
 *
 * @since 3.17.0
 */
type InlayHintKind protocol316.Integer

const (
	/**
	 * An inlay hint that for a type annotation.
	 */
	InlayHintKindType InlayHintKind = 1

	/**
	 * An inlay hint that is for a parameter.
	 */
	InlayHintKindParameter InlayHintKind = 2
)

const MethodInlayHintResolve = protocol316.Method("inlayHint/resolve")

type InlayHintResolveFunc func(context *glsp.Context, params *InlayHint) (*InlayHint, error)

const MethodWorkspaceInlayHintRefresh = protocol316.Method("workspace/inlayHint/refresh")

type WorkspaceInlayHintRefreshFunc func(context *glsp.Context) error

// ========================================================================================
// Inline Value
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#inlineValue
// ========================================================================================

/**
 * Client capabilities specific to inline values.
 *
 * @since 3.17.0
 */
type InlineValueClientCapabilities struct {
	/**
	 * Whether implementation supports dynamic registration for inline
	 * value providers.
	 */
	DynamicRegistration *bool `json:"dynamicRegistration,omitempty"`
}

/**
 * Client workspace capabilities specific to inline values.
 *
 * @since 3.17.0
 */
type InlineValueWorkspaceClientCapabilities struct {
	/**
	 * Whether the client implementation supports a refresh request sent from
	 * the server to the client.
	 *
	 * Note that this event is global and will force the client to refresh all
	 * inline values currently shown. It should be used with absolute care and
	 * is useful for situation where a server for example detect a project wide
	 * change that requires such a calculation.
	 */
	RefreshSupport *bool `json:"refreshSupport,omitempty"`
}

/**
 * Inline value options used during static registration.
 *
 * @since 3.17.0
 */
type InlineValueOptions struct {
	protocol316.WorkDoneProgressOptions
}

/**
 * Inline value options used during static or dynamic registration.
 *
 * @since 3.17.0
 */
type InlineValueRegistrationOptions struct {
	InlineValueOptions
	protocol316.TextDocumentRegistrationOptions
	protocol316.StaticRegistrationOptions
}

const MethodTextDocumentInlineValue = protocol316.Method("textDocument/inlineValue")

/**
 * A parameter literal used in inline value requests.
 *
 * @since 3.17.0
 */
type InlineValueParams struct {
	protocol316.WorkDoneProgressParams

	/**
	 * The text document.
	 */
	TextDocument protocol316.TextDocumentIdentifier `json:"textDocument"`

	/**
	 * The document range for which inline values should be computed.
	 */
	Range protocol316.Range `json:"range"`

	/**
	 * Additional information about the context in which inline values were
	 * requested.
	 */
	Context InlineValueContext `json:"context"`
}

type TextDocumentInlineValueFunc func(context *glsp.Context, params *InlineValueParams) ([]InlineValue, error)

/**
 * @since 3.17.0
 */
type InlineValueContext struct {
	/**
	 * The stack frame (as a DAP Id) where the execution has stopped.
	 */
	FrameID protocol316.Integer `json:"frameId"`

	/**
	 * The document range where execution has stopped.
	 * Typically the end position of the range denotes the line where the
	 * inline values are shown.
	 */
	StoppedLocation protocol316.Range `json:"stoppedLocation"`
}

/**
 * Inline value information can be provided by different means:
 * - directly as a text value (class InlineValueText).
 * - as a name to use for a variable lookup (class InlineValueVariableLookup)
 * - as an evaluatable expression (class InlineValueEvaluatableExpression)
 * The InlineValue types combines all inline value types into one type.
 *
 * @since 3.17.0
 */
type InlineValue any // InlineValueText | InlineValueVariableLookup | InlineValueEvaluatableExpression

/**
 * Provide inline value as text.
 *
 * @since 3.17.0
 */
type InlineValueText struct {
	/**
	 * The document range for which the inline value applies.
	 */
	Range protocol316.Range `json:"range"`

	/**
	 * The text of the inline value.
	 */
	Text string `json:"text"`
}

/**
 * Provide inline value through a variable lookup.
 *
 * If only a range is specified, the variable name will be extracted from
 * the underlying document.
 *
 * An optional variable name can be used to override the extracted name.
 *
 * @since 3.17.0
 */
type InlineValueVariableLookup struct {
	/**
	 * The document range for which the inline value applies.
	 * The range is used to extract the variable name from the underlying
	 * document.
	 */
	Range protocol316.Range `json:"range"`

	/**
	 * If specified the name of the variable to look up.
	 */
	VariableName *string `json:"variableName,omitempty"`

	/**
	 * How to perform the lookup.
	 */
	CaseSensitiveLookup bool `json:"caseSensitiveLookup"`
}

/**
 * Provide an inline value through an expression evaluation.
 *
 * If only a range is specified, the expression will be extracted from the
 * underlying document.
 *
 * An optional expression can be used to override the extracted expression.
 *
 * @since 3.17.0
 */
type InlineValueEvaluatableExpression struct {
	/**
	 * The document range for which the inline value applies.
	 * The range is used to extract the evaluatable expression from the
	 * underlying document.
	 */
	Range protocol316.Range `json:"range"`

	/**
	 * If specified the expression overrides the extracted expression.
	 */
	Expression *string `json:"expression,omitempty"`
}

const MethodWorkspaceInlineValueRefresh = protocol316.Method("workspace/inlineValue/refresh")

type WorkspaceInlineValueRefreshFunc func(context *glsp.Context) error
