package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlake_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove a user-specifed key and value tag from a data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlake_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(healthlake_untagResourceCmd).Standalone()

		healthlake_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the data store from which tags are being removed.")
		healthlake_untagResourceCmd.Flags().String("tag-keys", "", "The keys for the tags to be removed from the data store.")
		healthlake_untagResourceCmd.MarkFlagRequired("resource-arn")
		healthlake_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	healthlakeCmd.AddCommand(healthlake_untagResourceCmd)
}
