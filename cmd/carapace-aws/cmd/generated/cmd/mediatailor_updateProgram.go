package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_updateProgramCmd = &cobra.Command{
	Use:   "update-program",
	Short: "Updates a program within a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_updateProgramCmd).Standalone()

	mediatailor_updateProgramCmd.Flags().String("ad-breaks", "", "The ad break configuration settings.")
	mediatailor_updateProgramCmd.Flags().String("audience-media", "", "The list of AudienceMedia defined in program.")
	mediatailor_updateProgramCmd.Flags().String("channel-name", "", "The name of the channel for this Program.")
	mediatailor_updateProgramCmd.Flags().String("program-name", "", "The name of the Program.")
	mediatailor_updateProgramCmd.Flags().String("schedule-configuration", "", "The schedule configuration settings.")
	mediatailor_updateProgramCmd.MarkFlagRequired("channel-name")
	mediatailor_updateProgramCmd.MarkFlagRequired("program-name")
	mediatailor_updateProgramCmd.MarkFlagRequired("schedule-configuration")
	mediatailorCmd.AddCommand(mediatailor_updateProgramCmd)
}
