package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteBlueprintCmd = &cobra.Command{
	Use:   "delete-blueprint",
	Short: "Deletes an existing blueprint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteBlueprintCmd).Standalone()

	glue_deleteBlueprintCmd.Flags().String("name", "", "The name of the blueprint to delete.")
	glue_deleteBlueprintCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_deleteBlueprintCmd)
}
