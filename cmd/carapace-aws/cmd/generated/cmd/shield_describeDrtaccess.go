package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_describeDrtaccessCmd = &cobra.Command{
	Use:   "describe-drtaccess",
	Short: "Returns the current role and list of Amazon S3 log buckets used by the Shield Response Team (SRT) to access your Amazon Web Services account while assisting with attack mitigation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_describeDrtaccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_describeDrtaccessCmd).Standalone()

	})
	shieldCmd.AddCommand(shield_describeDrtaccessCmd)
}
