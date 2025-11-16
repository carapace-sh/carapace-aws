package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listLicenseEndpointsCmd = &cobra.Command{
	Use:   "list-license-endpoints",
	Short: "Lists license endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listLicenseEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listLicenseEndpointsCmd).Standalone()

		deadline_listLicenseEndpointsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listLicenseEndpointsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	})
	deadlineCmd.AddCommand(deadline_listLicenseEndpointsCmd)
}
