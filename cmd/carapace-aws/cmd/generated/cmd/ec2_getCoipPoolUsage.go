package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getCoipPoolUsageCmd = &cobra.Command{
	Use:   "get-coip-pool-usage",
	Short: "Describes the allocations from the specified customer-owned address pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getCoipPoolUsageCmd).Standalone()

	ec2_getCoipPoolUsageCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getCoipPoolUsageCmd.Flags().String("filters", "", "One or more filters.")
	ec2_getCoipPoolUsageCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_getCoipPoolUsageCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_getCoipPoolUsageCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getCoipPoolUsageCmd.Flags().String("pool-id", "", "The ID of the address pool.")
	ec2_getCoipPoolUsageCmd.Flag("no-dry-run").Hidden = true
	ec2_getCoipPoolUsageCmd.MarkFlagRequired("pool-id")
	ec2Cmd.AddCommand(ec2_getCoipPoolUsageCmd)
}
