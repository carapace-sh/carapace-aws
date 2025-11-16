package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_getUpgradeStatusCmd = &cobra.Command{
	Use:   "get-upgrade-status",
	Short: "Retrieves the latest status of the last upgrade or upgrade eligibility check that was performed on the domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_getUpgradeStatusCmd).Standalone()

	es_getUpgradeStatusCmd.Flags().String("domain-name", "", "")
	es_getUpgradeStatusCmd.MarkFlagRequired("domain-name")
	esCmd.AddCommand(es_getUpgradeStatusCmd)
}
