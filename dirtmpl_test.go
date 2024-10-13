package dirtmpl_test

import (
	"os"
	"testing"

	"github.com/imantung/dirtmpl"
	"github.com/stretchr/testify/require"
)

func TestTemplateDirectory_Entries(t *testing.T) {
	testcases := []struct {
		TestName    string
		TemplateDir dirtmpl.TemplateDir
		ExpectedErr string
	}{
		{
			TestName:    "empty root",
			TemplateDir: dirtmpl.TemplateDir{},
			ExpectedErr: "'Root' can't be empty",
		},
		{
			TestName:    "empty read dir",
			TemplateDir: dirtmpl.TemplateDir{Root: "root"},
			ExpectedErr: "'ReadDir' can't be nil",
		},
		{
			TestName:    "empty base prefix",
			TemplateDir: dirtmpl.TemplateDir{Root: "root", ReadDir: os.ReadDir},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			_, err := tt.TemplateDir.Entries()
			if err != nil {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
			}

		})
	}
}
