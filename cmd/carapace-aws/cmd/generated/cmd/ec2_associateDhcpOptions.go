package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_associateDhcpOptionsCmd = &cobra.Command{
	Use:   "associate-dhcp-options",
	Short: "Associates a set of DHCP options (that you've previously created) with the specified VPC, or associates no DHCP options with the VPC.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_associateDhcpOptionsCmd).Standalone()

	ec2_associateDhcpOptionsCmd.Flags().String("dhcp-options-id", "", "The ID of the DHCP options set, or `default` to associate no DHCP options with the VPC.")
	ec2_associateDhcpOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateDhcpOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_associateDhcpOptionsCmd.Flags().String("vpc-id", "", "The ID of the VPC.")
	ec2_associateDhcpOptionsCmd.MarkFlagRequired("dhcp-options-id")
	ec2_associateDhcpOptionsCmd.Flag("no-dry-run").Hidden = true
	ec2_associateDhcpOptionsCmd.MarkFlagRequired("vpc-id")
	ec2Cmd.AddCommand(ec2_associateDhcpOptionsCmd)
}
