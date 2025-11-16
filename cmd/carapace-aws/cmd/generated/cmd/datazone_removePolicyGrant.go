package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_removePolicyGrantCmd = &cobra.Command{
	Use:   "remove-policy-grant",
	Short: "Removes a policy grant.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_removePolicyGrantCmd).Standalone()

	datazone_removePolicyGrantCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
	datazone_removePolicyGrantCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to remove a policy grant.")
	datazone_removePolicyGrantCmd.Flags().String("entity-identifier", "", "The ID of the entity from which you want to remove a policy grant.")
	datazone_removePolicyGrantCmd.Flags().String("entity-type", "", "The type of the entity from which you want to remove a policy grant.")
	datazone_removePolicyGrantCmd.Flags().String("grant-identifier", "", "The ID of the policy grant that is to be removed from a specified entity.")
	datazone_removePolicyGrantCmd.Flags().String("policy-type", "", "The type of the policy that you want to remove.")
	datazone_removePolicyGrantCmd.Flags().String("principal", "", "The principal from which you want to remove a policy grant.")
	datazone_removePolicyGrantCmd.MarkFlagRequired("domain-identifier")
	datazone_removePolicyGrantCmd.MarkFlagRequired("entity-identifier")
	datazone_removePolicyGrantCmd.MarkFlagRequired("entity-type")
	datazone_removePolicyGrantCmd.MarkFlagRequired("policy-type")
	datazone_removePolicyGrantCmd.MarkFlagRequired("principal")
	datazoneCmd.AddCommand(datazone_removePolicyGrantCmd)
}
