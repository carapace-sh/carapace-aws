package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_listManagedEndpointsCmd = &cobra.Command{
	Use:   "list-managed-endpoints",
	Short: "Lists managed endpoints based on a set of parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_listManagedEndpointsCmd).Standalone()

	emrContainers_listManagedEndpointsCmd.Flags().String("created-after", "", "The date and time after which the endpoints are created.")
	emrContainers_listManagedEndpointsCmd.Flags().String("created-before", "", "The date and time before which the endpoints are created.")
	emrContainers_listManagedEndpointsCmd.Flags().String("max-results", "", "The maximum number of managed endpoints that can be listed.")
	emrContainers_listManagedEndpointsCmd.Flags().String("next-token", "", "The token for the next set of managed endpoints to return.")
	emrContainers_listManagedEndpointsCmd.Flags().String("states", "", "The states of the managed endpoints.")
	emrContainers_listManagedEndpointsCmd.Flags().String("types", "", "The types of the managed endpoints.")
	emrContainers_listManagedEndpointsCmd.Flags().String("virtual-cluster-id", "", "The ID of the virtual cluster.")
	emrContainers_listManagedEndpointsCmd.MarkFlagRequired("virtual-cluster-id")
	emrContainersCmd.AddCommand(emrContainers_listManagedEndpointsCmd)
}
