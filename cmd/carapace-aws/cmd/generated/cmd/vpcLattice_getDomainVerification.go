package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getDomainVerificationCmd = &cobra.Command{
	Use:   "get-domain-verification",
	Short: "Retrieves information about a domain verification.ß",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getDomainVerificationCmd).Standalone()

	vpcLattice_getDomainVerificationCmd.Flags().String("domain-verification-identifier", "", "The ID or ARN of the domain verification to retrieve.")
	vpcLattice_getDomainVerificationCmd.MarkFlagRequired("domain-verification-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_getDomainVerificationCmd)
}
