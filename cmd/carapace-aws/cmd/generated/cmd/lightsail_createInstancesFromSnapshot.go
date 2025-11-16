package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createInstancesFromSnapshotCmd = &cobra.Command{
	Use:   "create-instances-from-snapshot",
	Short: "Creates one or more new instances from a manual or automatic snapshot of an instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createInstancesFromSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_createInstancesFromSnapshotCmd).Standalone()

		lightsail_createInstancesFromSnapshotCmd.Flags().String("add-ons", "", "An array of objects representing the add-ons to enable for the new instance.")
		lightsail_createInstancesFromSnapshotCmd.Flags().String("attached-disk-mapping", "", "An object containing information about one or more disk mappings.")
		lightsail_createInstancesFromSnapshotCmd.Flags().String("availability-zone", "", "The Availability Zone where you want to create your instances.")
		lightsail_createInstancesFromSnapshotCmd.Flags().String("bundle-id", "", "The bundle of specification information for your virtual private server (or *instance*), including the pricing plan (`micro_x_x`).")
		lightsail_createInstancesFromSnapshotCmd.Flags().String("instance-names", "", "The names for your new instances.")
		lightsail_createInstancesFromSnapshotCmd.Flags().String("instance-snapshot-name", "", "The name of the instance snapshot on which you are basing your new instances.")
		lightsail_createInstancesFromSnapshotCmd.Flags().String("ip-address-type", "", "The IP address type for the instance.")
		lightsail_createInstancesFromSnapshotCmd.Flags().String("key-pair-name", "", "The name for your key pair.")
		lightsail_createInstancesFromSnapshotCmd.Flags().String("restore-date", "", "The date of the automatic snapshot to use for the new instance.")
		lightsail_createInstancesFromSnapshotCmd.Flags().String("source-instance-name", "", "The name of the source instance from which the source automatic snapshot was created.")
		lightsail_createInstancesFromSnapshotCmd.Flags().String("tags", "", "The tag keys and optional values to add to the resource during create.")
		lightsail_createInstancesFromSnapshotCmd.Flags().String("use-latest-restorable-auto-snapshot", "", "A Boolean value to indicate whether to use the latest available automatic snapshot.")
		lightsail_createInstancesFromSnapshotCmd.Flags().String("user-data", "", "You can create a launch script that configures a server with additional user data.")
		lightsail_createInstancesFromSnapshotCmd.MarkFlagRequired("availability-zone")
		lightsail_createInstancesFromSnapshotCmd.MarkFlagRequired("bundle-id")
		lightsail_createInstancesFromSnapshotCmd.MarkFlagRequired("instance-names")
	})
	lightsailCmd.AddCommand(lightsail_createInstancesFromSnapshotCmd)
}
