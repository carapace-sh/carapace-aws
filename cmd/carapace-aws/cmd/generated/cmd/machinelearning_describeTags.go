package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machinelearning_describeTagsCmd = &cobra.Command{
	Use:   "describe-tags",
	Short: "Describes one or more of the tags for your Amazon ML object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machinelearning_describeTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(machinelearning_describeTagsCmd).Standalone()

		machinelearning_describeTagsCmd.Flags().String("resource-id", "", "The ID of the ML object.")
		machinelearning_describeTagsCmd.Flags().String("resource-type", "", "The type of the ML object.")
		machinelearning_describeTagsCmd.MarkFlagRequired("resource-id")
		machinelearning_describeTagsCmd.MarkFlagRequired("resource-type")
	})
	machinelearningCmd.AddCommand(machinelearning_describeTagsCmd)
}
