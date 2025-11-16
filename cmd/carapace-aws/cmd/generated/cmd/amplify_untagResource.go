package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Untags a resource with a specified Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_untagResourceCmd).Standalone()

	amplify_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) to use to untag a resource.")
	amplify_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to use to untag a resource.")
	amplify_untagResourceCmd.MarkFlagRequired("resource-arn")
	amplify_untagResourceCmd.MarkFlagRequired("tag-keys")
	amplifyCmd.AddCommand(amplify_untagResourceCmd)
}
