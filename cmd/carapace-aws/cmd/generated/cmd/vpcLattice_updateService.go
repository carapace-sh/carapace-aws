package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_updateServiceCmd = &cobra.Command{
	Use:   "update-service",
	Short: "Updates the specified service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_updateServiceCmd).Standalone()

	vpcLattice_updateServiceCmd.Flags().String("auth-type", "", "The type of IAM policy.")
	vpcLattice_updateServiceCmd.Flags().String("certificate-arn", "", "The Amazon Resource Name (ARN) of the certificate.")
	vpcLattice_updateServiceCmd.Flags().String("service-identifier", "", "The ID or ARN of the service.")
	vpcLattice_updateServiceCmd.MarkFlagRequired("service-identifier")
	vpcLatticeCmd.AddCommand(vpcLattice_updateServiceCmd)
}
