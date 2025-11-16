package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_listEnabledBaselinesCmd = &cobra.Command{
	Use:   "list-enabled-baselines",
	Short: "Returns a list of summaries describing `EnabledBaseline` resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_listEnabledBaselinesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_listEnabledBaselinesCmd).Standalone()

		controltower_listEnabledBaselinesCmd.Flags().String("filter", "", "A filter applied on the `ListEnabledBaseline` operation.")
		controltower_listEnabledBaselinesCmd.Flags().Bool("include-children", false, "A value that can be set to include the child enabled baselines in responses.")
		controltower_listEnabledBaselinesCmd.Flags().String("max-results", "", "The maximum number of results to be shown.")
		controltower_listEnabledBaselinesCmd.Flags().String("next-token", "", "A pagination token.")
		controltower_listEnabledBaselinesCmd.Flags().Bool("no-include-children", false, "A value that can be set to include the child enabled baselines in responses.")
		controltower_listEnabledBaselinesCmd.Flag("no-include-children").Hidden = true
	})
	controltowerCmd.AddCommand(controltower_listEnabledBaselinesCmd)
}
