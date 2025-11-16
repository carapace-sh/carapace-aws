package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_getWorkloadIdentityCmd = &cobra.Command{
	Use:   "get-workload-identity",
	Short: "Retrieves information about a workload identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_getWorkloadIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_getWorkloadIdentityCmd).Standalone()

		bedrockAgentcoreControl_getWorkloadIdentityCmd.Flags().String("name", "", "The name of the workload identity to retrieve.")
		bedrockAgentcoreControl_getWorkloadIdentityCmd.MarkFlagRequired("name")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_getWorkloadIdentityCmd)
}
