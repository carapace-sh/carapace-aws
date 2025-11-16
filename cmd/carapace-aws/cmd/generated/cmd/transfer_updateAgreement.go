package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_updateAgreementCmd = &cobra.Command{
	Use:   "update-agreement",
	Short: "Updates some of the parameters for an existing agreement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_updateAgreementCmd).Standalone()

	transfer_updateAgreementCmd.Flags().String("access-role", "", "Connectors are used to send files using either the AS2 or SFTP protocol.")
	transfer_updateAgreementCmd.Flags().String("agreement-id", "", "A unique identifier for the agreement.")
	transfer_updateAgreementCmd.Flags().String("base-directory", "", "To change the landing directory (folder) for files that are transferred, provide the bucket folder that you want to use; for example, `/amzn-s3-demo-bucket/home/mydirectory`.")
	transfer_updateAgreementCmd.Flags().String("custom-directories", "", "A `CustomDirectoriesType` structure.")
	transfer_updateAgreementCmd.Flags().String("description", "", "To replace the existing description, provide a short description for the agreement.")
	transfer_updateAgreementCmd.Flags().String("enforce-message-signing", "", "Determines whether or not unsigned messages from your trading partners will be accepted.")
	transfer_updateAgreementCmd.Flags().String("local-profile-id", "", "A unique identifier for the AS2 local profile.")
	transfer_updateAgreementCmd.Flags().String("partner-profile-id", "", "A unique identifier for the partner profile.")
	transfer_updateAgreementCmd.Flags().String("preserve-filename", "", "Determines whether or not Transfer Family appends a unique string of characters to the end of the AS2 message payload filename when saving it.")
	transfer_updateAgreementCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server instance.")
	transfer_updateAgreementCmd.Flags().String("status", "", "You can update the status for the agreement, either activating an inactive agreement or the reverse.")
	transfer_updateAgreementCmd.MarkFlagRequired("agreement-id")
	transfer_updateAgreementCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_updateAgreementCmd)
}
