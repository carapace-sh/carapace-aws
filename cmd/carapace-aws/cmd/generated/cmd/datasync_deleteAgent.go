package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_deleteAgentCmd = &cobra.Command{
	Use:   "delete-agent",
	Short: "Removes an DataSync agent resource from your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_deleteAgentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_deleteAgentCmd).Standalone()

		datasync_deleteAgentCmd.Flags().String("agent-arn", "", "The Amazon Resource Name (ARN) of the agent to delete.")
		datasync_deleteAgentCmd.MarkFlagRequired("agent-arn")
	})
	datasyncCmd.AddCommand(datasync_deleteAgentCmd)
}
