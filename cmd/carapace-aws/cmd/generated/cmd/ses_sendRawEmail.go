package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_sendRawEmailCmd = &cobra.Command{
	Use:   "send-raw-email",
	Short: "Composes an email message and immediately queues it for sending.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_sendRawEmailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_sendRawEmailCmd).Standalone()

		ses_sendRawEmailCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to use when you send an email using `SendRawEmail`.")
		ses_sendRawEmailCmd.Flags().String("destinations", "", "A list of destinations for the message, consisting of To:, CC:, and BCC: addresses.")
		ses_sendRawEmailCmd.Flags().String("from-arn", "", "This parameter is used only for sending authorization.")
		ses_sendRawEmailCmd.Flags().String("raw-message", "", "The raw email message itself.")
		ses_sendRawEmailCmd.Flags().String("return-path-arn", "", "This parameter is used only for sending authorization.")
		ses_sendRawEmailCmd.Flags().String("source", "", "The identity's email address.")
		ses_sendRawEmailCmd.Flags().String("source-arn", "", "This parameter is used only for sending authorization.")
		ses_sendRawEmailCmd.Flags().String("tags", "", "A list of tags, in the form of name/value pairs, to apply to an email that you send using `SendRawEmail`.")
		ses_sendRawEmailCmd.MarkFlagRequired("raw-message")
	})
	sesCmd.AddCommand(ses_sendRawEmailCmd)
}
