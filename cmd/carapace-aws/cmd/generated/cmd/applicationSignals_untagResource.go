package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_untagResourceCmd).Standalone()

		applicationSignals_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the CloudWatch resource that you want to delete tags from.")
		applicationSignals_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
		applicationSignals_untagResourceCmd.MarkFlagRequired("resource-arn")
		applicationSignals_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	applicationSignalsCmd.AddCommand(applicationSignals_untagResourceCmd)
}
