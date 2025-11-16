package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listDomainUnitsForParentCmd = &cobra.Command{
	Use:   "list-domain-units-for-parent",
	Short: "Lists child domain units for the specified parent domain unit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listDomainUnitsForParentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_listDomainUnitsForParentCmd).Standalone()

		datazone_listDomainUnitsForParentCmd.Flags().String("domain-identifier", "", "The ID of the domain in which you want to list domain units for a parent domain unit.")
		datazone_listDomainUnitsForParentCmd.Flags().String("max-results", "", "The maximum number of domain units to return in a single call to ListDomainUnitsForParent.")
		datazone_listDomainUnitsForParentCmd.Flags().String("next-token", "", "When the number of domain units is greater than the default value for the MaxResults parameter, or if you explicitly specify a value for MaxResults that is less than the number of domain units, the response includes a pagination token named NextToken.")
		datazone_listDomainUnitsForParentCmd.Flags().String("parent-domain-unit-identifier", "", "The ID of the parent domain unit.")
		datazone_listDomainUnitsForParentCmd.MarkFlagRequired("domain-identifier")
		datazone_listDomainUnitsForParentCmd.MarkFlagRequired("parent-domain-unit-identifier")
	})
	datazoneCmd.AddCommand(datazone_listDomainUnitsForParentCmd)
}
