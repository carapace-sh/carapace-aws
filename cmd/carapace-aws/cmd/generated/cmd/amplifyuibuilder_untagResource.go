package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Untags a resource with a specified Amazon Resource Name (ARN).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifyuibuilder_untagResourceCmd).Standalone()

		amplifyuibuilder_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) to use to untag a resource.")
		amplifyuibuilder_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to use to untag a resource.")
		amplifyuibuilder_untagResourceCmd.MarkFlagRequired("resource-arn")
		amplifyuibuilder_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_untagResourceCmd)
}
