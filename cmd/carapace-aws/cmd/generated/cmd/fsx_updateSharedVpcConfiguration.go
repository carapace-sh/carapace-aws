package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_updateSharedVpcConfigurationCmd = &cobra.Command{
	Use:   "update-shared-vpc-configuration",
	Short: "Configures whether participant accounts in your organization can create Amazon FSx for NetApp ONTAP Multi-AZ file systems in subnets that are shared by a virtual private cloud (VPC) owner.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_updateSharedVpcConfigurationCmd).Standalone()

	fsx_updateSharedVpcConfigurationCmd.Flags().String("client-request-token", "", "")
	fsx_updateSharedVpcConfigurationCmd.Flags().String("enable-fsx-route-table-updates-from-participant-accounts", "", "Specifies whether participant accounts can create FSx for ONTAP Multi-AZ file systems in shared subnets.")
	fsxCmd.AddCommand(fsx_updateSharedVpcConfigurationCmd)
}
