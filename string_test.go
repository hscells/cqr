package cqr_test

import (
	"github.com/hscells/cqr"
	"github.com/hscells/groove/combinator"
	"github.com/hscells/transmute/backend"
	"github.com/hscells/transmute/lexer"
	"github.com/hscells/transmute/parser"
	"github.com/hscells/transmute/pipeline"
	"testing"
)

func TestName(t *testing.T) {
	cqrPipeline := pipeline.NewPipeline(
		parser.NewMedlineParser(),
		backend.NewCQRBackend(),
		pipeline.TransmutePipelineOptions{
			LexOptions: lexer.LexOptions{
				FormatParenthesis: false,
			},
			RequiresLexing: true,
		})

	rawQuery := `1. MMSE*.tw.
2. sMMSE.tw.
3. Folstein*.tw.
4. MiniMental.tw.
5. mini mental stat*.tw.
6. or/1-5`

	cq, err := cqrPipeline.Execute(rawQuery)
	if err != nil {
		t.Fatal(err)
	}

	repr, err := cq.Representation()
	if err != nil {
		t.Fatal(err)
	}

	q := repr.(cqr.CommonQueryRepresentation)

	t.Log(q.String())
	t.Log(q.String())
	t.Log(q.String())
	t.Log(q.String())
	t.Log(combinator.HashCQR(q))
	t.Log(combinator.HashCQR(q))
	t.Log(combinator.HashCQR(q))
	t.Log(combinator.HashCQR(q))
}
