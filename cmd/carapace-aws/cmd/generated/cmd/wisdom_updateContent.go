package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_updateContentCmd = &cobra.Command{
	Use:   "update-content",
	Short: "Updates information about the content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_updateContentCmd).Standalone()

	wisdom_updateContentCmd.Flags().String("content-id", "", "The identifier of the content.")
	wisdom_updateContentCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	wisdom_updateContentCmd.Flags().String("metadata", "", "A key/value map to store attributes without affecting tagging or recommendations.")
	wisdom_updateContentCmd.Flags().Bool("no-remove-override-link-out-uri", false, "Unset the existing `overrideLinkOutUri` if it exists.")
	wisdom_updateContentCmd.Flags().String("override-link-out-uri", "", "The URI for the article.")
	wisdom_updateContentCmd.Flags().Bool("remove-override-link-out-uri", false, "Unset the existing `overrideLinkOutUri` if it exists.")
	wisdom_updateContentCmd.Flags().String("revision-id", "", "The `revisionId` of the content resource to update, taken from an earlier call to `GetContent`, `GetContentSummary`, `SearchContent`, or `ListContents`.")
	wisdom_updateContentCmd.Flags().String("title", "", "The title of the content.")
	wisdom_updateContentCmd.Flags().String("upload-id", "", "A pointer to the uploaded asset.")
	wisdom_updateContentCmd.MarkFlagRequired("content-id")
	wisdom_updateContentCmd.MarkFlagRequired("knowledge-base-id")
	wisdom_updateContentCmd.Flag("no-remove-override-link-out-uri").Hidden = true
	wisdomCmd.AddCommand(wisdom_updateContentCmd)
}
