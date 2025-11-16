package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_listDomainVerificationsCmd = &cobra.Command{
	Use:   "list-domain-verifications",
	Short: "Lists the domain verifications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_listDomainVerificationsCmd).Standalone()

	vpcLattice_listDomainVerificationsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	vpcLattice_listDomainVerificationsCmd.Flags().String("next-token", "", "A pagination token for the next page of results.")
	vpcLatticeCmd.AddCommand(vpcLattice_listDomainVerificationsCmd)
}
