package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "Creates a job to import or export data between Amazon S3 and your on-premises data center.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_createJobCmd).Standalone()

	snowball_createJobCmd.Flags().String("address-id", "", "The ID for the address that you want the Snow device shipped to.")
	snowball_createJobCmd.Flags().String("cluster-id", "", "The ID of a cluster.")
	snowball_createJobCmd.Flags().String("description", "", "Defines an optional description of this specific job, for example `Important Photos 2016-08-11`.")
	snowball_createJobCmd.Flags().String("device-configuration", "", "Defines the device configuration for an Snowball Edge job.")
	snowball_createJobCmd.Flags().String("forwarding-address-id", "", "The forwarding address ID for a job.")
	snowball_createJobCmd.Flags().String("impact-level", "", "The highest impact level of data that will be stored or processed on the device, provided at job creation.")
	snowball_createJobCmd.Flags().String("job-type", "", "Defines the type of job that you're creating.")
	snowball_createJobCmd.Flags().String("kms-key-arn", "", "The `KmsKeyARN` that you want to associate with this job.")
	snowball_createJobCmd.Flags().String("long-term-pricing-id", "", "The ID of the long-term pricing type for the device.")
	snowball_createJobCmd.Flags().String("notification", "", "Defines the Amazon Simple Notification Service (Amazon SNS) notification settings for this job.")
	snowball_createJobCmd.Flags().String("on-device-service-configuration", "", "Specifies the service or services on the Snow Family device that your transferred data will be exported from or imported into.")
	snowball_createJobCmd.Flags().String("pickup-details", "", "Information identifying the person picking up the device.")
	snowball_createJobCmd.Flags().String("remote-management", "", "Allows you to securely operate and manage Snowcone devices remotely from outside of your internal network.")
	snowball_createJobCmd.Flags().String("resources", "", "Defines the Amazon S3 buckets associated with this job.")
	snowball_createJobCmd.Flags().String("role-arn", "", "The `RoleARN` that you want to associate with this job.")
	snowball_createJobCmd.Flags().String("shipping-option", "", "The shipping speed for this job.")
	snowball_createJobCmd.Flags().String("snowball-capacity-preference", "", "If your job is being created in one of the US regions, you have the option of specifying what size Snow device you'd like for this job.")
	snowball_createJobCmd.Flags().String("snowball-type", "", "The type of Snow Family devices to use for this job.")
	snowball_createJobCmd.Flags().String("tax-documents", "", "The tax documents required in your Amazon Web Services Region.")
	snowballCmd.AddCommand(snowball_createJobCmd)
}
