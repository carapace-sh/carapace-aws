package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeOutpostLagsCmd = &cobra.Command{
	Use:   "describe-outpost-lags",
	Short: "Describes the Outposts link aggregation groups (LAGs).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeOutpostLagsCmd).Standalone()

	ec2_describeOutpostLagsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeOutpostLagsCmd.Flags().String("filters", "", "The filters to use for narrowing down the request.")
	ec2_describeOutpostLagsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeOutpostLagsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeOutpostLagsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeOutpostLagsCmd.Flags().String("outpost-lag-ids", "", "The IDs of the Outpost LAGs.")
	ec2_describeOutpostLagsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeOutpostLagsCmd)
}
