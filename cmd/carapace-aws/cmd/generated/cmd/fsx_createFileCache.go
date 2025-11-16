package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_createFileCacheCmd = &cobra.Command{
	Use:   "create-file-cache",
	Short: "Creates a new Amazon File Cache resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_createFileCacheCmd).Standalone()

	fsx_createFileCacheCmd.Flags().String("client-request-token", "", "An idempotency token for resource creation, in a string of up to 63 ASCII characters.")
	fsx_createFileCacheCmd.Flags().String("copy-tags-to-data-repository-associations", "", "A boolean flag indicating whether tags for the cache should be copied to data repository associations.")
	fsx_createFileCacheCmd.Flags().String("data-repository-associations", "", "A list of up to 8 configurations for data repository associations (DRAs) to be created during the cache creation.")
	fsx_createFileCacheCmd.Flags().String("file-cache-type", "", "The type of cache that you're creating, which must be `LUSTRE`.")
	fsx_createFileCacheCmd.Flags().String("file-cache-type-version", "", "Sets the Lustre version for the cache that you're creating, which must be `2.12`.")
	fsx_createFileCacheCmd.Flags().String("kms-key-id", "", "Specifies the ID of the Key Management Service (KMS) key to use for encrypting data on an Amazon File Cache.")
	fsx_createFileCacheCmd.Flags().String("lustre-configuration", "", "The configuration for the Amazon File Cache resource being created.")
	fsx_createFileCacheCmd.Flags().String("security-group-ids", "", "A list of IDs specifying the security groups to apply to all network interfaces created for Amazon File Cache access.")
	fsx_createFileCacheCmd.Flags().String("storage-capacity", "", "The storage capacity of the cache in gibibytes (GiB).")
	fsx_createFileCacheCmd.Flags().String("subnet-ids", "", "")
	fsx_createFileCacheCmd.Flags().String("tags", "", "")
	fsx_createFileCacheCmd.MarkFlagRequired("file-cache-type")
	fsx_createFileCacheCmd.MarkFlagRequired("file-cache-type-version")
	fsx_createFileCacheCmd.MarkFlagRequired("storage-capacity")
	fsx_createFileCacheCmd.MarkFlagRequired("subnet-ids")
	fsxCmd.AddCommand(fsx_createFileCacheCmd)
}
