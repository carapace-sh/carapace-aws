package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearchserverless_getAccountSettingsCmd = &cobra.Command{
	Use:   "get-account-settings",
	Short: "Returns account-level settings related to OpenSearch Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearchserverless_getAccountSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearchserverless_getAccountSettingsCmd).Standalone()

	})
	opensearchserverlessCmd.AddCommand(opensearchserverless_getAccountSettingsCmd)
}
