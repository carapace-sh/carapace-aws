package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLatticeCmd = &cobra.Command{
	Use:   "vpc-lattice",
	Short: "Amazon VPC Lattice is a fully managed application networking service that you use to connect, secure, and monitor all of your services across multiple accounts and virtual private clouds (VPCs).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLatticeCmd).Standalone()

	rootCmd.AddCommand(vpcLatticeCmd)
}
