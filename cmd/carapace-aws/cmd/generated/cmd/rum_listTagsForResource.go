package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Displays the tags associated with a CloudWatch RUM resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rum_listTagsForResourceCmd).Standalone()

		rum_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource that you want to see the tags of.")
		rum_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	rumCmd.AddCommand(rum_listTagsForResourceCmd)
}
