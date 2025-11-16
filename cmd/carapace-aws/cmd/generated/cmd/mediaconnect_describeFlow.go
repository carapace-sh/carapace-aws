package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_describeFlowCmd = &cobra.Command{
	Use:   "describe-flow",
	Short: "Displays the details of a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_describeFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_describeFlowCmd).Standalone()

		mediaconnect_describeFlowCmd.Flags().String("flow-arn", "", "The ARN of the flow that you want to describe.")
		mediaconnect_describeFlowCmd.MarkFlagRequired("flow-arn")
	})
	mediaconnectCmd.AddCommand(mediaconnect_describeFlowCmd)
}
