package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var support_describeAttachmentCmd = &cobra.Command{
	Use:   "describe-attachment",
	Short: "Returns the attachment that has the specified ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(support_describeAttachmentCmd).Standalone()

	support_describeAttachmentCmd.Flags().String("attachment-id", "", "The ID of the attachment to return.")
	support_describeAttachmentCmd.MarkFlagRequired("attachment-id")
	supportCmd.AddCommand(support_describeAttachmentCmd)
}
