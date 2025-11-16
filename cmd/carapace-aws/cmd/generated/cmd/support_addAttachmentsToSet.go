package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_addAttachmentsToSetCmd = &cobra.Command{
	Use:   "add-attachments-to-set",
	Short: "Adds one or more attachments to an attachment set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_addAttachmentsToSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(support_addAttachmentsToSetCmd).Standalone()

		support_addAttachmentsToSetCmd.Flags().String("attachment-set-id", "", "The ID of the attachment set.")
		support_addAttachmentsToSetCmd.Flags().String("attachments", "", "One or more attachments to add to the set.")
		support_addAttachmentsToSetCmd.MarkFlagRequired("attachments")
	})
	supportCmd.AddCommand(support_addAttachmentsToSetCmd)
}
