package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_listApplicationInstanceDependenciesCmd = &cobra.Command{
	Use:   "list-application-instance-dependencies",
	Short: "Returns a list of application instance dependencies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_listApplicationInstanceDependenciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_listApplicationInstanceDependenciesCmd).Standalone()

		panorama_listApplicationInstanceDependenciesCmd.Flags().String("application-instance-id", "", "The application instance's ID.")
		panorama_listApplicationInstanceDependenciesCmd.Flags().String("max-results", "", "The maximum number of application instance dependencies to return in one page of results.")
		panorama_listApplicationInstanceDependenciesCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		panorama_listApplicationInstanceDependenciesCmd.MarkFlagRequired("application-instance-id")
	})
	panoramaCmd.AddCommand(panorama_listApplicationInstanceDependenciesCmd)
}
