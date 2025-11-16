package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Retrieves all applications or multiple applications by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_listApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_listApplicationsCmd).Standalone()

		mgn_listApplicationsCmd.Flags().String("account-id", "", "Applications list Account ID.")
		mgn_listApplicationsCmd.Flags().String("filters", "", "Applications list filters.")
		mgn_listApplicationsCmd.Flags().String("max-results", "", "Maximum results to return when listing applications.")
		mgn_listApplicationsCmd.Flags().String("next-token", "", "Request next token.")
	})
	mgnCmd.AddCommand(mgn_listApplicationsCmd)
}
