package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listServicesCmd = &cobra.Command{
	Use:   "list-services",
	Short: "Lists the services owned by the caller account or shared with the caller account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listServicesCmd).Standalone()

	vpcLattice_listServicesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	vpcLattice_listServicesCmd.Flags().String("next-token", "", "A pagination token for the next page of results.")
	vpcLatticeCmd.AddCommand(vpcLattice_listServicesCmd)
}
