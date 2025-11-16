package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_describeProgramCmd = &cobra.Command{
	Use:   "describe-program",
	Short: "Describes a program within a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_describeProgramCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_describeProgramCmd).Standalone()

		mediatailor_describeProgramCmd.Flags().String("channel-name", "", "The name of the channel associated with this Program.")
		mediatailor_describeProgramCmd.Flags().String("program-name", "", "The name of the program.")
		mediatailor_describeProgramCmd.MarkFlagRequired("channel-name")
		mediatailor_describeProgramCmd.MarkFlagRequired("program-name")
	})
	mediatailorCmd.AddCommand(mediatailor_describeProgramCmd)
}
