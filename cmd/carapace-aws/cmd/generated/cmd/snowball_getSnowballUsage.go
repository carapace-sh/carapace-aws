package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_getSnowballUsageCmd = &cobra.Command{
	Use:   "get-snowball-usage",
	Short: "Returns information about the Snow Family service limit for your account, and also the number of Snow devices your account has in use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_getSnowballUsageCmd).Standalone()

	snowballCmd.AddCommand(snowball_getSnowballUsageCmd)
}
