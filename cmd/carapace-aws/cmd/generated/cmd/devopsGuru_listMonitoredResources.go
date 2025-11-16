package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_listMonitoredResourcesCmd = &cobra.Command{
	Use:   "list-monitored-resources",
	Short: "Returns the list of all log groups that are being monitored and tagged by DevOps Guru.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_listMonitoredResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_listMonitoredResourcesCmd).Standalone()

		devopsGuru_listMonitoredResourcesCmd.Flags().String("filters", "", "Filters to determine which monitored resources you want to retrieve.")
		devopsGuru_listMonitoredResourcesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		devopsGuru_listMonitoredResourcesCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	})
	devopsGuruCmd.AddCommand(devopsGuru_listMonitoredResourcesCmd)
}
