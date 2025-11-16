package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listPolicyGrantsCmd = &cobra.Command{
	Use:   "list-policy-grants",
	Short: "Lists policy grants.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listPolicyGrantsCmd).Standalone()

	datazone_listPolicyGrantsCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to list policy grants.")
	datazone_listPolicyGrantsCmd.Flags().String("entity-identifier", "", "The ID of the entity for which you want to list policy grants.")
	datazone_listPolicyGrantsCmd.Flags().String("entity-type", "", "The type of entity for which you want to list policy grants.")
	datazone_listPolicyGrantsCmd.Flags().String("max-results", "", "The maximum number of grants to return in a single call to `ListPolicyGrants`.")
	datazone_listPolicyGrantsCmd.Flags().String("next-token", "", "When the number of grants is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of grants, the response includes a pagination token named `NextToken`.")
	datazone_listPolicyGrantsCmd.Flags().String("policy-type", "", "The type of policy that you want to list.")
	datazone_listPolicyGrantsCmd.MarkFlagRequired("domain-identifier")
	datazone_listPolicyGrantsCmd.MarkFlagRequired("entity-identifier")
	datazone_listPolicyGrantsCmd.MarkFlagRequired("entity-type")
	datazone_listPolicyGrantsCmd.MarkFlagRequired("policy-type")
	datazoneCmd.AddCommand(datazone_listPolicyGrantsCmd)
}
