package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_updateAccountSettingsCmd = &cobra.Command{
	Use:   "update-account-settings",
	Short: "Update the OpenSearch Serverless settings for the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_updateAccountSettingsCmd).Standalone()

	opensearchserverless_updateAccountSettingsCmd.Flags().String("capacity-limits", "", "")
	opensearchserverlessCmd.AddCommand(opensearchserverless_updateAccountSettingsCmd)
}
