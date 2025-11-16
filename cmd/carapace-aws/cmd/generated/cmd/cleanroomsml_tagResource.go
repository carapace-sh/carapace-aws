package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds metadata tags to a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_tagResourceCmd).Standalone()

	cleanroomsml_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to assign tags.")
	cleanroomsml_tagResourceCmd.Flags().String("tags", "", "The optional metadata that you apply to the resource to help you categorize and organize them.")
	cleanroomsml_tagResourceCmd.MarkFlagRequired("resource-arn")
	cleanroomsml_tagResourceCmd.MarkFlagRequired("tags")
	cleanroomsmlCmd.AddCommand(cleanroomsml_tagResourceCmd)
}
