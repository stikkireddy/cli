package version

import (
	"github.com/databricks/cli/internal/build"
	"github.com/databricks/cli/libs/cmdio"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "version",
		Args: cobra.NoArgs,

		Annotations: map[string]string{
			"template": "Databricks CLI v{{.Version}}\n",
		},
	}

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return cmdio.Render(cmd.Context(), build.GetInfo())
	}

	return cmd
}
