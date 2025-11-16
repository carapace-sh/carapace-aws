package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_untagResourceCmd).Standalone()

		iotManagedIntegrations_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to which to add tags.")
		iotManagedIntegrations_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys to remove from the resource.")
		iotManagedIntegrations_untagResourceCmd.MarkFlagRequired("resource-arn")
		iotManagedIntegrations_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_untagResourceCmd)
}
