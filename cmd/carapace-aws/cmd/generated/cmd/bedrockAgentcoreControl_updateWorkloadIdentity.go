package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_updateWorkloadIdentityCmd = &cobra.Command{
	Use:   "update-workload-identity",
	Short: "Updates an existing workload identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_updateWorkloadIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_updateWorkloadIdentityCmd).Standalone()

		bedrockAgentcoreControl_updateWorkloadIdentityCmd.Flags().String("allowed-resource-oauth2-return-urls", "", "The new list of allowed OAuth2 return URLs for resources associated with this workload identity.")
		bedrockAgentcoreControl_updateWorkloadIdentityCmd.Flags().String("name", "", "The name of the workload identity to update.")
		bedrockAgentcoreControl_updateWorkloadIdentityCmd.MarkFlagRequired("name")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_updateWorkloadIdentityCmd)
}
