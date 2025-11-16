package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeMovingAddressesCmd = &cobra.Command{
	Use:   "describe-moving-addresses",
	Short: "This action is deprecated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeMovingAddressesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeMovingAddressesCmd).Standalone()

		ec2_describeMovingAddressesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeMovingAddressesCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeMovingAddressesCmd.Flags().String("max-results", "", "The maximum number of results to return for the request in a single page.")
		ec2_describeMovingAddressesCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeMovingAddressesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeMovingAddressesCmd.Flags().String("public-ips", "", "One or more Elastic IP addresses.")
		ec2_describeMovingAddressesCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeMovingAddressesCmd)
}
