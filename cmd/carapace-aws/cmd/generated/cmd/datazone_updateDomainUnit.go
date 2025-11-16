package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateDomainUnitCmd = &cobra.Command{
	Use:   "update-domain-unit",
	Short: "Updates the domain unit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateDomainUnitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_updateDomainUnitCmd).Standalone()

		datazone_updateDomainUnitCmd.Flags().String("description", "", "The description of the domain unit that you want to update.")
		datazone_updateDomainUnitCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to update a domain unit.")
		datazone_updateDomainUnitCmd.Flags().String("identifier", "", "The ID of the domain unit that you want to update.")
		datazone_updateDomainUnitCmd.Flags().String("name", "", "The name of the domain unit that you want to update.")
		datazone_updateDomainUnitCmd.MarkFlagRequired("domain-identifier")
		datazone_updateDomainUnitCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_updateDomainUnitCmd)
}
