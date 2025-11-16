package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aiops_deleteInvestigationGroupPolicyCmd = &cobra.Command{
	Use:   "delete-investigation-group-policy",
	Short: "Removes the IAM resource policy from being associated with the investigation group that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aiops_deleteInvestigationGroupPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(aiops_deleteInvestigationGroupPolicyCmd).Standalone()

		aiops_deleteInvestigationGroupPolicyCmd.Flags().String("identifier", "", "Specify either the name or the ARN of the investigation group that you want to remove the policy from.")
		aiops_deleteInvestigationGroupPolicyCmd.MarkFlagRequired("identifier")
	})
	aiopsCmd.AddCommand(aiops_deleteInvestigationGroupPolicyCmd)
}
