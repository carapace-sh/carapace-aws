package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_validatePolicyCmd = &cobra.Command{
	Use:   "validate-policy",
	Short: "Requests the validation of a policy and returns a list of findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_validatePolicyCmd).Standalone()

	accessanalyzer_validatePolicyCmd.Flags().String("locale", "", "The locale to use for localizing the findings.")
	accessanalyzer_validatePolicyCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	accessanalyzer_validatePolicyCmd.Flags().String("next-token", "", "A token used for pagination of results returned.")
	accessanalyzer_validatePolicyCmd.Flags().String("policy-document", "", "The JSON policy document to use as the content for the policy.")
	accessanalyzer_validatePolicyCmd.Flags().String("policy-type", "", "The type of policy to validate.")
	accessanalyzer_validatePolicyCmd.Flags().String("validate-policy-resource-type", "", "The type of resource to attach to your resource policy.")
	accessanalyzer_validatePolicyCmd.MarkFlagRequired("policy-document")
	accessanalyzer_validatePolicyCmd.MarkFlagRequired("policy-type")
	accessanalyzerCmd.AddCommand(accessanalyzer_validatePolicyCmd)
}
