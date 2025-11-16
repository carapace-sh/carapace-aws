package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aiops_getInvestigationGroupPolicyCmd = &cobra.Command{
	Use:   "get-investigation-group-policy",
	Short: "Returns the JSON of the IAM resource policy associated with the specified investigation group in a string.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aiops_getInvestigationGroupPolicyCmd).Standalone()

	aiops_getInvestigationGroupPolicyCmd.Flags().String("identifier", "", "Specify either the name or the ARN of the investigation group that you want to view the policy of.")
	aiops_getInvestigationGroupPolicyCmd.MarkFlagRequired("identifier")
	aiopsCmd.AddCommand(aiops_getInvestigationGroupPolicyCmd)
}
