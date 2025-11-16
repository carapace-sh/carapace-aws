package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_createArchiveCmd = &cobra.Command{
	Use:   "create-archive",
	Short: "Creates a new email archive resource for storing and retaining emails.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_createArchiveCmd).Standalone()

	mailmanager_createArchiveCmd.Flags().String("archive-name", "", "A unique name for the new archive.")
	mailmanager_createArchiveCmd.Flags().String("client-token", "", "A unique token Amazon SES uses to recognize retries of this request.")
	mailmanager_createArchiveCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key for encrypting emails in the archive.")
	mailmanager_createArchiveCmd.Flags().String("retention", "", "The period for retaining emails in the archive before automatic deletion.")
	mailmanager_createArchiveCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for the resource.")
	mailmanager_createArchiveCmd.MarkFlagRequired("archive-name")
	mailmanagerCmd.AddCommand(mailmanager_createArchiveCmd)
}
