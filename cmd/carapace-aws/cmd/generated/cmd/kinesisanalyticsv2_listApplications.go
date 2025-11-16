package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_listApplicationsCmd = &cobra.Command{
	Use:   "list-applications",
	Short: "Returns a list of Managed Service for Apache Flink applications in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_listApplicationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_listApplicationsCmd).Standalone()

		kinesisanalyticsv2_listApplicationsCmd.Flags().String("limit", "", "The maximum number of applications to list.")
		kinesisanalyticsv2_listApplicationsCmd.Flags().String("next-token", "", "If a previous command returned a pagination token, pass it into this value to retrieve the next set of results.")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_listApplicationsCmd)
}
