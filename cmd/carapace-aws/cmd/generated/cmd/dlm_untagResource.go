package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dlm_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the specified tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dlm_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dlm_untagResourceCmd).Standalone()

		dlm_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		dlm_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys.")
		dlm_untagResourceCmd.MarkFlagRequired("resource-arn")
		dlm_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	dlmCmd.AddCommand(dlm_untagResourceCmd)
}
