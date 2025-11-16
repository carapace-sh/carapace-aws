package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_createStorediScsivolumeCmd = &cobra.Command{
	Use:   "create-storedi-scsivolume",
	Short: "Creates a volume on a specified gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_createStorediScsivolumeCmd).Standalone()

	storagegateway_createStorediScsivolumeCmd.Flags().String("disk-id", "", "The unique identifier for the gateway local disk that is configured as a stored volume.")
	storagegateway_createStorediScsivolumeCmd.Flags().String("gateway-arn", "", "")
	storagegateway_createStorediScsivolumeCmd.Flags().Bool("kmsencrypted", false, "Set to `true` to use Amazon S3 server-side encryption with your own KMS key, or `false` to use a key managed by Amazon S3.")
	storagegateway_createStorediScsivolumeCmd.Flags().String("kmskey", "", "The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for Amazon S3 server-side encryption.")
	storagegateway_createStorediScsivolumeCmd.Flags().String("network-interface-id", "", "The network interface of the gateway on which to expose the iSCSI target.")
	storagegateway_createStorediScsivolumeCmd.Flags().Bool("no-kmsencrypted", false, "Set to `true` to use Amazon S3 server-side encryption with your own KMS key, or `false` to use a key managed by Amazon S3.")
	storagegateway_createStorediScsivolumeCmd.Flags().String("preserve-existing-data", "", "Set to `true` if you want to preserve the data on the local disk.")
	storagegateway_createStorediScsivolumeCmd.Flags().String("snapshot-id", "", "The snapshot ID (e.g., \"snap-1122aabb\") of the snapshot to restore as the new stored volume.")
	storagegateway_createStorediScsivolumeCmd.Flags().String("tags", "", "A list of up to 50 tags that can be assigned to a stored volume.")
	storagegateway_createStorediScsivolumeCmd.Flags().String("target-name", "", "The name of the iSCSI target used by an initiator to connect to a volume and used as a suffix for the target ARN.")
	storagegateway_createStorediScsivolumeCmd.MarkFlagRequired("disk-id")
	storagegateway_createStorediScsivolumeCmd.MarkFlagRequired("gateway-arn")
	storagegateway_createStorediScsivolumeCmd.MarkFlagRequired("network-interface-id")
	storagegateway_createStorediScsivolumeCmd.Flag("no-kmsencrypted").Hidden = true
	storagegateway_createStorediScsivolumeCmd.MarkFlagRequired("preserve-existing-data")
	storagegateway_createStorediScsivolumeCmd.MarkFlagRequired("target-name")
	storagegatewayCmd.AddCommand(storagegateway_createStorediScsivolumeCmd)
}
