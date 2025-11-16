package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates an empty cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_createClusterCmd).Standalone()

	snowball_createClusterCmd.Flags().String("address-id", "", "The ID for the address that you want the cluster shipped to.")
	snowball_createClusterCmd.Flags().String("description", "", "An optional description of this specific cluster, for example `Environmental Data Cluster-01`.")
	snowball_createClusterCmd.Flags().Bool("force-create-jobs", false, "Force to create cluster when user attempts to overprovision or underprovision a cluster.")
	snowball_createClusterCmd.Flags().String("forwarding-address-id", "", "The forwarding address ID for a cluster.")
	snowball_createClusterCmd.Flags().String("initial-cluster-size", "", "If provided, each job will be automatically created and associated with the new cluster.")
	snowball_createClusterCmd.Flags().String("job-type", "", "The type of job for this cluster.")
	snowball_createClusterCmd.Flags().String("kms-key-arn", "", "The `KmsKeyARN` value that you want to associate with this cluster.")
	snowball_createClusterCmd.Flags().String("long-term-pricing-ids", "", "Lists long-term pricing id that will be used to associate with jobs automatically created for the new cluster.")
	snowball_createClusterCmd.Flags().Bool("no-force-create-jobs", false, "Force to create cluster when user attempts to overprovision or underprovision a cluster.")
	snowball_createClusterCmd.Flags().String("notification", "", "The Amazon Simple Notification Service (Amazon SNS) notification settings for this cluster.")
	snowball_createClusterCmd.Flags().String("on-device-service-configuration", "", "Specifies the service or services on the Snow Family device that your transferred data will be exported from or imported into.")
	snowball_createClusterCmd.Flags().String("remote-management", "", "Allows you to securely operate and manage Snow devices in a cluster remotely from outside of your internal network.")
	snowball_createClusterCmd.Flags().String("resources", "", "The resources associated with the cluster job.")
	snowball_createClusterCmd.Flags().String("role-arn", "", "The `RoleARN` that you want to associate with this cluster.")
	snowball_createClusterCmd.Flags().String("shipping-option", "", "The shipping speed for each node in this cluster.")
	snowball_createClusterCmd.Flags().String("snowball-capacity-preference", "", "If your job is being created in one of the US regions, you have the option of specifying what size Snow device you'd like for this job.")
	snowball_createClusterCmd.Flags().String("snowball-type", "", "The type of Snow Family devices to use for this cluster.")
	snowball_createClusterCmd.Flags().String("tax-documents", "", "The tax documents required in your Amazon Web Services Region.")
	snowball_createClusterCmd.MarkFlagRequired("address-id")
	snowball_createClusterCmd.MarkFlagRequired("job-type")
	snowball_createClusterCmd.Flag("no-force-create-jobs").Hidden = true
	snowball_createClusterCmd.MarkFlagRequired("shipping-option")
	snowball_createClusterCmd.MarkFlagRequired("snowball-type")
	snowballCmd.AddCommand(snowball_createClusterCmd)
}
