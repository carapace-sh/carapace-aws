package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSyncCmd = &cobra.Command{
	Use:   "cognito-sync",
	Short: "Amazon Cognito Sync\n\nAmazon Cognito Sync provides an AWS service and client library that enable cross-device syncing of application-related user data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSyncCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoSyncCmd).Standalone()

	})
	rootCmd.AddCommand(cognitoSyncCmd)
}
