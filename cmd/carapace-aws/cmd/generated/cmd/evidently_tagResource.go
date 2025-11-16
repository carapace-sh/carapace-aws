package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified CloudWatch Evidently resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_tagResourceCmd).Standalone()

	evidently_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the CloudWatch Evidently resource that you're adding tags to.")
	evidently_tagResourceCmd.Flags().String("tags", "", "The list of key-value pairs to associate with the resource.")
	evidently_tagResourceCmd.MarkFlagRequired("resource-arn")
	evidently_tagResourceCmd.MarkFlagRequired("tags")
	evidentlyCmd.AddCommand(evidently_tagResourceCmd)
}
