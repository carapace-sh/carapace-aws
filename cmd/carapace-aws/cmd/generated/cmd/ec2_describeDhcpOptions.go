package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeDhcpOptionsCmd = &cobra.Command{
	Use:   "describe-dhcp-options",
	Short: "Describes your DHCP option sets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeDhcpOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeDhcpOptionsCmd).Standalone()

		ec2_describeDhcpOptionsCmd.Flags().String("dhcp-options-ids", "", "The IDs of DHCP option sets.")
		ec2_describeDhcpOptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeDhcpOptionsCmd.Flags().String("filters", "", "The filters.")
		ec2_describeDhcpOptionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		ec2_describeDhcpOptionsCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
		ec2_describeDhcpOptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeDhcpOptionsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeDhcpOptionsCmd)
}
