package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_createDataRepositoryAssociationCmd = &cobra.Command{
	Use:   "create-data-repository-association",
	Short: "Creates an Amazon FSx for Lustre data repository association (DRA).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_createDataRepositoryAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_createDataRepositoryAssociationCmd).Standalone()

		fsx_createDataRepositoryAssociationCmd.Flags().String("batch-import-meta-data-on-create", "", "Set to `true` to run an import data repository task to import metadata from the data repository to the file system after the data repository association is created.")
		fsx_createDataRepositoryAssociationCmd.Flags().String("client-request-token", "", "")
		fsx_createDataRepositoryAssociationCmd.Flags().String("data-repository-path", "", "The path to the Amazon S3 data repository that will be linked to the file system.")
		fsx_createDataRepositoryAssociationCmd.Flags().String("file-system-id", "", "")
		fsx_createDataRepositoryAssociationCmd.Flags().String("file-system-path", "", "A path on the file system that points to a high-level directory (such as `/ns1/`) or subdirectory (such as `/ns1/subdir/`) that will be mapped 1-1 with `DataRepositoryPath`.")
		fsx_createDataRepositoryAssociationCmd.Flags().String("imported-file-chunk-size", "", "For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.")
		fsx_createDataRepositoryAssociationCmd.Flags().String("s3", "", "The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association.")
		fsx_createDataRepositoryAssociationCmd.Flags().String("tags", "", "")
		fsx_createDataRepositoryAssociationCmd.MarkFlagRequired("data-repository-path")
		fsx_createDataRepositoryAssociationCmd.MarkFlagRequired("file-system-id")
	})
	fsxCmd.AddCommand(fsx_createDataRepositoryAssociationCmd)
}
