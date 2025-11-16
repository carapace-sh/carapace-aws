package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_attachTypedLinkCmd = &cobra.Command{
	Use:   "attach-typed-link",
	Short: "Attaches a typed link to a specified source and target object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_attachTypedLinkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_attachTypedLinkCmd).Standalone()

		clouddirectory_attachTypedLinkCmd.Flags().String("attributes", "", "A set of attributes that are associated with the typed link.")
		clouddirectory_attachTypedLinkCmd.Flags().String("directory-arn", "", "The Amazon Resource Name (ARN) of the directory where you want to attach the typed link.")
		clouddirectory_attachTypedLinkCmd.Flags().String("source-object-reference", "", "Identifies the source object that the typed link will attach to.")
		clouddirectory_attachTypedLinkCmd.Flags().String("target-object-reference", "", "Identifies the target object that the typed link will attach to.")
		clouddirectory_attachTypedLinkCmd.Flags().String("typed-link-facet", "", "Identifies the typed link facet that is associated with the typed link.")
		clouddirectory_attachTypedLinkCmd.MarkFlagRequired("attributes")
		clouddirectory_attachTypedLinkCmd.MarkFlagRequired("directory-arn")
		clouddirectory_attachTypedLinkCmd.MarkFlagRequired("source-object-reference")
		clouddirectory_attachTypedLinkCmd.MarkFlagRequired("target-object-reference")
		clouddirectory_attachTypedLinkCmd.MarkFlagRequired("typed-link-facet")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_attachTypedLinkCmd)
}
