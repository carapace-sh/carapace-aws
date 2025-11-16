package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_getAccountConfigurationCmd = &cobra.Command{
	Use:   "get-account-configuration",
	Short: "Returns the account configuration options associated with an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_getAccountConfigurationCmd).Standalone()

	acmCmd.AddCommand(acm_getAccountConfigurationCmd)
}
