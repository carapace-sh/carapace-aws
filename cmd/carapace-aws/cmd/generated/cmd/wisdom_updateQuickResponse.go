package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_updateQuickResponseCmd = &cobra.Command{
	Use:   "update-quick-response",
	Short: "Updates an existing Wisdom quick response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_updateQuickResponseCmd).Standalone()

	wisdom_updateQuickResponseCmd.Flags().String("channels", "", "The Amazon Connect contact channels this quick response applies to.")
	wisdom_updateQuickResponseCmd.Flags().String("content", "", "The updated content of the quick response.")
	wisdom_updateQuickResponseCmd.Flags().String("content-type", "", "The media type of the quick response content.")
	wisdom_updateQuickResponseCmd.Flags().String("description", "", "The updated description of the quick response.")
	wisdom_updateQuickResponseCmd.Flags().String("grouping-configuration", "", "The updated grouping configuration of the quick response.")
	wisdom_updateQuickResponseCmd.Flags().Bool("is-active", false, "Whether the quick response is active.")
	wisdom_updateQuickResponseCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	wisdom_updateQuickResponseCmd.Flags().String("language", "", "The language code value for the language in which the quick response is written.")
	wisdom_updateQuickResponseCmd.Flags().String("name", "", "The name of the quick response.")
	wisdom_updateQuickResponseCmd.Flags().Bool("no-is-active", false, "Whether the quick response is active.")
	wisdom_updateQuickResponseCmd.Flags().Bool("no-remove-description", false, "Whether to remove the description from the quick response.")
	wisdom_updateQuickResponseCmd.Flags().Bool("no-remove-grouping-configuration", false, "Whether to remove the grouping configuration of the quick response.")
	wisdom_updateQuickResponseCmd.Flags().Bool("no-remove-shortcut-key", false, "Whether to remove the shortcut key of the quick response.")
	wisdom_updateQuickResponseCmd.Flags().String("quick-response-id", "", "The identifier of the quick response.")
	wisdom_updateQuickResponseCmd.Flags().Bool("remove-description", false, "Whether to remove the description from the quick response.")
	wisdom_updateQuickResponseCmd.Flags().Bool("remove-grouping-configuration", false, "Whether to remove the grouping configuration of the quick response.")
	wisdom_updateQuickResponseCmd.Flags().Bool("remove-shortcut-key", false, "Whether to remove the shortcut key of the quick response.")
	wisdom_updateQuickResponseCmd.Flags().String("shortcut-key", "", "The shortcut key of the quick response.")
	wisdom_updateQuickResponseCmd.MarkFlagRequired("knowledge-base-id")
	wisdom_updateQuickResponseCmd.Flag("no-is-active").Hidden = true
	wisdom_updateQuickResponseCmd.Flag("no-remove-description").Hidden = true
	wisdom_updateQuickResponseCmd.Flag("no-remove-grouping-configuration").Hidden = true
	wisdom_updateQuickResponseCmd.Flag("no-remove-shortcut-key").Hidden = true
	wisdom_updateQuickResponseCmd.MarkFlagRequired("quick-response-id")
	wisdomCmd.AddCommand(wisdom_updateQuickResponseCmd)
}
