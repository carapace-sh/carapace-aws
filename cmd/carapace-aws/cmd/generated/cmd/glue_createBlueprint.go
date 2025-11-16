package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createBlueprintCmd = &cobra.Command{
	Use:   "create-blueprint",
	Short: "Registers a blueprint with Glue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createBlueprintCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_createBlueprintCmd).Standalone()

		glue_createBlueprintCmd.Flags().String("blueprint-location", "", "Specifies a path in Amazon S3 where the blueprint is published.")
		glue_createBlueprintCmd.Flags().String("description", "", "A description of the blueprint.")
		glue_createBlueprintCmd.Flags().String("name", "", "The name of the blueprint.")
		glue_createBlueprintCmd.Flags().String("tags", "", "The tags to be applied to this blueprint.")
		glue_createBlueprintCmd.MarkFlagRequired("blueprint-location")
		glue_createBlueprintCmd.MarkFlagRequired("name")
	})
	glueCmd.AddCommand(glue_createBlueprintCmd)
}
