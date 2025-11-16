package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_createTapesCmd = &cobra.Command{
	Use:   "create-tapes",
	Short: "Creates one or more virtual tapes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_createTapesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_createTapesCmd).Standalone()

		storagegateway_createTapesCmd.Flags().String("client-token", "", "A unique identifier that you use to retry a request.")
		storagegateway_createTapesCmd.Flags().String("gateway-arn", "", "The unique Amazon Resource Name (ARN) that represents the gateway to associate the virtual tapes with.")
		storagegateway_createTapesCmd.Flags().Bool("kmsencrypted", false, "Set to `true` to use Amazon S3 server-side encryption with your own KMS key, or `false` to use a key managed by Amazon S3.")
		storagegateway_createTapesCmd.Flags().String("kmskey", "", "The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for Amazon S3 server-side encryption.")
		storagegateway_createTapesCmd.Flags().Bool("no-kmsencrypted", false, "Set to `true` to use Amazon S3 server-side encryption with your own KMS key, or `false` to use a key managed by Amazon S3.")
		storagegateway_createTapesCmd.Flags().String("num-tapes-to-create", "", "The number of virtual tapes that you want to create.")
		storagegateway_createTapesCmd.Flags().String("pool-id", "", "The ID of the pool that you want to add your tape to for archiving.")
		storagegateway_createTapesCmd.Flags().String("tags", "", "A list of up to 50 tags that can be assigned to a virtual tape.")
		storagegateway_createTapesCmd.Flags().String("tape-barcode-prefix", "", "A prefix that you append to the barcode of the virtual tape you are creating.")
		storagegateway_createTapesCmd.Flags().String("tape-size-in-bytes", "", "The size, in bytes, of the virtual tapes that you want to create.")
		storagegateway_createTapesCmd.Flags().String("worm", "", "Set to `TRUE` if the tape you are creating is to be configured as a write-once-read-many (WORM) tape.")
		storagegateway_createTapesCmd.MarkFlagRequired("client-token")
		storagegateway_createTapesCmd.MarkFlagRequired("gateway-arn")
		storagegateway_createTapesCmd.Flag("no-kmsencrypted").Hidden = true
		storagegateway_createTapesCmd.MarkFlagRequired("num-tapes-to-create")
		storagegateway_createTapesCmd.MarkFlagRequired("tape-barcode-prefix")
		storagegateway_createTapesCmd.MarkFlagRequired("tape-size-in-bytes")
	})
	storagegatewayCmd.AddCommand(storagegateway_createTapesCmd)
}
