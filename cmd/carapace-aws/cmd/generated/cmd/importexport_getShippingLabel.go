package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importexport_getShippingLabelCmd = &cobra.Command{
	Use:   "get-shipping-label",
	Short: "This operation generates a pre-paid UPS shipping label that you will use to ship your device to AWS for processing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importexport_getShippingLabelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(importexport_getShippingLabelCmd).Standalone()

		importexport_getShippingLabelCmd.Flags().String("apiversion", "", "")
		importexport_getShippingLabelCmd.Flags().String("city", "", "")
		importexport_getShippingLabelCmd.Flags().String("company", "", "")
		importexport_getShippingLabelCmd.Flags().String("country", "", "")
		importexport_getShippingLabelCmd.Flags().String("job-ids", "", "")
		importexport_getShippingLabelCmd.Flags().String("name", "", "")
		importexport_getShippingLabelCmd.Flags().String("phone-number", "", "")
		importexport_getShippingLabelCmd.Flags().String("postal-code", "", "")
		importexport_getShippingLabelCmd.Flags().String("state-or-province", "", "")
		importexport_getShippingLabelCmd.Flags().String("street1", "", "")
		importexport_getShippingLabelCmd.Flags().String("street2", "", "")
		importexport_getShippingLabelCmd.Flags().String("street3", "", "")
		importexport_getShippingLabelCmd.MarkFlagRequired("job-ids")
	})
	importexportCmd.AddCommand(importexport_getShippingLabelCmd)
}
