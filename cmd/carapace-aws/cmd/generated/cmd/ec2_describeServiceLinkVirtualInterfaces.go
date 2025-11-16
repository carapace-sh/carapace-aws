package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeServiceLinkVirtualInterfacesCmd = &cobra.Command{
	Use:   "describe-service-link-virtual-interfaces",
	Short: "Describes the Outpost service link virtual interfaces.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeServiceLinkVirtualInterfacesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeServiceLinkVirtualInterfacesCmd).Standalone()

		ec2_describeServiceLinkVirtualInterfacesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeServiceLinkVirtualInterfacesCmd.Flags().String("filters", "", "The filters to use for narrowing down the request.")
		ec2_describeServiceLinkVirtualInterfacesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeServiceLinkVirtualInterfacesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeServiceLinkVirtualInterfacesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeServiceLinkVirtualInterfacesCmd.Flags().String("service-link-virtual-interface-ids", "", "The IDs of the service link virtual interfaces.")
		ec2_describeServiceLinkVirtualInterfacesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeServiceLinkVirtualInterfacesCmd)
}
