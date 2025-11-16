package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags for the resource identified by a specified Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_listTagsForResourceCmd).Standalone()

		codedeploy_listTagsForResourceCmd.Flags().String("next-token", "", "An identifier returned from the previous `ListTagsForResource` call.")
		codedeploy_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of a CodeDeploy resource.")
		codedeploy_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	codedeployCmd.AddCommand(codedeploy_listTagsForResourceCmd)
}
