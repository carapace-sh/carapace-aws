package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listEffectivePolicyValidationErrorsCmd = &cobra.Command{
	Use:   "list-effective-policy-validation-errors",
	Short: "Lists all the validation errors on an [effective policy](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_effective.html) for a specified account and policy type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listEffectivePolicyValidationErrorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_listEffectivePolicyValidationErrorsCmd).Standalone()

		organizations_listEffectivePolicyValidationErrorsCmd.Flags().String("account-id", "", "The ID of the account that you want details about.")
		organizations_listEffectivePolicyValidationErrorsCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
		organizations_listEffectivePolicyValidationErrorsCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
		organizations_listEffectivePolicyValidationErrorsCmd.Flags().String("policy-type", "", "The type of policy that you want information about.")
		organizations_listEffectivePolicyValidationErrorsCmd.MarkFlagRequired("account-id")
		organizations_listEffectivePolicyValidationErrorsCmd.MarkFlagRequired("policy-type")
	})
	organizationsCmd.AddCommand(organizations_listEffectivePolicyValidationErrorsCmd)
}
