package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_describeAccountOverviewCmd = &cobra.Command{
	Use:   "describe-account-overview",
	Short: "For the time range passed in, returns the number of open reactive insight that were created, the number of open proactive insights that were created, and the Mean Time to Recover (MTTR) for all closed reactive insights.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_describeAccountOverviewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_describeAccountOverviewCmd).Standalone()

		devopsGuru_describeAccountOverviewCmd.Flags().String("from-time", "", "The start of the time range passed in.")
		devopsGuru_describeAccountOverviewCmd.Flags().String("to-time", "", "The end of the time range passed in.")
		devopsGuru_describeAccountOverviewCmd.MarkFlagRequired("from-time")
	})
	devopsGuruCmd.AddCommand(devopsGuru_describeAccountOverviewCmd)
}
