package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_createCachediScsivolumeCmd = &cobra.Command{
	Use:   "create-cachedi-scsivolume",
	Short: "Creates a cached volume on a specified cached volume gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_createCachediScsivolumeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_createCachediScsivolumeCmd).Standalone()

		storagegateway_createCachediScsivolumeCmd.Flags().String("client-token", "", "A unique identifier that you use to retry a request.")
		storagegateway_createCachediScsivolumeCmd.Flags().String("gateway-arn", "", "")
		storagegateway_createCachediScsivolumeCmd.Flags().Bool("kmsencrypted", false, "Set to `true` to use Amazon S3 server-side encryption with your own KMS key, or `false` to use a key managed by Amazon S3.")
		storagegateway_createCachediScsivolumeCmd.Flags().String("kmskey", "", "The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for Amazon S3 server-side encryption.")
		storagegateway_createCachediScsivolumeCmd.Flags().String("network-interface-id", "", "The network interface of the gateway on which to expose the iSCSI target.")
		storagegateway_createCachediScsivolumeCmd.Flags().Bool("no-kmsencrypted", false, "Set to `true` to use Amazon S3 server-side encryption with your own KMS key, or `false` to use a key managed by Amazon S3.")
		storagegateway_createCachediScsivolumeCmd.Flags().String("snapshot-id", "", "The snapshot ID (e.g. \"snap-1122aabb\") of the snapshot to restore as the new cached volume.")
		storagegateway_createCachediScsivolumeCmd.Flags().String("source-volume-arn", "", "The ARN for an existing volume.")
		storagegateway_createCachediScsivolumeCmd.Flags().String("tags", "", "A list of up to 50 tags that you can assign to a cached volume.")
		storagegateway_createCachediScsivolumeCmd.Flags().String("target-name", "", "The name of the iSCSI target used by an initiator to connect to a volume and used as a suffix for the target ARN.")
		storagegateway_createCachediScsivolumeCmd.Flags().String("volume-size-in-bytes", "", "The size of the volume in bytes.")
		storagegateway_createCachediScsivolumeCmd.MarkFlagRequired("client-token")
		storagegateway_createCachediScsivolumeCmd.MarkFlagRequired("gateway-arn")
		storagegateway_createCachediScsivolumeCmd.MarkFlagRequired("network-interface-id")
		storagegateway_createCachediScsivolumeCmd.Flag("no-kmsencrypted").Hidden = true
		storagegateway_createCachediScsivolumeCmd.MarkFlagRequired("target-name")
		storagegateway_createCachediScsivolumeCmd.MarkFlagRequired("volume-size-in-bytes")
	})
	storagegatewayCmd.AddCommand(storagegateway_createCachediScsivolumeCmd)
}
