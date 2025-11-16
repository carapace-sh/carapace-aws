package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_describeSharedVpcConfigurationCmd = &cobra.Command{
	Use:   "describe-shared-vpc-configuration",
	Short: "Indicates whether participant accounts in your organization can create Amazon FSx for NetApp ONTAP Multi-AZ file systems in subnets that are shared by a virtual private cloud (VPC) owner.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_describeSharedVpcConfigurationCmd).Standalone()

	fsxCmd.AddCommand(fsx_describeSharedVpcConfigurationCmd)
}
