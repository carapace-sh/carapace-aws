package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_getRuleCmd = &cobra.Command{
	Use:   "get-rule",
	Short: "Retrieves information about the specified listener rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_getRuleCmd).Standalone()

	vpcLattice_getRuleCmd.Flags().String("listener-identifier", "", "The ID or ARN of the listener.")
	vpcLattice_getRuleCmd.Flags().String("rule-identifier", "", "The ID or ARN of the listener rule.")
	vpcLattice_getRuleCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
	vpcLattice_getRuleCmd.MarkFlagRequired("listener-identifier")
	vpcLattice_getRuleCmd.MarkFlagRequired("rule-identifier")
	vpcLattice_getRuleCmd.MarkFlagRequired("service-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_getRuleCmd)
}
