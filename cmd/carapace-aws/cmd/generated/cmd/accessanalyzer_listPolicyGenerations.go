package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_listPolicyGenerationsCmd = &cobra.Command{
	Use:   "list-policy-generations",
	Short: "Lists all of the policy generations requested in the last seven days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_listPolicyGenerationsCmd).Standalone()

	accessanalyzer_listPolicyGenerationsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	accessanalyzer_listPolicyGenerationsCmd.Flags().String("next-token", "", "A token used for pagination of results returned.")
	accessanalyzer_listPolicyGenerationsCmd.Flags().String("principal-arn", "", "The ARN of the IAM entity (user or role) for which you are generating a policy.")
	accessanalyzerCmd.AddCommand(accessanalyzer_listPolicyGenerationsCmd)
}
