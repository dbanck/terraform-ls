package command

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"

	"github.com/creachadair/jrpc2/code"
	"github.com/hashicorp/hcl/v2"
	lsctx "github.com/hashicorp/terraform-ls/internal/context"
	"github.com/hashicorp/terraform-ls/internal/langserver/cmd"
	"github.com/hashicorp/terraform-ls/internal/langserver/diagnostics"
	ilsp "github.com/hashicorp/terraform-ls/internal/lsp"
	lsp "github.com/hashicorp/terraform-ls/internal/protocol"
	"github.com/hashicorp/terraform-ls/internal/terraform/module"
	"github.com/hashicorp/terraform-ls/internal/uri"
)

type TfLintRange struct {
	Filename string `json:"filename"`
	Start    struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"start"`
	End struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"end"`
}

type TfLintIssue struct {
	Rule struct {
		Name     string `json:"name"`
		Severity string `json:"severity"`
		Link     string `json:"link"`
	} `json:"rule"`
	Message string        `json:"message"`
	Range   *TfLintRange  `json:"range"`
	Callers []interface{} `json:"callers"`
}

type TfLintOutput struct {
	Issues []TfLintIssue `json:"issues"`
	Errors []interface{} `json:"errors"`
}

func LintHandler(ctx context.Context, args cmd.CommandArgs) (interface{}, error) {
	dirUri, ok := args.GetString("uri")
	if !ok || dirUri == "" {
		return nil, fmt.Errorf("%w: expected module uri argument to be set", code.InvalidParams.Err())
	}

	if !uri.IsURIValid(dirUri) {
		return nil, fmt.Errorf("URI %q is not valid", dirUri)
	}

	dh := ilsp.FileHandlerFromDirURI(lsp.DocumentURI(dirUri))

	modMgr, err := lsctx.ModuleManager(ctx)
	if err != nil {
		return nil, err
	}

	mod, err := modMgr.ModuleByPath(dh.Dir())
	if err != nil {
		if module.IsModuleNotFound(err) {
			mod, err = modMgr.AddModule(dh.Dir())
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	notifier, err := lsctx.DiagnosticsNotifier(ctx)
	if err != nil {
		return nil, err
	}
	log.Printf("HELLLOOOOO")

	cmd := exec.Command("/usr/local/bin/tflint", "--format=json", mod.Path)
	var outBuf = bytes.Buffer{}
	cmd.Stdout = &outBuf

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("OUTPUT %#v", outBuf)

	var ret TfLintOutput
	jsonErr := json.Unmarshal(outBuf.Bytes(), &ret)
	if jsonErr != nil {
		return nil, jsonErr
	}
	log.Printf("RET %#v", ret)

	diags := diagnostics.NewDiagnostics()
	lintDiags := HCLDiagsFromTfLint(ret.Issues)
	diags.EmptyRootDiagnostic()
	diags.Append("tflint", lintDiags)
	diags.Append("HCL", mod.ModuleDiagnostics.AsMap())
	diags.Append("HCL", mod.VarsDiagnostics.AutoloadedOnly().AsMap())

	notifier.PublishHCLDiags(ctx, mod.Path, diags)

	return nil, nil
}

func HCLDiagsFromTfLint(jsonDiags []TfLintIssue) map[string]hcl.Diagnostics {
	diagsMap := make(map[string]hcl.Diagnostics)

	for _, d := range jsonDiags {
		file := ""
		if d.Range != nil {
			file = d.Range.Filename
		}

		diags := diagsMap[file]

		var severity hcl.DiagnosticSeverity
		if d.Rule.Severity == "error" {
			severity = hcl.DiagError
		} else if d.Rule.Severity == "warning" {
			severity = hcl.DiagWarning
		}

		diag := &hcl.Diagnostic{
			Severity: severity,
			Summary:  d.Rule.Name,
			Detail:   d.Message,
		}

		if d.Range != nil {
			diag.Subject = &hcl.Range{
				Filename: d.Range.Filename,
				Start: hcl.Pos{
					Line:   d.Range.Start.Line,
					Column: d.Range.Start.Column,
				},
				End: hcl.Pos{
					Line:   d.Range.End.Line,
					Column: d.Range.End.Column,
				},
			}
		}

		diags = append(diags, diag)

		diagsMap[file] = diags
	}

	return diagsMap
}