package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_listApplicationVersionsCmd = &cobra.Command{
	Use:   "list-application-versions",
	Short: "Lists versions for the specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_listApplicationVersionsCmd).Standalone()

	serverlessrepo_listApplicationVersionsCmd.Flags().String("application-id", "", "The Amazon Resource Name (ARN) of the application.")
	serverlessrepo_listApplicationVersionsCmd.Flags().String("max-items", "", "The total number of items to return.")
	serverlessrepo_listApplicationVersionsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	serverlessrepo_listApplicationVersionsCmd.MarkFlagRequired("application-id")
	serverlessrepoCmd.AddCommand(serverlessrepo_listApplicationVersionsCmd)
}
