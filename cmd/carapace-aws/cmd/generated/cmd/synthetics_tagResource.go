package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified canary or group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(synthetics_tagResourceCmd).Standalone()

		synthetics_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the canary or group that you're adding tags to.")
		synthetics_tagResourceCmd.Flags().String("tags", "", "The list of key-value pairs to associate with the resource.")
		synthetics_tagResourceCmd.MarkFlagRequired("resource-arn")
		synthetics_tagResourceCmd.MarkFlagRequired("tags")
	})
	syntheticsCmd.AddCommand(synthetics_tagResourceCmd)
}
