package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes specified tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_untagResourceCmd).Standalone()

		billingconductor_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to which to delete tags.")
		billingconductor_untagResourceCmd.Flags().String("tag-keys", "", "The tags to delete from the resource as a list of key-value pairs.")
		billingconductor_untagResourceCmd.MarkFlagRequired("resource-arn")
		billingconductor_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	billingconductorCmd.AddCommand(billingconductor_untagResourceCmd)
}
