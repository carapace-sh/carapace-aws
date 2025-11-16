package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteDomainUnitCmd = &cobra.Command{
	Use:   "delete-domain-unit",
	Short: "Deletes a domain unit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteDomainUnitCmd).Standalone()

	datazone_deleteDomainUnitCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to delete a domain unit.")
	datazone_deleteDomainUnitCmd.Flags().String("identifier", "", "The ID of the domain unit that you want to delete.")
	datazone_deleteDomainUnitCmd.MarkFlagRequired("domain-identifier")
	datazone_deleteDomainUnitCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_deleteDomainUnitCmd)
}
