package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getUsageLimitCmd = &cobra.Command{
	Use:   "get-usage-limit",
	Short: "Returns information about a usage limit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getUsageLimitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_getUsageLimitCmd).Standalone()

		redshiftServerless_getUsageLimitCmd.Flags().String("usage-limit-id", "", "The unique identifier of the usage limit to return information for.")
		redshiftServerless_getUsageLimitCmd.MarkFlagRequired("usage-limit-id")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_getUsageLimitCmd)
}
