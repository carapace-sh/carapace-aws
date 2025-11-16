package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates the list of tags in the input `Tags` parameter with the resource identified by the `ResourceArn` input parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_tagResourceCmd).Standalone()

		codedeploy_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of a resource, such as a CodeDeploy application or deployment group.")
		codedeploy_tagResourceCmd.Flags().String("tags", "", "A list of tags that `TagResource` associates with a resource.")
		codedeploy_tagResourceCmd.MarkFlagRequired("resource-arn")
		codedeploy_tagResourceCmd.MarkFlagRequired("tags")
	})
	codedeployCmd.AddCommand(codedeploy_tagResourceCmd)
}
