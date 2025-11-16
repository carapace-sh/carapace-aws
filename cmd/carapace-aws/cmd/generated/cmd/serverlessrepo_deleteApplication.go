package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes the specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_deleteApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serverlessrepo_deleteApplicationCmd).Standalone()

		serverlessrepo_deleteApplicationCmd.Flags().String("application-id", "", "The Amazon Resource Name (ARN) of the application.")
		serverlessrepo_deleteApplicationCmd.MarkFlagRequired("application-id")
	})
	serverlessrepoCmd.AddCommand(serverlessrepo_deleteApplicationCmd)
}
