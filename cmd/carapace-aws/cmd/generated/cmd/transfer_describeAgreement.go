package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_describeAgreementCmd = &cobra.Command{
	Use:   "describe-agreement",
	Short: "Describes the agreement that's identified by the `AgreementId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_describeAgreementCmd).Standalone()

	transfer_describeAgreementCmd.Flags().String("agreement-id", "", "A unique identifier for the agreement.")
	transfer_describeAgreementCmd.Flags().String("server-id", "", "The server identifier that's associated with the agreement.")
	transfer_describeAgreementCmd.MarkFlagRequired("agreement-id")
	transfer_describeAgreementCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_describeAgreementCmd)
}
