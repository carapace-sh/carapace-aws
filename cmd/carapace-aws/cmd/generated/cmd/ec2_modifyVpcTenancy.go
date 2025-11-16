package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyVpcTenancyCmd = &cobra.Command{
	Use:   "modify-vpc-tenancy",
	Short: "Modifies the instance tenancy attribute of the specified VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyVpcTenancyCmd).Standalone()

	ec2_modifyVpcTenancyCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVpcTenancyCmd.Flags().String("instance-tenancy", "", "The instance tenancy attribute for the VPC.")
	ec2_modifyVpcTenancyCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyVpcTenancyCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
	ec2_modifyVpcTenancyCmd.MarkFlagRequired("instance-tenancy")
	ec2_modifyVpcTenancyCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyVpcTenancyCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_modifyVpcTenancyCmd)
}
