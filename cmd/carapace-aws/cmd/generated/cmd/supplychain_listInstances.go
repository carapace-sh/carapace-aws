package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var supplychain_listInstancesCmd = &cobra.Command{
	Use:   "list-instances",
	Short: "List all Amazon Web Services Supply Chain instances for a specific account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(supplychain_listInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(supplychain_listInstancesCmd).Standalone()

		supplychain_listInstancesCmd.Flags().String("instance-name-filter", "", "The filter to ListInstances based on their names.")
		supplychain_listInstancesCmd.Flags().String("instance-state-filter", "", "The filter to ListInstances based on their state.")
		supplychain_listInstancesCmd.Flags().String("max-results", "", "Specify the maximum number of instances to fetch in this paginated request.")
		supplychain_listInstancesCmd.Flags().String("next-token", "", "The pagination token to fetch the next page of instances.")
	})
	supplychainCmd.AddCommand(supplychain_listInstancesCmd)
}
