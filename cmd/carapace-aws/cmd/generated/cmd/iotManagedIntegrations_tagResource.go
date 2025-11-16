package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotManagedIntegrations_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotManagedIntegrations_tagResourceCmd).Standalone()

	iotManagedIntegrations_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to which to add tags.")
	iotManagedIntegrations_tagResourceCmd.Flags().String("tags", "", "A set of key/value pairs that are used to manage the resource")
	iotManagedIntegrations_tagResourceCmd.MarkFlagRequired("resource-arn")
	iotManagedIntegrations_tagResourceCmd.MarkFlagRequired("tags")
	iotManagedIntegrationsCmd.AddCommand(iotManagedIntegrations_tagResourceCmd)
}
