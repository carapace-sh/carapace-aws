package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_createAgreementCmd = &cobra.Command{
	Use:   "create-agreement",
	Short: "Creates an agreement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_createAgreementCmd).Standalone()

	transfer_createAgreementCmd.Flags().String("access-role", "", "Connectors are used to send files using either the AS2 or SFTP protocol.")
	transfer_createAgreementCmd.Flags().String("base-directory", "", "The landing directory (folder) for files transferred by using the AS2 protocol.")
	transfer_createAgreementCmd.Flags().String("custom-directories", "", "A `CustomDirectoriesType` structure.")
	transfer_createAgreementCmd.Flags().String("description", "", "A name or short description to identify the agreement.")
	transfer_createAgreementCmd.Flags().String("enforce-message-signing", "", "Determines whether or not unsigned messages from your trading partners will be accepted.")
	transfer_createAgreementCmd.Flags().String("local-profile-id", "", "A unique identifier for the AS2 local profile.")
	transfer_createAgreementCmd.Flags().String("partner-profile-id", "", "A unique identifier for the partner profile used in the agreement.")
	transfer_createAgreementCmd.Flags().String("preserve-filename", "", "Determines whether or not Transfer Family appends a unique string of characters to the end of the AS2 message payload filename when saving it.")
	transfer_createAgreementCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server instance.")
	transfer_createAgreementCmd.Flags().String("status", "", "The status of the agreement.")
	transfer_createAgreementCmd.Flags().String("tags", "", "Key-value pairs that can be used to group and search for agreements.")
	transfer_createAgreementCmd.MarkFlagRequired("access-role")
	transfer_createAgreementCmd.MarkFlagRequired("local-profile-id")
	transfer_createAgreementCmd.MarkFlagRequired("partner-profile-id")
	transfer_createAgreementCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_createAgreementCmd)
}
