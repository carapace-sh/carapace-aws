package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_deleteDomainVerificationCmd = &cobra.Command{
	Use:   "delete-domain-verification",
	Short: "Deletes the specified domain verification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_deleteDomainVerificationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_deleteDomainVerificationCmd).Standalone()

		vpcLattice_deleteDomainVerificationCmd.Flags().String("domain-verification-identifier", "", "The ID of the domain verification to delete.")
		vpcLattice_deleteDomainVerificationCmd.MarkFlagRequired("domain-verification-identifier")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_deleteDomainVerificationCmd)
}
