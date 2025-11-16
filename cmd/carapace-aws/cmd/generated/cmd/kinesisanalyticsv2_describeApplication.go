package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_describeApplicationCmd = &cobra.Command{
	Use:   "describe-application",
	Short: "Returns information about a specific Managed Service for Apache Flink application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_describeApplicationCmd).Standalone()

	kinesisanalyticsv2_describeApplicationCmd.Flags().String("application-name", "", "The name of the application.")
	kinesisanalyticsv2_describeApplicationCmd.Flags().String("include-additional-details", "", "Displays verbose information about a Managed Service for Apache Flink application, including the application's job plan.")
	kinesisanalyticsv2_describeApplicationCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_describeApplicationCmd)
}
