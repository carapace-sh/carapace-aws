package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controltower_listEnabledControlsCmd = &cobra.Command{
	Use:   "list-enabled-controls",
	Short: "Lists the controls enabled by Amazon Web Services Control Tower on the specified organizational unit and the accounts it contains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controltower_listEnabledControlsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(controltower_listEnabledControlsCmd).Standalone()

		controltower_listEnabledControlsCmd.Flags().String("filter", "", "An input filter for the `ListEnabledControls` API that lets you select the types of control operations to view.")
		controltower_listEnabledControlsCmd.Flags().Bool("include-children", false, "A boolean value that determines whether to include enabled controls from child organizational units in the response.")
		controltower_listEnabledControlsCmd.Flags().String("max-results", "", "How many results to return per API call.")
		controltower_listEnabledControlsCmd.Flags().String("next-token", "", "The token to continue the list from a previous API call with the same parameters.")
		controltower_listEnabledControlsCmd.Flags().Bool("no-include-children", false, "A boolean value that determines whether to include enabled controls from child organizational units in the response.")
		controltower_listEnabledControlsCmd.Flags().String("target-identifier", "", "The ARN of the organizational unit.")
		controltower_listEnabledControlsCmd.Flag("no-include-children").Hidden = true
	})
	controltowerCmd.AddCommand(controltower_listEnabledControlsCmd)
}
