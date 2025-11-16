package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listCommandsCmd = &cobra.Command{
	Use:   "list-commands",
	Short: "Lists the commands requested by users of the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listCommandsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_listCommandsCmd).Standalone()

		ssm_listCommandsCmd.Flags().String("command-id", "", "(Optional) If provided, lists only the specified command.")
		ssm_listCommandsCmd.Flags().String("filters", "", "(Optional) One or more filters.")
		ssm_listCommandsCmd.Flags().String("instance-id", "", "(Optional) Lists commands issued against this managed node ID.")
		ssm_listCommandsCmd.Flags().String("max-results", "", "(Optional) The maximum number of items to return for this call.")
		ssm_listCommandsCmd.Flags().String("next-token", "", "(Optional) The token for the next set of items to return.")
	})
	ssmCmd.AddCommand(ssm_listCommandsCmd)
}
