package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_updateDataRepositoryAssociationCmd = &cobra.Command{
	Use:   "update-data-repository-association",
	Short: "Updates the configuration of an existing data repository association on an Amazon FSx for Lustre file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_updateDataRepositoryAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_updateDataRepositoryAssociationCmd).Standalone()

		fsx_updateDataRepositoryAssociationCmd.Flags().String("association-id", "", "The ID of the data repository association that you are updating.")
		fsx_updateDataRepositoryAssociationCmd.Flags().String("client-request-token", "", "")
		fsx_updateDataRepositoryAssociationCmd.Flags().String("imported-file-chunk-size", "", "For files imported from a data repository, this value determines the stripe count and maximum amount of data per file (in MiB) stored on a single physical disk.")
		fsx_updateDataRepositoryAssociationCmd.Flags().String("s3", "", "The configuration for an Amazon S3 data repository linked to an Amazon FSx Lustre file system with a data repository association.")
		fsx_updateDataRepositoryAssociationCmd.MarkFlagRequired("association-id")
	})
	fsxCmd.AddCommand(fsx_updateDataRepositoryAssociationCmd)
}
