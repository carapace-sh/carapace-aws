package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createInputCmd = &cobra.Command{
	Use:   "create-input",
	Short: "Create an input",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createInputCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_createInputCmd).Standalone()

		medialive_createInputCmd.Flags().String("destinations", "", "Destination settings for PUSH type inputs.")
		medialive_createInputCmd.Flags().String("input-devices", "", "Settings for the devices.")
		medialive_createInputCmd.Flags().String("input-network-location", "", "The location of this input.")
		medialive_createInputCmd.Flags().String("input-security-groups", "", "A list of security groups referenced by IDs to attach to the input.")
		medialive_createInputCmd.Flags().String("media-connect-flows", "", "A list of the MediaConnect Flows that you want to use in this input.")
		medialive_createInputCmd.Flags().String("multicast-settings", "", "Multicast Input settings.")
		medialive_createInputCmd.Flags().String("name", "", "Name of the input.")
		medialive_createInputCmd.Flags().String("request-id", "", "Unique identifier of the request to ensure the request is handled exactly once in case of retries.")
		medialive_createInputCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the role this input assumes during and after creation.")
		medialive_createInputCmd.Flags().String("sdi-sources", "", "")
		medialive_createInputCmd.Flags().String("smpte2110-receiver-group-settings", "", "Include this parameter if the input is a SMPTE 2110 input, to identify the stream sources for this input.")
		medialive_createInputCmd.Flags().String("sources", "", "The source URLs for a PULL-type input.")
		medialive_createInputCmd.Flags().String("srt-settings", "", "The settings associated with an SRT input.")
		medialive_createInputCmd.Flags().String("tags", "", "A collection of key-value pairs.")
		medialive_createInputCmd.Flags().String("type", "", "")
		medialive_createInputCmd.Flags().String("vpc", "", "")
	})
	medialiveCmd.AddCommand(medialive_createInputCmd)
}
