package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastoreData_putObjectCmd = &cobra.Command{
	Use:   "put-object",
	Short: "Uploads an object to the specified path.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastoreData_putObjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastoreData_putObjectCmd).Standalone()

		mediastoreData_putObjectCmd.Flags().String("body", "", "The bytes to be stored.")
		mediastoreData_putObjectCmd.Flags().String("cache-control", "", "An optional `CacheControl` header that allows the caller to control the object's cache behavior.")
		mediastoreData_putObjectCmd.Flags().String("content-type", "", "The content type of the object.")
		mediastoreData_putObjectCmd.Flags().String("path", "", "The path (including the file name) where the object is stored in the container.")
		mediastoreData_putObjectCmd.Flags().String("storage-class", "", "Indicates the storage class of a `Put` request.")
		mediastoreData_putObjectCmd.Flags().String("upload-availability", "", "Indicates the availability of an object while it is still uploading.")
		mediastoreData_putObjectCmd.MarkFlagRequired("body")
		mediastoreData_putObjectCmd.MarkFlagRequired("path")
	})
	mediastoreDataCmd.AddCommand(mediastoreData_putObjectCmd)
}
