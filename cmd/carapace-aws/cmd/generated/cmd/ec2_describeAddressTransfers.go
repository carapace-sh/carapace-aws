package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeAddressTransfersCmd = &cobra.Command{
	Use:   "describe-address-transfers",
	Short: "Describes an Elastic IP address transfer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeAddressTransfersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeAddressTransfersCmd).Standalone()

		ec2_describeAddressTransfersCmd.Flags().String("allocation-ids", "", "The allocation IDs of Elastic IP addresses.")
		ec2_describeAddressTransfersCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeAddressTransfersCmd.Flags().String("max-results", "", "The maximum number of address transfers to return in one page of results.")
		ec2_describeAddressTransfersCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		ec2_describeAddressTransfersCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeAddressTransfersCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeAddressTransfersCmd)
}
