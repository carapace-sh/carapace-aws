package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createDomainUnitCmd = &cobra.Command{
	Use:   "create-domain-unit",
	Short: "Creates a domain unit in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createDomainUnitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createDomainUnitCmd).Standalone()

		datazone_createDomainUnitCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_createDomainUnitCmd.Flags().String("description", "", "The description of the domain unit.")
		datazone_createDomainUnitCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to crate a domain unit.")
		datazone_createDomainUnitCmd.Flags().String("name", "", "The name of the domain unit.")
		datazone_createDomainUnitCmd.Flags().String("parent-domain-unit-identifier", "", "The ID of the parent domain unit.")
		datazone_createDomainUnitCmd.MarkFlagRequired("domain-identifier")
		datazone_createDomainUnitCmd.MarkFlagRequired("name")
		datazone_createDomainUnitCmd.MarkFlagRequired("parent-domain-unit-identifier")
	})
	datazoneCmd.AddCommand(datazone_createDomainUnitCmd)
}
