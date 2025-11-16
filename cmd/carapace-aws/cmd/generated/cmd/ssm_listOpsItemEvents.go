package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listOpsItemEventsCmd = &cobra.Command{
	Use:   "list-ops-item-events",
	Short: "Returns a list of all OpsItem events in the current Amazon Web Services Region and Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listOpsItemEventsCmd).Standalone()

	ssm_listOpsItemEventsCmd.Flags().String("filters", "", "One or more OpsItem filters.")
	ssm_listOpsItemEventsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_listOpsItemEventsCmd.Flags().String("next-token", "", "A token to start the list.")
	ssmCmd.AddCommand(ssm_listOpsItemEventsCmd)
}
