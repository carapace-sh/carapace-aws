package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_confirmCustomerAgreementCmd = &cobra.Command{
	Use:   "confirm-customer-agreement",
	Short: "The confirmation of the terms of agreement when creating the connection/link aggregation group (LAG).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_confirmCustomerAgreementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_confirmCustomerAgreementCmd).Standalone()

		directconnect_confirmCustomerAgreementCmd.Flags().String("agreement-name", "", "The name of the customer agreement.")
	})
	directconnectCmd.AddCommand(directconnect_confirmCustomerAgreementCmd)
}
