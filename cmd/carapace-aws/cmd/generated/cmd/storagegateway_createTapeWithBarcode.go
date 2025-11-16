package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_createTapeWithBarcodeCmd = &cobra.Command{
	Use:   "create-tape-with-barcode",
	Short: "Creates a virtual tape by using your own barcode.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_createTapeWithBarcodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(storagegateway_createTapeWithBarcodeCmd).Standalone()

		storagegateway_createTapeWithBarcodeCmd.Flags().String("gateway-arn", "", "The unique Amazon Resource Name (ARN) that represents the gateway to associate the virtual tape with.")
		storagegateway_createTapeWithBarcodeCmd.Flags().Bool("kmsencrypted", false, "Set to `true` to use Amazon S3 server-side encryption with your own KMS key, or `false` to use a key managed by Amazon S3.")
		storagegateway_createTapeWithBarcodeCmd.Flags().String("kmskey", "", "The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for Amazon S3 server-side encryption.")
		storagegateway_createTapeWithBarcodeCmd.Flags().Bool("no-kmsencrypted", false, "Set to `true` to use Amazon S3 server-side encryption with your own KMS key, or `false` to use a key managed by Amazon S3.")
		storagegateway_createTapeWithBarcodeCmd.Flags().String("pool-id", "", "The ID of the pool that you want to add your tape to for archiving.")
		storagegateway_createTapeWithBarcodeCmd.Flags().String("tags", "", "A list of up to 50 tags that can be assigned to a virtual tape that has a barcode.")
		storagegateway_createTapeWithBarcodeCmd.Flags().String("tape-barcode", "", "The barcode that you want to assign to the tape.")
		storagegateway_createTapeWithBarcodeCmd.Flags().String("tape-size-in-bytes", "", "The size, in bytes, of the virtual tape that you want to create.")
		storagegateway_createTapeWithBarcodeCmd.Flags().String("worm", "", "Set to `TRUE` if the tape you are creating is to be configured as a write-once-read-many (WORM) tape.")
		storagegateway_createTapeWithBarcodeCmd.MarkFlagRequired("gateway-arn")
		storagegateway_createTapeWithBarcodeCmd.Flag("no-kmsencrypted").Hidden = true
		storagegateway_createTapeWithBarcodeCmd.MarkFlagRequired("tape-barcode")
		storagegateway_createTapeWithBarcodeCmd.MarkFlagRequired("tape-size-in-bytes")
	})
	storagegatewayCmd.AddCommand(storagegateway_createTapeWithBarcodeCmd)
}
