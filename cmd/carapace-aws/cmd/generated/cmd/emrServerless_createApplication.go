package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_createApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrServerless_createApplicationCmd).Standalone()

		emrServerless_createApplicationCmd.Flags().String("architecture", "", "The CPU architecture of an application.")
		emrServerless_createApplicationCmd.Flags().String("auto-start-configuration", "", "The configuration for an application to automatically start on job submission.")
		emrServerless_createApplicationCmd.Flags().String("auto-stop-configuration", "", "The configuration for an application to automatically stop after a certain amount of time being idle.")
		emrServerless_createApplicationCmd.Flags().String("client-token", "", "The client idempotency token of the application to create.")
		emrServerless_createApplicationCmd.Flags().String("identity-center-configuration", "", "The IAM Identity Center Configuration accepts the Identity Center instance parameter required to enable trusted identity propagation.")
		emrServerless_createApplicationCmd.Flags().String("image-configuration", "", "The image configuration for all worker types.")
		emrServerless_createApplicationCmd.Flags().String("initial-capacity", "", "The capacity to initialize when the application is created.")
		emrServerless_createApplicationCmd.Flags().String("interactive-configuration", "", "The interactive configuration object that enables the interactive use cases to use when running an application.")
		emrServerless_createApplicationCmd.Flags().String("maximum-capacity", "", "The maximum capacity to allocate when the application is created.")
		emrServerless_createApplicationCmd.Flags().String("monitoring-configuration", "", "The configuration setting for monitoring.")
		emrServerless_createApplicationCmd.Flags().String("name", "", "The name of the application.")
		emrServerless_createApplicationCmd.Flags().String("network-configuration", "", "The network configuration for customer VPC connectivity.")
		emrServerless_createApplicationCmd.Flags().String("release-label", "", "The Amazon EMR release associated with the application.")
		emrServerless_createApplicationCmd.Flags().String("runtime-configuration", "", "The [Configuration](https://docs.aws.amazon.com/emr-serverless/latest/APIReference/API_Configuration.html) specifications to use when creating an application.")
		emrServerless_createApplicationCmd.Flags().String("scheduler-configuration", "", "The scheduler configuration for batch and streaming jobs running on this application.")
		emrServerless_createApplicationCmd.Flags().String("tags", "", "The tags assigned to the application.")
		emrServerless_createApplicationCmd.Flags().String("type", "", "The type of application you want to start, such as Spark or Hive.")
		emrServerless_createApplicationCmd.Flags().String("worker-type-specifications", "", "The key-value pairs that specify worker type to `WorkerTypeSpecificationInput`.")
		emrServerless_createApplicationCmd.MarkFlagRequired("client-token")
		emrServerless_createApplicationCmd.MarkFlagRequired("release-label")
		emrServerless_createApplicationCmd.MarkFlagRequired("type")
	})
	emrServerlessCmd.AddCommand(emrServerless_createApplicationCmd)
}
