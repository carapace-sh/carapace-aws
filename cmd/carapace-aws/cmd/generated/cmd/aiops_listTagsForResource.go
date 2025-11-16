package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aiops_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Displays the tags associated with a CloudWatch investigations resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aiops_listTagsForResourceCmd).Standalone()

	aiops_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the CloudWatch investigations resource that you want to view tags for.")
	aiops_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	aiopsCmd.AddCommand(aiops_listTagsForResourceCmd)
}
