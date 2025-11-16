package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_describeManagedEndpointCmd = &cobra.Command{
	Use:   "describe-managed-endpoint",
	Short: "Displays detailed information about a managed endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_describeManagedEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrContainers_describeManagedEndpointCmd).Standalone()

		emrContainers_describeManagedEndpointCmd.Flags().String("id", "", "This output displays ID of the managed endpoint.")
		emrContainers_describeManagedEndpointCmd.Flags().String("virtual-cluster-id", "", "The ID of the endpoint's virtual cluster.")
		emrContainers_describeManagedEndpointCmd.MarkFlagRequired("id")
		emrContainers_describeManagedEndpointCmd.MarkFlagRequired("virtual-cluster-id")
	})
	emrContainersCmd.AddCommand(emrContainers_describeManagedEndpointCmd)
}
