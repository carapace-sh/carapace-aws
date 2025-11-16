package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_startDomainVerificationCmd = &cobra.Command{
	Use:   "start-domain-verification",
	Short: "Starts the domain verification process for a custom domain name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_startDomainVerificationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(vpcLattice_startDomainVerificationCmd).Standalone()

		vpcLattice_startDomainVerificationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		vpcLattice_startDomainVerificationCmd.Flags().String("domain-name", "", "The domain name to verify ownership for.")
		vpcLattice_startDomainVerificationCmd.Flags().String("tags", "", "The tags for the domain verification.")
		vpcLattice_startDomainVerificationCmd.MarkFlagRequired("domain-name")
	})
	vpcLatticeCmd.AddCommand(vpcLattice_startDomainVerificationCmd)
}
