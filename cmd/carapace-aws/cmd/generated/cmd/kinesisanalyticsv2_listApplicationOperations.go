package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_listApplicationOperationsCmd = &cobra.Command{
	Use:   "list-application-operations",
	Short: "Lists all the operations performed for the specified application such as UpdateApplication, StartApplication etc.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_listApplicationOperationsCmd).Standalone()

	kinesisanalyticsv2_listApplicationOperationsCmd.Flags().String("application-name", "", "")
	kinesisanalyticsv2_listApplicationOperationsCmd.Flags().String("limit", "", "")
	kinesisanalyticsv2_listApplicationOperationsCmd.Flags().String("next-token", "", "")
	kinesisanalyticsv2_listApplicationOperationsCmd.Flags().String("operation", "", "")
	kinesisanalyticsv2_listApplicationOperationsCmd.Flags().String("operation-status", "", "")
	kinesisanalyticsv2_listApplicationOperationsCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_listApplicationOperationsCmd)
}
