package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_addPolicyGrantCmd = &cobra.Command{
	Use:   "add-policy-grant",
	Short: "Adds a policy grant (an authorization policy) to a specified entity, including domain units, environment blueprint configurations, or environment profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_addPolicyGrantCmd).Standalone()

	datazone_addPolicyGrantCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
	datazone_addPolicyGrantCmd.Flags().String("detail", "", "The details of the policy grant.")
	datazone_addPolicyGrantCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to add a policy grant.")
	datazone_addPolicyGrantCmd.Flags().String("entity-identifier", "", "The ID of the entity (resource) to which you want to add a policy grant.")
	datazone_addPolicyGrantCmd.Flags().String("entity-type", "", "The type of entity (resource) to which the grant is added.")
	datazone_addPolicyGrantCmd.Flags().String("policy-type", "", "The type of policy that you want to grant.")
	datazone_addPolicyGrantCmd.Flags().String("principal", "", "The principal to whom the permissions are granted.")
	datazone_addPolicyGrantCmd.MarkFlagRequired("detail")
	datazone_addPolicyGrantCmd.MarkFlagRequired("domain-identifier")
	datazone_addPolicyGrantCmd.MarkFlagRequired("entity-identifier")
	datazone_addPolicyGrantCmd.MarkFlagRequired("entity-type")
	datazone_addPolicyGrantCmd.MarkFlagRequired("policy-type")
	datazone_addPolicyGrantCmd.MarkFlagRequired("principal")
	datazoneCmd.AddCommand(datazone_addPolicyGrantCmd)
}
