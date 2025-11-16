package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags for a provided resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listTagsForResourceCmd).Standalone()

	cleanroomsml_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you are interested in.")
	cleanroomsml_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	cleanroomsmlCmd.AddCommand(cleanroomsml_listTagsForResourceCmd)
}
