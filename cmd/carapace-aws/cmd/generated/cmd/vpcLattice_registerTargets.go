package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_registerTargetsCmd = &cobra.Command{
	Use:   "register-targets",
	Short: "Registers the targets with the target group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_registerTargetsCmd).Standalone()

	vpcLattice_registerTargetsCmd.Flags().String("target-group-identifier", "", "The ID or ARN of the target group.")
	vpcLattice_registerTargetsCmd.Flags().String("targets", "", "The targets.")
	vpcLattice_registerTargetsCmd.MarkFlagRequired("target-group-identifier")
	vpcLattice_registerTargetsCmd.MarkFlagRequired("targets")
	vpcLatticeCmd.AddCommand(vpcLattice_registerTargetsCmd)
}
