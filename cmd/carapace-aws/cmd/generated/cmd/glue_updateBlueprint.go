package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateBlueprintCmd = &cobra.Command{
	Use:   "update-blueprint",
	Short: "Updates a registered blueprint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateBlueprintCmd).Standalone()

	glue_updateBlueprintCmd.Flags().String("blueprint-location", "", "Specifies a path in Amazon S3 where the blueprint is published.")
	glue_updateBlueprintCmd.Flags().String("description", "", "A description of the blueprint.")
	glue_updateBlueprintCmd.Flags().String("name", "", "The name of the blueprint.")
	glue_updateBlueprintCmd.MarkFlagRequired("blueprint-location")
	glue_updateBlueprintCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_updateBlueprintCmd)
}
