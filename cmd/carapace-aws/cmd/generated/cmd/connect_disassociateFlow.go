package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_disassociateFlowCmd = &cobra.Command{
	Use:   "disassociate-flow",
	Short: "Disassociates a connect resource from a flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_disassociateFlowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_disassociateFlowCmd).Standalone()

		connect_disassociateFlowCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_disassociateFlowCmd.Flags().String("resource-id", "", "The identifier of the resource.")
		connect_disassociateFlowCmd.Flags().String("resource-type", "", "A valid resource type.")
		connect_disassociateFlowCmd.MarkFlagRequired("instance-id")
		connect_disassociateFlowCmd.MarkFlagRequired("resource-id")
		connect_disassociateFlowCmd.MarkFlagRequired("resource-type")
	})
	connectCmd.AddCommand(connect_disassociateFlowCmd)
}
