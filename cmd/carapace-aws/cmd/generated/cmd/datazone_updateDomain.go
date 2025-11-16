package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateDomainCmd = &cobra.Command{
	Use:   "update-domain",
	Short: "Updates a Amazon DataZone domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_updateDomainCmd).Standalone()

		datazone_updateDomainCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_updateDomainCmd.Flags().String("description", "", "The description to be updated as part of the `UpdateDomain` action.")
		datazone_updateDomainCmd.Flags().String("domain-execution-role", "", "The domain execution role to be updated as part of the `UpdateDomain` action.")
		datazone_updateDomainCmd.Flags().String("identifier", "", "The ID of the Amazon Web Services domain that is to be updated.")
		datazone_updateDomainCmd.Flags().String("name", "", "The name to be updated as part of the `UpdateDomain` action.")
		datazone_updateDomainCmd.Flags().String("service-role", "", "The service role of the domain.")
		datazone_updateDomainCmd.Flags().String("single-sign-on", "", "The single sign-on option to be updated as part of the `UpdateDomain` action.")
		datazone_updateDomainCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_updateDomainCmd)
}
