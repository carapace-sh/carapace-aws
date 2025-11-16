package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_getApplicationCmd = &cobra.Command{
	Use:   "get-application",
	Short: "Gets the specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_getApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serverlessrepo_getApplicationCmd).Standalone()

		serverlessrepo_getApplicationCmd.Flags().String("application-id", "", "The Amazon Resource Name (ARN) of the application.")
		serverlessrepo_getApplicationCmd.Flags().String("semantic-version", "", "The semantic version of the application to get.")
		serverlessrepo_getApplicationCmd.MarkFlagRequired("application-id")
	})
	serverlessrepoCmd.AddCommand(serverlessrepo_getApplicationCmd)
}
