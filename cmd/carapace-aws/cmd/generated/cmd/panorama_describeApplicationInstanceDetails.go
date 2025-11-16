package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_describeApplicationInstanceDetailsCmd = &cobra.Command{
	Use:   "describe-application-instance-details",
	Short: "Returns information about an application instance's configuration manifest.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_describeApplicationInstanceDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_describeApplicationInstanceDetailsCmd).Standalone()

		panorama_describeApplicationInstanceDetailsCmd.Flags().String("application-instance-id", "", "The application instance's ID.")
		panorama_describeApplicationInstanceDetailsCmd.MarkFlagRequired("application-instance-id")
	})
	panoramaCmd.AddCommand(panorama_describeApplicationInstanceDetailsCmd)
}
