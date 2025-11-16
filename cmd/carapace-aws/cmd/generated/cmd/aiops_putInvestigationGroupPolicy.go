package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aiops_putInvestigationGroupPolicyCmd = &cobra.Command{
	Use:   "put-investigation-group-policy",
	Short: "Creates an IAM resource policy and assigns it to the specified investigation group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aiops_putInvestigationGroupPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(aiops_putInvestigationGroupPolicyCmd).Standalone()

		aiops_putInvestigationGroupPolicyCmd.Flags().String("identifier", "", "Specify either the name or the ARN of the investigation group that you want to assign the policy to.")
		aiops_putInvestigationGroupPolicyCmd.Flags().String("policy", "", "The policy, in JSON format.")
		aiops_putInvestigationGroupPolicyCmd.MarkFlagRequired("identifier")
		aiops_putInvestigationGroupPolicyCmd.MarkFlagRequired("policy")
	})
	aiopsCmd.AddCommand(aiops_putInvestigationGroupPolicyCmd)
}
