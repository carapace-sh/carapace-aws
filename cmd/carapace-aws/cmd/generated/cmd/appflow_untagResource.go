package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from the specified flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_untagResourceCmd).Standalone()

		appflow_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the flow that you want to untag.")
		appflow_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys associated with the tag that you want to remove from your flow.")
		appflow_untagResourceCmd.MarkFlagRequired("resource-arn")
		appflow_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	appflowCmd.AddCommand(appflow_untagResourceCmd)
}
