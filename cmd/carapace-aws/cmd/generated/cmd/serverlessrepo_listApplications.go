package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverlessrepo_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Lists applications owned by the requester.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverlessrepo_listApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serverlessrepo_listApplicationsCmd).Standalone()

		serverlessrepo_listApplicationsCmd.Flags().String("max-items", "", "The total number of items to return.")
		serverlessrepo_listApplicationsCmd.Flags().String("next-token", "", "A token to specify where to start paginating.")
	})
	serverlessrepoCmd.AddCommand(serverlessrepo_listApplicationsCmd)
}
