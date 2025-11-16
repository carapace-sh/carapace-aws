package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeFlowLogsCmd = &cobra.Command{
	Use:   "describe-flow-logs",
	Short: "Describes one or more flow logs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeFlowLogsCmd).Standalone()

	ec2_describeFlowLogsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeFlowLogsCmd.Flags().String("filter", "", "One or more filters.")
	ec2_describeFlowLogsCmd.Flags().String("flow-log-ids", "", "One or more flow log IDs.")
	ec2_describeFlowLogsCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	ec2_describeFlowLogsCmd.Flags().String("next-token", "", "The token to request the next page of items.")
	ec2_describeFlowLogsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeFlowLogsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeFlowLogsCmd)
}
