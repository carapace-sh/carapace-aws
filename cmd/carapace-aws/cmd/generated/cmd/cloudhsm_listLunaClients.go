package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_listLunaClientsCmd = &cobra.Command{
	Use:   "list-luna-clients",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_listLunaClientsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudhsm_listLunaClientsCmd).Standalone()

		cloudhsm_listLunaClientsCmd.Flags().String("next-token", "", "The `NextToken` value from a previous call to `ListLunaClients`.")
	})
	cloudhsmCmd.AddCommand(cloudhsm_listLunaClientsCmd)
}
