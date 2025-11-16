package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateAgentCmd = &cobra.Command{
	Use:   "update-agent",
	Short: "Updates the name of an DataSync agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateAgentCmd).Standalone()

	datasync_updateAgentCmd.Flags().String("agent-arn", "", "The Amazon Resource Name (ARN) of the agent to update.")
	datasync_updateAgentCmd.Flags().String("name", "", "The name that you want to use to configure the agent.")
	datasync_updateAgentCmd.MarkFlagRequired("agent-arn")
	datasyncCmd.AddCommand(datasync_updateAgentCmd)
}
