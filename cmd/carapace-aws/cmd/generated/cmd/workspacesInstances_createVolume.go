package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesInstances_createVolumeCmd = &cobra.Command{
	Use:   "create-volume",
	Short: "Creates a new volume for WorkSpace Instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesInstances_createVolumeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesInstances_createVolumeCmd).Standalone()

		workspacesInstances_createVolumeCmd.Flags().String("availability-zone", "", "Availability zone for the volume.")
		workspacesInstances_createVolumeCmd.Flags().String("client-token", "", "Unique token to prevent duplicate volume creation.")
		workspacesInstances_createVolumeCmd.Flags().Bool("encrypted", false, "Indicates if the volume should be encrypted.")
		workspacesInstances_createVolumeCmd.Flags().String("iops", "", "Input/output operations per second for the volume.")
		workspacesInstances_createVolumeCmd.Flags().String("kms-key-id", "", "KMS key for volume encryption.")
		workspacesInstances_createVolumeCmd.Flags().Bool("no-encrypted", false, "Indicates if the volume should be encrypted.")
		workspacesInstances_createVolumeCmd.Flags().String("size-in-gb", "", "Volume size in gigabytes.")
		workspacesInstances_createVolumeCmd.Flags().String("snapshot-id", "", "Source snapshot for volume creation.")
		workspacesInstances_createVolumeCmd.Flags().String("tag-specifications", "", "Metadata tags for the volume.")
		workspacesInstances_createVolumeCmd.Flags().String("throughput", "", "Volume throughput performance.")
		workspacesInstances_createVolumeCmd.Flags().String("volume-type", "", "Type of EBS volume.")
		workspacesInstances_createVolumeCmd.MarkFlagRequired("availability-zone")
		workspacesInstances_createVolumeCmd.Flag("no-encrypted").Hidden = true
	})
	workspacesInstancesCmd.AddCommand(workspacesInstances_createVolumeCmd)
}
