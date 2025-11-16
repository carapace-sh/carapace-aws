package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags for a resource in CodeCommit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_untagResourceCmd).Standalone()

		codecommit_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to which you want to remove tags.")
		codecommit_untagResourceCmd.Flags().String("tag-keys", "", "The tag key for each tag that you want to remove from the resource.")
		codecommit_untagResourceCmd.MarkFlagRequired("resource-arn")
		codecommit_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	codecommitCmd.AddCommand(codecommit_untagResourceCmd)
}
