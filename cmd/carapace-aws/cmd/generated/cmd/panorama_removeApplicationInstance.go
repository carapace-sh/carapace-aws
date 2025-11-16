package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_removeApplicationInstanceCmd = &cobra.Command{
	Use:   "remove-application-instance",
	Short: "Removes an application instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_removeApplicationInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_removeApplicationInstanceCmd).Standalone()

		panorama_removeApplicationInstanceCmd.Flags().String("application-instance-id", "", "An application instance ID.")
		panorama_removeApplicationInstanceCmd.MarkFlagRequired("application-instance-id")
	})
	panoramaCmd.AddCommand(panorama_removeApplicationInstanceCmd)
}
