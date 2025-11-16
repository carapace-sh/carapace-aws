package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listCommandInvocationsCmd = &cobra.Command{
	Use:   "list-command-invocations",
	Short: "An invocation is copy of a command sent to a specific managed node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listCommandInvocationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_listCommandInvocationsCmd).Standalone()

		ssm_listCommandInvocationsCmd.Flags().String("command-id", "", "(Optional) The invocations for a specific command ID.")
		ssm_listCommandInvocationsCmd.Flags().Bool("details", false, "(Optional) If set this returns the response of the command executions and any command output.")
		ssm_listCommandInvocationsCmd.Flags().String("filters", "", "(Optional) One or more filters.")
		ssm_listCommandInvocationsCmd.Flags().String("instance-id", "", "(Optional) The command execution details for a specific managed node ID.")
		ssm_listCommandInvocationsCmd.Flags().String("max-results", "", "(Optional) The maximum number of items to return for this call.")
		ssm_listCommandInvocationsCmd.Flags().String("next-token", "", "(Optional) The token for the next set of items to return.")
		ssm_listCommandInvocationsCmd.Flags().Bool("no-details", false, "(Optional) If set this returns the response of the command executions and any command output.")
		ssm_listCommandInvocationsCmd.Flag("no-details").Hidden = true
	})
	ssmCmd.AddCommand(ssm_listCommandInvocationsCmd)
}
