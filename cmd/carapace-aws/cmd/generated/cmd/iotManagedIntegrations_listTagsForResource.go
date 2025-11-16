package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotManagedIntegrations_listTagsForResourceCmd).Standalone()

		iotManagedIntegrations_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource for which to list tags.")
		iotManagedIntegrations_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_listTagsForResourceCmd)
}
