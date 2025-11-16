package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getBlueprintRunCmd = &cobra.Command{
	Use:   "get-blueprint-run",
	Short: "Retrieves the details of a blueprint run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getBlueprintRunCmd).Standalone()

	glue_getBlueprintRunCmd.Flags().String("blueprint-name", "", "The name of the blueprint.")
	glue_getBlueprintRunCmd.Flags().String("run-id", "", "The run ID for the blueprint run you want to retrieve.")
	glue_getBlueprintRunCmd.MarkFlagRequired("blueprint-name")
	glue_getBlueprintRunCmd.MarkFlagRequired("run-id")
	glueCmd.AddCommand(glue_getBlueprintRunCmd)
}
