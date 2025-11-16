package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrServerless_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates a specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrServerless_updateApplicationCmd).Standalone()

	emrServerless_updateApplicationCmd.Flags().String("application-id", "", "The ID of the application to update.")
	emrServerless_updateApplicationCmd.Flags().String("architecture", "", "The CPU architecture of an application.")
	emrServerless_updateApplicationCmd.Flags().String("auto-start-configuration", "", "The configuration for an application to automatically start on job submission.")
	emrServerless_updateApplicationCmd.Flags().String("auto-stop-configuration", "", "The configuration for an application to automatically stop after a certain amount of time being idle.")
	emrServerless_updateApplicationCmd.Flags().String("client-token", "", "The client idempotency token of the application to update.")
	emrServerless_updateApplicationCmd.Flags().String("identity-center-configuration", "", "Specifies the IAM Identity Center configuration used to enable or disable trusted identity propagation.")
	emrServerless_updateApplicationCmd.Flags().String("image-configuration", "", "The image configuration to be used for all worker types.")
	emrServerless_updateApplicationCmd.Flags().String("initial-capacity", "", "The capacity to initialize when the application is updated.")
	emrServerless_updateApplicationCmd.Flags().String("interactive-configuration", "", "The interactive configuration object that contains new interactive use cases when the application is updated.")
	emrServerless_updateApplicationCmd.Flags().String("maximum-capacity", "", "The maximum capacity to allocate when the application is updated.")
	emrServerless_updateApplicationCmd.Flags().String("monitoring-configuration", "", "The configuration setting for monitoring.")
	emrServerless_updateApplicationCmd.Flags().String("network-configuration", "", "")
	emrServerless_updateApplicationCmd.Flags().String("release-label", "", "The Amazon EMR release label for the application.")
	emrServerless_updateApplicationCmd.Flags().String("runtime-configuration", "", "The [Configuration](https://docs.aws.amazon.com/emr-serverless/latest/APIReference/API_Configuration.html) specifications to use when updating an application.")
	emrServerless_updateApplicationCmd.Flags().String("scheduler-configuration", "", "The scheduler configuration for batch and streaming jobs running on this application.")
	emrServerless_updateApplicationCmd.Flags().String("worker-type-specifications", "", "The key-value pairs that specify worker type to `WorkerTypeSpecificationInput`.")
	emrServerless_updateApplicationCmd.MarkFlagRequired("application-id")
	emrServerless_updateApplicationCmd.MarkFlagRequired("client-token")
	emrServerlessCmd.AddCommand(emrServerless_updateApplicationCmd)
}
