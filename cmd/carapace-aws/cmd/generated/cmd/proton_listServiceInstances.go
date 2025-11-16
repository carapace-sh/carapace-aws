package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proton_listServiceInstancesCmd = &cobra.Command{
	Use:   "list-service-instances",
	Short: "List service instances with summary data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proton_listServiceInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(proton_listServiceInstancesCmd).Standalone()

		proton_listServiceInstancesCmd.Flags().String("filters", "", "An array of filtering criteria that scope down the result list.")
		proton_listServiceInstancesCmd.Flags().String("max-results", "", "The maximum number of service instances to list.")
		proton_listServiceInstancesCmd.Flags().String("next-token", "", "A token that indicates the location of the next service in the array of service instances, after the list of service instances that was previously requested.")
		proton_listServiceInstancesCmd.Flags().String("service-name", "", "The name of the service that the service instance belongs to.")
		proton_listServiceInstancesCmd.Flags().String("sort-by", "", "The field that the result list is sorted by.")
		proton_listServiceInstancesCmd.Flags().String("sort-order", "", "Result list sort order.")
	})
	protonCmd.AddCommand(proton_listServiceInstancesCmd)
}
