package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_listApplicationInstanceNodeInstancesCmd = &cobra.Command{
	Use:   "list-application-instance-node-instances",
	Short: "Returns a list of application node instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_listApplicationInstanceNodeInstancesCmd).Standalone()

	panorama_listApplicationInstanceNodeInstancesCmd.Flags().String("application-instance-id", "", "The node instances' application instance ID.")
	panorama_listApplicationInstanceNodeInstancesCmd.Flags().String("max-results", "", "The maximum number of node instances to return in one page of results.")
	panorama_listApplicationInstanceNodeInstancesCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	panorama_listApplicationInstanceNodeInstancesCmd.MarkFlagRequired("application-instance-id")
	panoramaCmd.AddCommand(panorama_listApplicationInstanceNodeInstancesCmd)
}
