package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_deleteProgramCmd = &cobra.Command{
	Use:   "delete-program",
	Short: "Deletes a program within a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_deleteProgramCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_deleteProgramCmd).Standalone()

		mediatailor_deleteProgramCmd.Flags().String("channel-name", "", "The name of the channel.")
		mediatailor_deleteProgramCmd.Flags().String("program-name", "", "The name of the program.")
		mediatailor_deleteProgramCmd.MarkFlagRequired("channel-name")
		mediatailor_deleteProgramCmd.MarkFlagRequired("program-name")
	})
	mediatailorCmd.AddCommand(mediatailor_deleteProgramCmd)
}
