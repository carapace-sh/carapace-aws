package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_describeApplicationOperationCmd = &cobra.Command{
	Use:   "describe-application-operation",
	Short: "Provides a detailed description of a specified application operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_describeApplicationOperationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_describeApplicationOperationCmd).Standalone()

		kinesisanalyticsv2_describeApplicationOperationCmd.Flags().String("application-name", "", "")
		kinesisanalyticsv2_describeApplicationOperationCmd.Flags().String("operation-id", "", "")
		kinesisanalyticsv2_describeApplicationOperationCmd.MarkFlagRequired("application-name")
		kinesisanalyticsv2_describeApplicationOperationCmd.MarkFlagRequired("operation-id")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_describeApplicationOperationCmd)
}
