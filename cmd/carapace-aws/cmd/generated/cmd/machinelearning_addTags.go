package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_addTagsCmd = &cobra.Command{
	Use:   "add-tags",
	Short: "Adds one or more tags to an object, up to a limit of 10. Each tag consists of a key and an optional value.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_addTagsCmd).Standalone()

	machinelearning_addTagsCmd.Flags().String("resource-id", "", "The ID of the ML object to tag.")
	machinelearning_addTagsCmd.Flags().String("resource-type", "", "The type of the ML object to tag.")
	machinelearning_addTagsCmd.Flags().String("tags", "", "The key-value pairs to use to create tags.")
	machinelearning_addTagsCmd.MarkFlagRequired("resource-id")
	machinelearning_addTagsCmd.MarkFlagRequired("resource-type")
	machinelearning_addTagsCmd.MarkFlagRequired("tags")
	machinelearningCmd.AddCommand(machinelearning_addTagsCmd)
}
