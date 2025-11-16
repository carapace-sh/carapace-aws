package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getDomainUnitCmd = &cobra.Command{
	Use:   "get-domain-unit",
	Short: "Gets the details of the specified domain unit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getDomainUnitCmd).Standalone()

	datazone_getDomainUnitCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to get a domain unit.")
	datazone_getDomainUnitCmd.Flags().String("identifier", "", "The identifier of the domain unit that you want to get.")
	datazone_getDomainUnitCmd.MarkFlagRequired("domain-identifier")
	datazone_getDomainUnitCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getDomainUnitCmd)
}
