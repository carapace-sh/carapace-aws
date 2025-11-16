package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_listBaselinesCmd = &cobra.Command{
	Use:   "list-baselines",
	Short: "Returns a summary list of all available baselines.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_listBaselinesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_listBaselinesCmd).Standalone()

		controltower_listBaselinesCmd.Flags().String("max-results", "", "The maximum number of results to be shown.")
		controltower_listBaselinesCmd.Flags().String("next-token", "", "A pagination token.")
	})
	controltowerCmd.AddCommand(controltower_listBaselinesCmd)
}
