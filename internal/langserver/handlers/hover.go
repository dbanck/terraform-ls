package handlers

import (
	"context"

	"github.com/hashicorp/hcl/v2"
	lsctx "github.com/hashicorp/terraform-ls/internal/context"
	ilsp "github.com/hashicorp/terraform-ls/internal/lsp"
	lsp "github.com/hashicorp/terraform-ls/internal/protocol"
)

type Diagnostic struct {
	Dir      string
	Filename string
	Range    hcl.Range
	Severity hcl.DiagnosticSeverity

	// The diagnostic's code, which might appear in the user interface
	// e.g. "terraform_deprecated_interpolation"
	Code string
	// An optional property to describe the error code
	CodeDescription string

	// A human-readable string describing the source, e.g. "tflint"
	Source string
	// The diagnostic's message.
	// e.g. "Interpolation-only expressions are deprecated"
	Message string

	// Optional data entry field
	Data interface{}
}

func (svc *service) TextDocumentHover(ctx context.Context, params lsp.TextDocumentPositionParams) (*lsp.Hover, error) {
	fs, err := lsctx.DocumentStorage(ctx)
	if err != nil {
		return nil, err
	}

	cc, err := ilsp.ClientCapabilities(ctx)
	if err != nil {
		return nil, err
	}

	doc, err := fs.GetDocument(ilsp.FileHandlerFromDocumentURI(params.TextDocument.URI))
	if err != nil {
		return nil, err
	}

	d, err := svc.decoderForDocument(ctx, doc)
	if err != nil {
		return nil, err
	}

	fPos, err := ilsp.FilePositionFromDocumentPosition(params, doc)
	if err != nil {
		return nil, err
	}

	svc.logger.Printf("Looking for hover data at %q -> %#v", doc.Filename(), fPos.Position())
	hoverData, err := d.HoverAtPos(doc.Filename(), fPos.Position())
	svc.logger.Printf("received hover data: %#v", hoverData)
	if err != nil {
		return nil, err
	}

	return ilsp.HoverData(hoverData, cc.TextDocument), nil
}
