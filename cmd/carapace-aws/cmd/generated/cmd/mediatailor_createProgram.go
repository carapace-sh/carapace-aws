package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_createProgramCmd = &cobra.Command{
	Use:   "create-program",
	Short: "Creates a program within a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_createProgramCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_createProgramCmd).Standalone()

		mediatailor_createProgramCmd.Flags().String("ad-breaks", "", "The ad break configuration settings.")
		mediatailor_createProgramCmd.Flags().String("audience-media", "", "The list of AudienceMedia defined in program.")
		mediatailor_createProgramCmd.Flags().String("channel-name", "", "The name of the channel for this Program.")
		mediatailor_createProgramCmd.Flags().String("live-source-name", "", "The name of the LiveSource for this Program.")
		mediatailor_createProgramCmd.Flags().String("program-name", "", "The name of the Program.")
		mediatailor_createProgramCmd.Flags().String("schedule-configuration", "", "The schedule configuration settings.")
		mediatailor_createProgramCmd.Flags().String("source-location-name", "", "The name of the source location.")
		mediatailor_createProgramCmd.Flags().String("vod-source-name", "", "The name that's used to refer to a VOD source.")
		mediatailor_createProgramCmd.MarkFlagRequired("channel-name")
		mediatailor_createProgramCmd.MarkFlagRequired("program-name")
		mediatailor_createProgramCmd.MarkFlagRequired("schedule-configuration")
		mediatailor_createProgramCmd.MarkFlagRequired("source-location-name")
	})
	mediatailorCmd.AddCommand(mediatailor_createProgramCmd)
}
