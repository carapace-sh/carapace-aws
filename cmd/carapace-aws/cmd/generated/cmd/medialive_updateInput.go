package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateInputCmd = &cobra.Command{
	Use:   "update-input",
	Short: "Updates an input.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateInputCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_updateInputCmd).Standalone()

		medialive_updateInputCmd.Flags().String("destinations", "", "Destination settings for PUSH type inputs.")
		medialive_updateInputCmd.Flags().String("input-devices", "", "Settings for the devices.")
		medialive_updateInputCmd.Flags().String("input-id", "", "Unique ID of the input.")
		medialive_updateInputCmd.Flags().String("input-security-groups", "", "A list of security groups referenced by IDs to attach to the input.")
		medialive_updateInputCmd.Flags().String("media-connect-flows", "", "A list of the MediaConnect Flow ARNs that you want to use as the source of the input.")
		medialive_updateInputCmd.Flags().String("multicast-settings", "", "Multicast Input settings.")
		medialive_updateInputCmd.Flags().String("name", "", "Name of the input.")
		medialive_updateInputCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the role this input assumes during and after creation.")
		medialive_updateInputCmd.Flags().String("sdi-sources", "", "")
		medialive_updateInputCmd.Flags().String("smpte2110-receiver-group-settings", "", "Include this parameter if the input is a SMPTE 2110 input, to identify the stream sources for this input.")
		medialive_updateInputCmd.Flags().String("sources", "", "The source URLs for a PULL-type input.")
		medialive_updateInputCmd.Flags().String("srt-settings", "", "The settings associated with an SRT input.")
		medialive_updateInputCmd.MarkFlagRequired("input-id")
	})
	medialiveCmd.AddCommand(medialive_updateInputCmd)
}
