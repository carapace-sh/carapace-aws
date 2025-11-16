package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_detachTypedLinkCmd = &cobra.Command{
	Use:   "detach-typed-link",
	Short: "Detaches a typed link from a specified source and target object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_detachTypedLinkCmd).Standalone()

	clouddirectory_detachTypedLinkCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) of the directory where you want to detach the typed link.")
	clouddirectory_detachTypedLinkCmd.Flags().String("typed-link-specifier", "", "Used to accept a typed link specifier as input.")
	clouddirectory_detachTypedLinkCmd.MarkFlagRequired("directory-arn")
	clouddirectory_detachTypedLinkCmd.MarkFlagRequired("typed-link-specifier")
	clouddirectoryCmd.AddCommand(clouddirectory_detachTypedLinkCmd)
}
