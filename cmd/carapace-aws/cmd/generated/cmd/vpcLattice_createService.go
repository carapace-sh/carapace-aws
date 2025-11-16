package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vpcLattice_createServiceCmd = &cobra.Command{
	Use:   "create-service",
	Short: "Creates a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vpcLattice_createServiceCmd).Standalone()

	vpcLattice_createServiceCmd.Flags().String("auth-type", "", "The type of IAM policy.")
	vpcLattice_createServiceCmd.Flags().String("certificate-arn", "", "The Amazon Resource Name (ARN) of the certificate.")
	vpcLattice_createServiceCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	vpcLattice_createServiceCmd.Flags().String("custom-domain-name", "", "The custom domain name of the service.")
	vpcLattice_createServiceCmd.Flags().String("name", "", "The name of the service.")
	vpcLattice_createServiceCmd.Flags().String("tags", "", "The tags for the service.")
	vpcLattice_createServiceCmd.MarkFlagRequired("name")
	vpcLatticeCmd.AddCommand(vpcLattice_createServiceCmd)
}
