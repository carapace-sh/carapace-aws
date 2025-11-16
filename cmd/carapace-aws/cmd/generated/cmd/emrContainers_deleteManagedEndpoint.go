package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_deleteManagedEndpointCmd = &cobra.Command{
	Use:   "delete-managed-endpoint",
	Short: "Deletes a managed endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_deleteManagedEndpointCmd).Standalone()

	emrContainers_deleteManagedEndpointCmd.Flags().String("id", "", "The ID of the managed endpoint.")
	emrContainers_deleteManagedEndpointCmd.Flags().String("virtual-cluster-id", "", "The ID of the endpoint's virtual cluster.")
	emrContainers_deleteManagedEndpointCmd.MarkFlagRequired("id")
	emrContainers_deleteManagedEndpointCmd.MarkFlagRequired("virtual-cluster-id")
	emrContainersCmd.AddCommand(emrContainers_deleteManagedEndpointCmd)
}
