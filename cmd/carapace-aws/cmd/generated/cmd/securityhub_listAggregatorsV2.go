package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listAggregatorsV2Cmd = &cobra.Command{
	Use:   "list-aggregators-v2",
	Short: "Retrieves a list of V2 aggregators.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listAggregatorsV2Cmd).Standalone()

	securityhub_listAggregatorsV2Cmd.Flags().String("max-results", "", "The maximum number of results to return.")
	securityhub_listAggregatorsV2Cmd.Flags().String("next-token", "", "The token required for pagination.")
	securityhubCmd.AddCommand(securityhub_listAggregatorsV2Cmd)
}
