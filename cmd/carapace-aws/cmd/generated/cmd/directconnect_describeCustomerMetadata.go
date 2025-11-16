package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeCustomerMetadataCmd = &cobra.Command{
	Use:   "describe-customer-metadata",
	Short: "Get and view a list of customer agreements, along with their signed status and whether the customer is an NNIPartner, NNIPartnerV2, or a nonPartner.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeCustomerMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_describeCustomerMetadataCmd).Standalone()

	})
	directconnectCmd.AddCommand(directconnect_describeCustomerMetadataCmd)
}
