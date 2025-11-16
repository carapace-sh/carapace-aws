package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_listDevicesJobsCmd = &cobra.Command{
	Use:   "list-devices-jobs",
	Short: "Returns a list of jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_listDevicesJobsCmd).Standalone()

	panorama_listDevicesJobsCmd.Flags().String("device-id", "", "Filter results by the job's target device ID.")
	panorama_listDevicesJobsCmd.Flags().String("max-results", "", "The maximum number of device jobs to return in one page of results.")
	panorama_listDevicesJobsCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	panoramaCmd.AddCommand(panorama_listDevicesJobsCmd)
}
