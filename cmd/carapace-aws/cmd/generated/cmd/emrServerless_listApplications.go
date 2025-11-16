package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Lists applications based on a set of parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_listApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrServerless_listApplicationsCmd).Standalone()

		emrServerless_listApplicationsCmd.Flags().String("max-results", "", "The maximum number of applications that can be listed.")
		emrServerless_listApplicationsCmd.Flags().String("next-token", "", "The token for the next set of application results.")
		emrServerless_listApplicationsCmd.Flags().String("states", "", "An optional filter for application states.")
	})
	emrServerlessCmd.AddCommand(emrServerless_listApplicationsCmd)
}
