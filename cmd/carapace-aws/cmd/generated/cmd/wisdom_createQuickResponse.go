package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_createQuickResponseCmd = &cobra.Command{
	Use:   "create-quick-response",
	Short: "Creates a Wisdom quick response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_createQuickResponseCmd).Standalone()

	wisdom_createQuickResponseCmd.Flags().String("channels", "", "The Amazon Connect channels this quick response applies to.")
	wisdom_createQuickResponseCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	wisdom_createQuickResponseCmd.Flags().String("content", "", "The content of the quick response.")
	wisdom_createQuickResponseCmd.Flags().String("content-type", "", "The media type of the quick response content.")
	wisdom_createQuickResponseCmd.Flags().String("description", "", "The description of the quick response.")
	wisdom_createQuickResponseCmd.Flags().String("grouping-configuration", "", "The configuration information of the user groups that the quick response is accessible to.")
	wisdom_createQuickResponseCmd.Flags().Bool("is-active", false, "Whether the quick response is active.")
	wisdom_createQuickResponseCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	wisdom_createQuickResponseCmd.Flags().String("language", "", "The language code value for the language in which the quick response is written.")
	wisdom_createQuickResponseCmd.Flags().String("name", "", "The name of the quick response.")
	wisdom_createQuickResponseCmd.Flags().Bool("no-is-active", false, "Whether the quick response is active.")
	wisdom_createQuickResponseCmd.Flags().String("shortcut-key", "", "The shortcut key of the quick response.")
	wisdom_createQuickResponseCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	wisdom_createQuickResponseCmd.MarkFlagRequired("content")
	wisdom_createQuickResponseCmd.MarkFlagRequired("knowledge-base-id")
	wisdom_createQuickResponseCmd.MarkFlagRequired("name")
	wisdom_createQuickResponseCmd.Flag("no-is-active").Hidden = true
	wisdomCmd.AddCommand(wisdom_createQuickResponseCmd)
}
