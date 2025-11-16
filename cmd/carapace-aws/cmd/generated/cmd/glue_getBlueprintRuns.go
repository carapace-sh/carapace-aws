package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getBlueprintRunsCmd = &cobra.Command{
	Use:   "get-blueprint-runs",
	Short: "Retrieves the details of blueprint runs for a specified blueprint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getBlueprintRunsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getBlueprintRunsCmd).Standalone()

		glue_getBlueprintRunsCmd.Flags().String("blueprint-name", "", "The name of the blueprint.")
		glue_getBlueprintRunsCmd.Flags().String("max-results", "", "The maximum size of a list to return.")
		glue_getBlueprintRunsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation request.")
		glue_getBlueprintRunsCmd.MarkFlagRequired("blueprint-name")
	})
	glueCmd.AddCommand(glue_getBlueprintRunsCmd)
}
