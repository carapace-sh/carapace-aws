package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_tagResourceCmd).Standalone()

		securityhub_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to apply the tags to.")
		securityhub_tagResourceCmd.Flags().String("tags", "", "The tags to add to the resource.")
		securityhub_tagResourceCmd.MarkFlagRequired("resource-arn")
		securityhub_tagResourceCmd.MarkFlagRequired("tags")
	})
	securityhubCmd.AddCommand(securityhub_tagResourceCmd)
}
