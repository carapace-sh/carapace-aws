package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteFlowLogsCmd = &cobra.Command{
	Use:   "delete-flow-logs",
	Short: "Deletes one or more flow logs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteFlowLogsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteFlowLogsCmd).Standalone()

		ec2_deleteFlowLogsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteFlowLogsCmd.Flags().String("flow-log-ids", "", "One or more flow log IDs.")
		ec2_deleteFlowLogsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteFlowLogsCmd.MarkFlagRequired("flow-log-ids")
		ec2_deleteFlowLogsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteFlowLogsCmd)
}
