package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeTapeArchivesCmd = &cobra.Command{
	Use:   "describe-tape-archives",
	Short: "Returns a description of specified virtual tapes in the virtual tape shelf (VTS).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeTapeArchivesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_describeTapeArchivesCmd).Standalone()

		storagegateway_describeTapeArchivesCmd.Flags().String("limit", "", "Specifies that the number of virtual tapes described be limited to the specified number.")
		storagegateway_describeTapeArchivesCmd.Flags().String("marker", "", "An opaque string that indicates the position at which to begin describing virtual tapes.")
		storagegateway_describeTapeArchivesCmd.Flags().String("tape-arns", "", "Specifies one or more unique Amazon Resource Names (ARNs) that represent the virtual tapes you want to describe.")
	})
	storagegatewayCmd.AddCommand(storagegateway_describeTapeArchivesCmd)
}
