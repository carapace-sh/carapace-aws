package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pricing_getAttributeValuesCmd = &cobra.Command{
	Use:   "get-attribute-values",
	Short: "Returns a list of attribute values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pricing_getAttributeValuesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pricing_getAttributeValuesCmd).Standalone()

		pricing_getAttributeValuesCmd.Flags().String("attribute-name", "", "The name of the attribute that you want to retrieve the values for, such as `volumeType`.")
		pricing_getAttributeValuesCmd.Flags().String("max-results", "", "The maximum number of results to return in response.")
		pricing_getAttributeValuesCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results that you want to retrieve.")
		pricing_getAttributeValuesCmd.Flags().String("service-code", "", "The service code for the service whose attributes you want to retrieve.")
		pricing_getAttributeValuesCmd.MarkFlagRequired("attribute-name")
		pricing_getAttributeValuesCmd.MarkFlagRequired("service-code")
	})
	pricingCmd.AddCommand(pricing_getAttributeValuesCmd)
}
