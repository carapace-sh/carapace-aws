package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteDhcpOptionsCmd = &cobra.Command{
	Use:   "delete-dhcp-options",
	Short: "Deletes the specified set of DHCP options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteDhcpOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteDhcpOptionsCmd).Standalone()

		ec2_deleteDhcpOptionsCmd.Flags().String("dhcp-options-id", "", "The ID of the DHCP options set.")
		ec2_deleteDhcpOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteDhcpOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteDhcpOptionsCmd.MarkFlagRequired("dhcp-options-id")
		ec2_deleteDhcpOptionsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteDhcpOptionsCmd)
}
