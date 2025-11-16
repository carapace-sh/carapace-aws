package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Disassociates a resource from a list of tags.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codedeploy_untagResourceCmd).Standalone()

		codedeploy_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that specifies from which resource to disassociate the tags with the keys in the `TagKeys` input parameter.")
		codedeploy_untagResourceCmd.Flags().String("tag-keys", "", "A list of keys of `Tag` objects.")
		codedeploy_untagResourceCmd.MarkFlagRequired("resource-arn")
		codedeploy_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	codedeployCmd.AddCommand(codedeploy_untagResourceCmd)
}
