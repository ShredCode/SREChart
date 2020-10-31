package parse

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/shredcode/srechart/ast"
)

func newFlows(ents interface{}) ([]ast.TX, error) {
	out := []ast.TX{}
	entries, err := toII(ents)
	if err != nil {
		return out, errors.Wrap(err, "toII")
	}
	for _, ent := range entries {
		tx, err := toTX(ent)
		if err != nil {
			return out, errors.Wrap(err, "toTX")
		}
		out = append(out, tx)
	}
	return out, nil
}


func newFlow(start, end, options interface{}) (ast.TX, error) {
	st, err := toString(start)
	if err != nil {
		return ast.TX{}, errors.Wrap(err, "toString")
	}

	options, err := toOptionss(options)
	if err != nil {
		return ast.TX{}, errors.Wrap(err, "toPostings")
	}
  
  	en, err := toString(end)
	if err != nil {
		return ast.TX{}, errors.Wrap(err, "toString")
	}

	return ast.TX{
		Start:     st,
		End:    en,
		Options:  options,
	}, nil
}

