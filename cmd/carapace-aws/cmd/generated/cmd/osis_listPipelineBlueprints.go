package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var osis_listPipelineBlueprintsCmd = &cobra.Command{
	Use:   "list-pipeline-blueprints",
	Short: "Retrieves a list of all available blueprints for Data Prepper.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(osis_listPipelineBlueprintsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(osis_listPipelineBlueprintsCmd).Standalone()

	})
	osisCmd.AddCommand(osis_listPipelineBlueprintsCmd)
}
