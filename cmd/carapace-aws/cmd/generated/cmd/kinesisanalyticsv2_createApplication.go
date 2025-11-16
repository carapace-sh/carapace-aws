package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates a Managed Service for Apache Flink application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_createApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_createApplicationCmd).Standalone()

		kinesisanalyticsv2_createApplicationCmd.Flags().String("application-configuration", "", "Use this parameter to configure the application.")
		kinesisanalyticsv2_createApplicationCmd.Flags().String("application-description", "", "A summary description of the application.")
		kinesisanalyticsv2_createApplicationCmd.Flags().String("application-mode", "", "Use the `STREAMING` mode to create a Managed Service for Apache Flink application.")
		kinesisanalyticsv2_createApplicationCmd.Flags().String("application-name", "", "The name of your application (for example, `sample-app`).")
		kinesisanalyticsv2_createApplicationCmd.Flags().String("cloud-watch-logging-options", "", "Use this parameter to configure an Amazon CloudWatch log stream to monitor application configuration errors.")
		kinesisanalyticsv2_createApplicationCmd.Flags().String("runtime-environment", "", "The runtime environment for the application.")
		kinesisanalyticsv2_createApplicationCmd.Flags().String("service-execution-role", "", "The IAM role used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.")
		kinesisanalyticsv2_createApplicationCmd.Flags().String("tags", "", "A list of one or more tags to assign to the application.")
		kinesisanalyticsv2_createApplicationCmd.MarkFlagRequired("application-name")
		kinesisanalyticsv2_createApplicationCmd.MarkFlagRequired("runtime-environment")
		kinesisanalyticsv2_createApplicationCmd.MarkFlagRequired("service-execution-role")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_createApplicationCmd)
}
