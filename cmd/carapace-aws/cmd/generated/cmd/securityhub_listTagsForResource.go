package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags associated with a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listTagsForResourceCmd).Standalone()

	securityhub_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to retrieve tags for.")
	securityhub_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	securityhubCmd.AddCommand(securityhub_listTagsForResourceCmd)
}
