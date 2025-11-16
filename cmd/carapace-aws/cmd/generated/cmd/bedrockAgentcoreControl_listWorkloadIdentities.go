package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_listWorkloadIdentitiesCmd = &cobra.Command{
	Use:   "list-workload-identities",
	Short: "Lists all workload identities in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_listWorkloadIdentitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_listWorkloadIdentitiesCmd).Standalone()

		bedrockAgentcoreControl_listWorkloadIdentitiesCmd.Flags().String("max-results", "", "Maximum number of results to return.")
		bedrockAgentcoreControl_listWorkloadIdentitiesCmd.Flags().String("next-token", "", "Pagination token.")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_listWorkloadIdentitiesCmd)
}
