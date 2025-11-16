package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_updateServiceAccessPoliciesCmd = &cobra.Command{
	Use:   "update-service-access-policies",
	Short: "Configures the access rules that control access to the domain's document and search endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_updateServiceAccessPoliciesCmd).Standalone()

	cloudsearch_updateServiceAccessPoliciesCmd.Flags().String("access-policies", "", "The access rules you want to configure.")
	cloudsearch_updateServiceAccessPoliciesCmd.Flags().String("domain-name", "", "")
	cloudsearch_updateServiceAccessPoliciesCmd.MarkFlagRequired("access-policies")
	cloudsearch_updateServiceAccessPoliciesCmd.MarkFlagRequired("domain-name")
	cloudsearchCmd.AddCommand(cloudsearch_updateServiceAccessPoliciesCmd)
}
