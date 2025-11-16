package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_describeServiceAccessPoliciesCmd = &cobra.Command{
	Use:   "describe-service-access-policies",
	Short: "Gets information about the access policies that control access to the domain's document and search endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_describeServiceAccessPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_describeServiceAccessPoliciesCmd).Standalone()

		cloudsearch_describeServiceAccessPoliciesCmd.Flags().Bool("deployed", false, "Whether to display the deployed configuration (`true`) or include any pending changes (`false`).")
		cloudsearch_describeServiceAccessPoliciesCmd.Flags().String("domain-name", "", "The name of the domain you want to describe.")
		cloudsearch_describeServiceAccessPoliciesCmd.Flags().Bool("no-deployed", false, "Whether to display the deployed configuration (`true`) or include any pending changes (`false`).")
		cloudsearch_describeServiceAccessPoliciesCmd.MarkFlagRequired("domain-name")
		cloudsearch_describeServiceAccessPoliciesCmd.Flag("no-deployed").Hidden = true
	})
	cloudsearchCmd.AddCommand(cloudsearch_describeServiceAccessPoliciesCmd)
}
