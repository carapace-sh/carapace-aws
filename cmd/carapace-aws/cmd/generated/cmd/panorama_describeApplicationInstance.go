package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_describeApplicationInstanceCmd = &cobra.Command{
	Use:   "describe-application-instance",
	Short: "Returns information about an application instance on a device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_describeApplicationInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_describeApplicationInstanceCmd).Standalone()

		panorama_describeApplicationInstanceCmd.Flags().String("application-instance-id", "", "The application instance's ID.")
		panorama_describeApplicationInstanceCmd.MarkFlagRequired("application-instance-id")
	})
	panoramaCmd.AddCommand(panorama_describeApplicationInstanceCmd)
}
