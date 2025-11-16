package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_deleteWorkloadIdentityCmd = &cobra.Command{
	Use:   "delete-workload-identity",
	Short: "Deletes a workload identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_deleteWorkloadIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_deleteWorkloadIdentityCmd).Standalone()

		bedrockAgentcoreControl_deleteWorkloadIdentityCmd.Flags().String("name", "", "The name of the workload identity to delete.")
		bedrockAgentcoreControl_deleteWorkloadIdentityCmd.MarkFlagRequired("name")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_deleteWorkloadIdentityCmd)
}
