package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createDhcpOptionsCmd = &cobra.Command{
	Use:   "create-dhcp-options",
	Short: "Creates a custom set of DHCP options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createDhcpOptionsCmd).Standalone()

	ec2_createDhcpOptionsCmd.Flags().String("dhcp-configurations", "", "A DHCP configuration option.")
	ec2_createDhcpOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createDhcpOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createDhcpOptionsCmd.Flags().String("tag-specifications", "", "The tags to assign to the DHCP option.")
	ec2_createDhcpOptionsCmd.MarkFlagRequired("dhcp-configurations")
	ec2_createDhcpOptionsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_createDhcpOptionsCmd)
}
