package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_listApplicationInstancesCmd = &cobra.Command{
	Use:   "list-application-instances",
	Short: "Returns a list of application instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_listApplicationInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_listApplicationInstancesCmd).Standalone()

		panorama_listApplicationInstancesCmd.Flags().String("device-id", "", "The application instances' device ID.")
		panorama_listApplicationInstancesCmd.Flags().String("max-results", "", "The maximum number of application instances to return in one page of results.")
		panorama_listApplicationInstancesCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		panorama_listApplicationInstancesCmd.Flags().String("status-filter", "", "Only include instances with a specific status.")
	})
	panoramaCmd.AddCommand(panorama_listApplicationInstancesCmd)
}
