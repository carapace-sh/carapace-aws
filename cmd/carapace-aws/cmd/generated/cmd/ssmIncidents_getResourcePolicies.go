package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmIncidents_getResourcePoliciesCmd = &cobra.Command{
	Use:   "get-resource-policies",
	Short: "Retrieves the resource policies attached to the specified response plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmIncidents_getResourcePoliciesCmd).Standalone()

	ssmIncidents_getResourcePoliciesCmd.Flags().String("max-results", "", "The maximum number of resource policies to display for each page of results.")
	ssmIncidents_getResourcePoliciesCmd.Flags().String("next-token", "", "The pagination token for the next set of items to return.")
	ssmIncidents_getResourcePoliciesCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the response plan with the attached resource policy.")
	ssmIncidents_getResourcePoliciesCmd.MarkFlagRequired("resource-arn")
	ssmIncidentsCmd.AddCommand(ssmIncidents_getResourcePoliciesCmd)
}
