package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_startPolicyGenerationCmd = &cobra.Command{
	Use:   "start-policy-generation",
	Short: "Starts the policy generation request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_startPolicyGenerationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_startPolicyGenerationCmd).Standalone()

		accessanalyzer_startPolicyGenerationCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		accessanalyzer_startPolicyGenerationCmd.Flags().String("cloud-trail-details", "", "A `CloudTrailDetails` object that contains details about a `Trail` that you want to analyze to generate policies.")
		accessanalyzer_startPolicyGenerationCmd.Flags().String("policy-generation-details", "", "Contains the ARN of the IAM entity (user or role) for which you are generating a policy.")
		accessanalyzer_startPolicyGenerationCmd.MarkFlagRequired("policy-generation-details")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_startPolicyGenerationCmd)
}
