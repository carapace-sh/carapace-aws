package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listListenersCmd = &cobra.Command{
	Use:   "list-listeners",
	Short: "Lists the listeners for the specified service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listListenersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_listListenersCmd).Standalone()

		vpcLattice_listListenersCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		vpcLattice_listListenersCmd.Flags().String("next-token", "", "A pagination token for the next page of results.")
		vpcLattice_listListenersCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
		vpcLattice_listListenersCmd.MarkFlagRequired("service-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_listListenersCmd)
}
