package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_associateFlowCmd = &cobra.Command{
	Use:   "associate-flow",
	Short: "Associates a connect resource to a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_associateFlowCmd).Standalone()

	connect_associateFlowCmd.Flags().String("flow-id", "", "The identifier of the flow.")
	connect_associateFlowCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_associateFlowCmd.Flags().String("resource-id", "", "The identifier of the resource.")
	connect_associateFlowCmd.Flags().String("resource-type", "", "A valid resource type.")
	connect_associateFlowCmd.MarkFlagRequired("flow-id")
	connect_associateFlowCmd.MarkFlagRequired("instance-id")
	connect_associateFlowCmd.MarkFlagRequired("resource-id")
	connect_associateFlowCmd.MarkFlagRequired("resource-type")
	connectCmd.AddCommand(connect_associateFlowCmd)
}
