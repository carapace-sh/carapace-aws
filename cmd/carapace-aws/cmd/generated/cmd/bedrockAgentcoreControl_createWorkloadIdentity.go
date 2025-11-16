package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_createWorkloadIdentityCmd = &cobra.Command{
	Use:   "create-workload-identity",
	Short: "Creates a new workload identity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_createWorkloadIdentityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_createWorkloadIdentityCmd).Standalone()

		bedrockAgentcoreControl_createWorkloadIdentityCmd.Flags().String("allowed-resource-oauth2-return-urls", "", "The list of allowed OAuth2 return URLs for resources associated with this workload identity.")
		bedrockAgentcoreControl_createWorkloadIdentityCmd.Flags().String("name", "", "The name of the workload identity.")
		bedrockAgentcoreControl_createWorkloadIdentityCmd.Flags().String("tags", "", "A map of tag keys and values to assign to the workload identity.")
		bedrockAgentcoreControl_createWorkloadIdentityCmd.MarkFlagRequired("name")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_createWorkloadIdentityCmd)
}
