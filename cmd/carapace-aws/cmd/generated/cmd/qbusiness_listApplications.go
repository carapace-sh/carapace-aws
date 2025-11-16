package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Lists Amazon Q Business applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listApplicationsCmd).Standalone()

	qbusiness_listApplicationsCmd.Flags().String("max-results", "", "The maximum number of Amazon Q Business applications to return.")
	qbusiness_listApplicationsCmd.Flags().String("next-token", "", "If the `maxResults` response was incomplete because there is more data to retrieve, Amazon Q Business returns a pagination token in the response.")
	qbusinessCmd.AddCommand(qbusiness_listApplicationsCmd)
}
