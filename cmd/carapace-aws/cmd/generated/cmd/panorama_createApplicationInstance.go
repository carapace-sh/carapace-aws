package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_createApplicationInstanceCmd = &cobra.Command{
	Use:   "create-application-instance",
	Short: "Creates an application instance and deploys it to a device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_createApplicationInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_createApplicationInstanceCmd).Standalone()

		panorama_createApplicationInstanceCmd.Flags().String("application-instance-id-to-replace", "", "The ID of an application instance to replace with the new instance.")
		panorama_createApplicationInstanceCmd.Flags().String("default-runtime-context-device", "", "A device's ID.")
		panorama_createApplicationInstanceCmd.Flags().String("description", "", "A description for the application instance.")
		panorama_createApplicationInstanceCmd.Flags().String("manifest-overrides-payload", "", "Setting overrides for the application manifest.")
		panorama_createApplicationInstanceCmd.Flags().String("manifest-payload", "", "The application's manifest document.")
		panorama_createApplicationInstanceCmd.Flags().String("name", "", "A name for the application instance.")
		panorama_createApplicationInstanceCmd.Flags().String("runtime-role-arn", "", "The ARN of a runtime role for the application instance.")
		panorama_createApplicationInstanceCmd.Flags().String("tags", "", "Tags for the application instance.")
		panorama_createApplicationInstanceCmd.MarkFlagRequired("default-runtime-context-device")
		panorama_createApplicationInstanceCmd.MarkFlagRequired("manifest-payload")
	})
	panoramaCmd.AddCommand(panorama_createApplicationInstanceCmd)
}
