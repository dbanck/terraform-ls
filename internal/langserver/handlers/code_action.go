package handlers

import (
	"context"
	"fmt"

	lsctx "github.com/hashicorp/terraform-ls/internal/context"
	"github.com/hashicorp/terraform-ls/internal/langserver/errors"
	ilsp "github.com/hashicorp/terraform-ls/internal/lsp"
	lsp "github.com/hashicorp/terraform-ls/internal/protocol"
	"github.com/hashicorp/terraform-ls/internal/terraform/module"
)

func (h *logHandler) TextDocumentCodeAction(ctx context.Context, params lsp.CodeActionParams) []lsp.CodeAction {
	ca, err := h.textDocumentCodeAction(ctx, params)
	if err != nil {
		h.logger.Printf("code action failed: %s", err)
	}

	return ca
}

func (h *logHandler) textDocumentCodeAction(ctx context.Context, params lsp.CodeActionParams) ([]lsp.CodeAction, error) {
	var ca []lsp.CodeAction
	h.logger.Printf("CODE ACTIONS HERE")

	// For action definitions, refer to https://code.visualstudio.com/api/references/vscode-api#CodeActionKind
	// We only support format type code actions at the moment, and do not want to format without the client asking for
	// them, so exit early here if nothing is requested.
	// if len(params.Context.Only) == 0 {
	// 	h.logger.Printf("No code action requested, exiting")
	// 	return ca, nil
	// }

	for _, o := range params.Context.Only {
		h.logger.Printf("Code actions requested: %q", o)
	}

	// The Only field of the context specifies which code actions the client wants.
	// If Only is empty, assume that the client wants all of the non-explicit code actions.
	var wantedCodeActions map[lsp.CodeActionKind]bool

	if len(params.Context.Only) == 0 {
		wantedCodeActions = ilsp.SupportedCodeActions // TODO! filter by type
	} else {
		wantedCodeActions = ilsp.SupportedCodeActions.Only(params.Context.Only)
	}

	if len(wantedCodeActions) == 0 {
		return nil, fmt.Errorf("could not find a supported code action to execute for %s, wanted %v",
			params.TextDocument.URI, params.Context.Only)
	}

	h.logger.Printf("Code actions supported: %v", wantedCodeActions)

	fh := ilsp.FileHandlerFromDocumentURI(params.TextDocument.URI)

	fs, err := lsctx.DocumentStorage(ctx)
	if err != nil {
		return ca, err
	}
	file, err := fs.GetDocument(fh)
	if err != nil {
		return ca, err
	}
	original, err := file.Text()
	if err != nil {
		return ca, err
	}

	for action := range wantedCodeActions {
		switch action {
		case lsp.RefactorExtract:
			// TODO figure out if the current selection can be extracted

			// TODO build a workspace edit OR command to do said extraction

		case lsp.QuickFix:
			// TODO! only offer quickfixes with matching diagnostics
			ca = append(ca, lsp.CodeAction{
				Title: "Fix issue",
				Kind:  action,
				// Edit:  lsp.WorkspaceEdit{},
				Command: &lsp.Command{
					Title:   "Fill me",
					Command: "notsupported",
				},
			})
		case ilsp.SourceFormatAllTerraform:
			tfExec, err := module.TerraformExecutorForModule(ctx, fh.Dir())
			if err != nil {
				return ca, errors.EnrichTfExecError(err)
			}

			h.logger.Printf("Formatting document via %q", tfExec.GetExecPath())

			edits, err := formatDocument(ctx, tfExec, original, file)
			if err != nil {
				return ca, err
			}

			ca = append(ca, lsp.CodeAction{
				Title: "Format Document",
				Kind:  action,
				Edit: lsp.WorkspaceEdit{
					Changes: map[string][]lsp.TextEdit{
						string(fh.URI()): edits,
					},
				},
			})
		}
	}

	return ca, nil
}
