package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var acm_putAccountConfigurationCmd = &cobra.Command{
	Use:   "put-account-configuration",
	Short: "Adds or modifies account-level configurations in ACM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(acm_putAccountConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(acm_putAccountConfigurationCmd).Standalone()

		acm_putAccountConfigurationCmd.Flags().String("expiry-events", "", "Specifies expiration events associated with an account.")
		acm_putAccountConfigurationCmd.Flags().String("idempotency-token", "", "Customer-chosen string used to distinguish between calls to `PutAccountConfiguration`.")
		acm_putAccountConfigurationCmd.MarkFlagRequired("idempotency-token")
	})
	acmCmd.AddCommand(acm_putAccountConfigurationCmd)
}
