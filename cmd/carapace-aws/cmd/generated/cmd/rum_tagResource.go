package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified CloudWatch RUM resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_tagResourceCmd).Standalone()

	rum_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the CloudWatch RUM resource that you're adding tags to.")
	rum_tagResourceCmd.Flags().String("tags", "", "The list of key-value pairs to associate with the resource.")
	rum_tagResourceCmd.MarkFlagRequired("resource-arn")
	rum_tagResourceCmd.MarkFlagRequired("tags")
	rumCmd.AddCommand(rum_tagResourceCmd)
}
