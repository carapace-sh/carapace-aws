package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_listApplicationVersionsCmd = &cobra.Command{
	Use:   "list-application-versions",
	Short: "Lists all the versions for the specified application, including versions that were rolled back.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_listApplicationVersionsCmd).Standalone()

	kinesisanalyticsv2_listApplicationVersionsCmd.Flags().String("application-name", "", "The name of the application for which you want to list all versions.")
	kinesisanalyticsv2_listApplicationVersionsCmd.Flags().String("limit", "", "The maximum number of versions to list in this invocation of the operation.")
	kinesisanalyticsv2_listApplicationVersionsCmd.Flags().String("next-token", "", "If a previous invocation of this operation returned a pagination token, pass it into this value to retrieve the next set of results.")
	kinesisanalyticsv2_listApplicationVersionsCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_listApplicationVersionsCmd)
}
