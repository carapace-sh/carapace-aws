package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_updateQuickResponseCmd = &cobra.Command{
	Use:   "update-quick-response",
	Short: "Updates an existing Amazon Q in Connect quick response.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_updateQuickResponseCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_updateQuickResponseCmd).Standalone()

		qconnect_updateQuickResponseCmd.Flags().String("channels", "", "The Amazon Connect contact channels this quick response applies to.")
		qconnect_updateQuickResponseCmd.Flags().String("content", "", "The updated content of the quick response.")
		qconnect_updateQuickResponseCmd.Flags().String("content-type", "", "The media type of the quick response content.")
		qconnect_updateQuickResponseCmd.Flags().String("description", "", "The updated description of the quick response.")
		qconnect_updateQuickResponseCmd.Flags().String("grouping-configuration", "", "The updated grouping configuration of the quick response.")
		qconnect_updateQuickResponseCmd.Flags().Bool("is-active", false, "Whether the quick response is active.")
		qconnect_updateQuickResponseCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_updateQuickResponseCmd.Flags().String("language", "", "The language code value for the language in which the quick response is written.")
		qconnect_updateQuickResponseCmd.Flags().String("name", "", "The name of the quick response.")
		qconnect_updateQuickResponseCmd.Flags().Bool("no-is-active", false, "Whether the quick response is active.")
		qconnect_updateQuickResponseCmd.Flags().Bool("no-remove-description", false, "Whether to remove the description from the quick response.")
		qconnect_updateQuickResponseCmd.Flags().Bool("no-remove-grouping-configuration", false, "Whether to remove the grouping configuration of the quick response.")
		qconnect_updateQuickResponseCmd.Flags().Bool("no-remove-shortcut-key", false, "Whether to remove the shortcut key of the quick response.")
		qconnect_updateQuickResponseCmd.Flags().String("quick-response-id", "", "The identifier of the quick response.")
		qconnect_updateQuickResponseCmd.Flags().Bool("remove-description", false, "Whether to remove the description from the quick response.")
		qconnect_updateQuickResponseCmd.Flags().Bool("remove-grouping-configuration", false, "Whether to remove the grouping configuration of the quick response.")
		qconnect_updateQuickResponseCmd.Flags().Bool("remove-shortcut-key", false, "Whether to remove the shortcut key of the quick response.")
		qconnect_updateQuickResponseCmd.Flags().String("shortcut-key", "", "The shortcut key of the quick response.")
		qconnect_updateQuickResponseCmd.MarkFlagRequired("knowledge-base-id")
		qconnect_updateQuickResponseCmd.Flag("no-is-active").Hidden = true
		qconnect_updateQuickResponseCmd.Flag("no-remove-description").Hidden = true
		qconnect_updateQuickResponseCmd.Flag("no-remove-grouping-configuration").Hidden = true
		qconnect_updateQuickResponseCmd.Flag("no-remove-shortcut-key").Hidden = true
		qconnect_updateQuickResponseCmd.MarkFlagRequired("quick-response-id")
	})
	qconnectCmd.AddCommand(qconnect_updateQuickResponseCmd)
}
