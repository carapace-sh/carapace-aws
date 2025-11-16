package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_listPackagesCmd = &cobra.Command{
	Use:   "list-packages",
	Short: "Returns a list of packages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_listPackagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_listPackagesCmd).Standalone()

		panorama_listPackagesCmd.Flags().String("max-results", "", "The maximum number of packages to return in one page of results.")
		panorama_listPackagesCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	})
	panoramaCmd.AddCommand(panorama_listPackagesCmd)
}
