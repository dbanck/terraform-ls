package lsp

// Code generated by gen/unmarshalers.go. DO NOT EDIT.

import (
	"encoding/json"
)

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CancelParams) UnmarshalJSON(b []byte) error {
	type t CancelParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ClientCapabilities) UnmarshalJSON(b []byte) error {
	type t ClientCapabilities
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ClientInfo) UnmarshalJSON(b []byte) error {
	type t ClientInfo
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CodeActionContext) UnmarshalJSON(b []byte) error {
	type t CodeActionContext
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CodeActionKind) UnmarshalJSON(b []byte) error {
	type t CodeActionKind
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CodeActionParams) UnmarshalJSON(b []byte) error {
	type t CodeActionParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CodeLens) UnmarshalJSON(b []byte) error {
	type t CodeLens
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CodeLensOptions) UnmarshalJSON(b []byte) error {
	type t CodeLensOptions
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CodeLensParams) UnmarshalJSON(b []byte) error {
	type t CodeLensParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *Command) UnmarshalJSON(b []byte) error {
	type t Command
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CompletionContext) UnmarshalJSON(b []byte) error {
	type t CompletionContext
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CompletionItem) UnmarshalJSON(b []byte) error {
	type t CompletionItem
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CompletionItemKind) UnmarshalJSON(b []byte) error {
	type t CompletionItemKind
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CompletionList) UnmarshalJSON(b []byte) error {
	type t CompletionList
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CompletionOptions) UnmarshalJSON(b []byte) error {
	type t CompletionOptions
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CompletionParams) UnmarshalJSON(b []byte) error {
	type t CompletionParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *CompletionTriggerKind) UnmarshalJSON(b []byte) error {
	type t CompletionTriggerKind
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ConfigurationItem) UnmarshalJSON(b []byte) error {
	type t ConfigurationItem
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ConfigurationParams) UnmarshalJSON(b []byte) error {
	type t ConfigurationParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ConfigurationResult) UnmarshalJSON(b []byte) error {
	type t ConfigurationResult
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *Diagnostic) UnmarshalJSON(b []byte) error {
	type t Diagnostic
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DiagnosticSeverity) UnmarshalJSON(b []byte) error {
	type t DiagnosticSeverity
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DidChangeConfigurationParams) UnmarshalJSON(b []byte) error {
	type t DidChangeConfigurationParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DidChangeTextDocumentParams) UnmarshalJSON(b []byte) error {
	type t DidChangeTextDocumentParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DidChangeWatchedFilesParams) UnmarshalJSON(b []byte) error {
	type t DidChangeWatchedFilesParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DidCloseTextDocumentParams) UnmarshalJSON(b []byte) error {
	type t DidCloseTextDocumentParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DidOpenTextDocumentParams) UnmarshalJSON(b []byte) error {
	type t DidOpenTextDocumentParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DidSaveTextDocumentParams) UnmarshalJSON(b []byte) error {
	type t DidSaveTextDocumentParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DocumentFormattingParams) UnmarshalJSON(b []byte) error {
	type t DocumentFormattingParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DocumentHighlight) UnmarshalJSON(b []byte) error {
	type t DocumentHighlight
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DocumentHighlightKind) UnmarshalJSON(b []byte) error {
	type t DocumentHighlightKind
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DocumentOnTypeFormattingOptions) UnmarshalJSON(b []byte) error {
	type t DocumentOnTypeFormattingOptions
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DocumentOnTypeFormattingParams) UnmarshalJSON(b []byte) error {
	type t DocumentOnTypeFormattingParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DocumentRangeFormattingParams) UnmarshalJSON(b []byte) error {
	type t DocumentRangeFormattingParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DocumentSymbolParams) UnmarshalJSON(b []byte) error {
	type t DocumentSymbolParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DocumentURI) UnmarshalJSON(b []byte) error {
	type t DocumentURI
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *DocumentationFormat) UnmarshalJSON(b []byte) error {
	type t DocumentationFormat
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ExecuteCommandOptions) UnmarshalJSON(b []byte) error {
	type t ExecuteCommandOptions
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ExecuteCommandParams) UnmarshalJSON(b []byte) error {
	type t ExecuteCommandParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *FileChangeType) UnmarshalJSON(b []byte) error {
	type t FileChangeType
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *FileEvent) UnmarshalJSON(b []byte) error {
	type t FileEvent
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *FormattingOptions) UnmarshalJSON(b []byte) error {
	type t FormattingOptions
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *Hover) UnmarshalJSON(b []byte) error {
	type t Hover
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *InitializeError) UnmarshalJSON(b []byte) error {
	type t InitializeError
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *InitializeParams) UnmarshalJSON(b []byte) error {
	type t InitializeParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *InitializeResult) UnmarshalJSON(b []byte) error {
	type t InitializeResult
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *InsertTextFormat) UnmarshalJSON(b []byte) error {
	type t InsertTextFormat
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *Location) UnmarshalJSON(b []byte) error {
	type t Location
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *LogMessageParams) UnmarshalJSON(b []byte) error {
	type t LogMessageParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *MessageActionItem) UnmarshalJSON(b []byte) error {
	type t MessageActionItem
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *MessageType) UnmarshalJSON(b []byte) error {
	type t MessageType
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *None) UnmarshalJSON(b []byte) error {
	type t None
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ParameterInformation) UnmarshalJSON(b []byte) error {
	type t ParameterInformation
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *Position) UnmarshalJSON(b []byte) error {
	type t Position
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *PublishDiagnosticsParams) UnmarshalJSON(b []byte) error {
	type t PublishDiagnosticsParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *Range) UnmarshalJSON(b []byte) error {
	type t Range
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ReferenceContext) UnmarshalJSON(b []byte) error {
	type t ReferenceContext
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ReferenceParams) UnmarshalJSON(b []byte) error {
	type t ReferenceParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *RenameParams) UnmarshalJSON(b []byte) error {
	type t RenameParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ResourceOperation) UnmarshalJSON(b []byte) error {
	type t ResourceOperation
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *SaveOptions) UnmarshalJSON(b []byte) error {
	type t SaveOptions
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *SemanticHighlightingOptions) UnmarshalJSON(b []byte) error {
	type t SemanticHighlightingOptions
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *SemanticHighlightingParams) UnmarshalJSON(b []byte) error {
	type t SemanticHighlightingParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *SemanticHighlightingToken) UnmarshalJSON(b []byte) error {
	type t SemanticHighlightingToken
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *SemanticHighlightingTokens) UnmarshalJSON(b []byte) error {
	type t SemanticHighlightingTokens
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ServerCapabilities) UnmarshalJSON(b []byte) error {
	type t ServerCapabilities
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ShowMessageParams) UnmarshalJSON(b []byte) error {
	type t ShowMessageParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *ShowMessageRequestParams) UnmarshalJSON(b []byte) error {
	type t ShowMessageRequestParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *SignatureHelp) UnmarshalJSON(b []byte) error {
	type t SignatureHelp
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *SignatureHelpOptions) UnmarshalJSON(b []byte) error {
	type t SignatureHelpOptions
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *SignatureInformation) UnmarshalJSON(b []byte) error {
	type t SignatureInformation
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *SymbolInformation) UnmarshalJSON(b []byte) error {
	type t SymbolInformation
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *SymbolKind) UnmarshalJSON(b []byte) error {
	type t SymbolKind
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *TextDocumentClientCapabilities) UnmarshalJSON(b []byte) error {
	type t TextDocumentClientCapabilities
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *TextDocumentContentChangeEvent) UnmarshalJSON(b []byte) error {
	type t TextDocumentContentChangeEvent
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *TextDocumentIdentifier) UnmarshalJSON(b []byte) error {
	type t TextDocumentIdentifier
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *TextDocumentItem) UnmarshalJSON(b []byte) error {
	type t TextDocumentItem
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *TextDocumentPositionParams) UnmarshalJSON(b []byte) error {
	type t TextDocumentPositionParams
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *TextDocumentSyncKind) UnmarshalJSON(b []byte) error {
	type t TextDocumentSyncKind
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *TextDocumentSyncOptions) UnmarshalJSON(b []byte) error {
	type t TextDocumentSyncOptions
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *TextEdit) UnmarshalJSON(b []byte) error {
	type t TextEdit
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *Trace) UnmarshalJSON(b []byte) error {
	type t Trace
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *VersionedTextDocumentIdentifier) UnmarshalJSON(b []byte) error {
	type t VersionedTextDocumentIdentifier
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *WindowClientCapabilities) UnmarshalJSON(b []byte) error {
	type t WindowClientCapabilities
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *WorkspaceClientCapabilities) UnmarshalJSON(b []byte) error {
	type t WorkspaceClientCapabilities
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *WorkspaceEdit) UnmarshalJSON(b []byte) error {
	type t WorkspaceEdit
	return json.Unmarshal(b, (*t)(v))
}

// UnmarshalJSON implements non-strict json.Unmarshaler.
func (v *WorkspaceSymbolParams) UnmarshalJSON(b []byte) error {
	type t WorkspaceSymbolParams
	return json.Unmarshal(b, (*t)(v))
}