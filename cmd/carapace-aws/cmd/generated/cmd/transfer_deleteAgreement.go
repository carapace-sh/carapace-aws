package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_deleteAgreementCmd = &cobra.Command{
	Use:   "delete-agreement",
	Short: "Delete the agreement that's specified in the provided `AgreementId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_deleteAgreementCmd).Standalone()

	transfer_deleteAgreementCmd.Flags().String("agreement-id", "", "A unique identifier for the agreement.")
	transfer_deleteAgreementCmd.Flags().String("server-id", "", "The server identifier associated with the agreement that you are deleting.")
	transfer_deleteAgreementCmd.MarkFlagRequired("agreement-id")
	transfer_deleteAgreementCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_deleteAgreementCmd)
}
