package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_describeAgentCmd = &cobra.Command{
	Use:   "describe-agent",
	Short: "Returns information about an DataSync agent, such as its name, service endpoint type, and status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_describeAgentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_describeAgentCmd).Standalone()

		datasync_describeAgentCmd.Flags().String("agent-arn", "", "Specifies the Amazon Resource Name (ARN) of the DataSync agent that you want information about.")
		datasync_describeAgentCmd.MarkFlagRequired("agent-arn")
	})
	datasyncCmd.AddCommand(datasync_describeAgentCmd)
}
