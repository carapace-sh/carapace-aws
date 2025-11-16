package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_updateContentCmd = &cobra.Command{
	Use:   "update-content",
	Short: "Updates information about the content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_updateContentCmd).Standalone()

	qconnect_updateContentCmd.Flags().String("content-id", "", "The identifier of the content.")
	qconnect_updateContentCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_updateContentCmd.Flags().String("metadata", "", "A key/value map to store attributes without affecting tagging or recommendations.")
	qconnect_updateContentCmd.Flags().Bool("no-remove-override-link-out-uri", false, "Unset the existing `overrideLinkOutUri` if it exists.")
	qconnect_updateContentCmd.Flags().String("override-link-out-uri", "", "The URI for the article.")
	qconnect_updateContentCmd.Flags().Bool("remove-override-link-out-uri", false, "Unset the existing `overrideLinkOutUri` if it exists.")
	qconnect_updateContentCmd.Flags().String("revision-id", "", "The `revisionId` of the content resource to update, taken from an earlier call to `GetContent`, `GetContentSummary`, `SearchContent`, or `ListContents`.")
	qconnect_updateContentCmd.Flags().String("title", "", "The title of the content.")
	qconnect_updateContentCmd.Flags().String("upload-id", "", "A pointer to the uploaded asset.")
	qconnect_updateContentCmd.MarkFlagRequired("content-id")
	qconnect_updateContentCmd.MarkFlagRequired("knowledge-base-id")
	qconnect_updateContentCmd.Flag("no-remove-override-link-out-uri").Hidden = true
	qconnectCmd.AddCommand(qconnect_updateContentCmd)
}
