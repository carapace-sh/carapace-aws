package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_deleteTagsCmd = &cobra.Command{
	Use:   "delete-tags",
	Short: "Deletes the specified tags associated with an ML object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_deleteTagsCmd).Standalone()

	machinelearning_deleteTagsCmd.Flags().String("resource-id", "", "The ID of the tagged ML object.")
	machinelearning_deleteTagsCmd.Flags().String("resource-type", "", "The type of the tagged ML object.")
	machinelearning_deleteTagsCmd.Flags().String("tag-keys", "", "One or more tags to delete.")
	machinelearning_deleteTagsCmd.MarkFlagRequired("resource-id")
	machinelearning_deleteTagsCmd.MarkFlagRequired("resource-type")
	machinelearning_deleteTagsCmd.MarkFlagRequired("tag-keys")
	machinelearningCmd.AddCommand(machinelearning_deleteTagsCmd)
}
