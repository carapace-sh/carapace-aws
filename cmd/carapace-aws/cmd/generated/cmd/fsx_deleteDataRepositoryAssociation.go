package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_deleteDataRepositoryAssociationCmd = &cobra.Command{
	Use:   "delete-data-repository-association",
	Short: "Deletes a data repository association on an Amazon FSx for Lustre file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_deleteDataRepositoryAssociationCmd).Standalone()

	fsx_deleteDataRepositoryAssociationCmd.Flags().String("association-id", "", "The ID of the data repository association that you want to delete.")
	fsx_deleteDataRepositoryAssociationCmd.Flags().String("client-request-token", "", "")
	fsx_deleteDataRepositoryAssociationCmd.Flags().String("delete-data-in-file-system", "", "Set to `true` to delete the data in the file system that corresponds to the data repository association.")
	fsx_deleteDataRepositoryAssociationCmd.MarkFlagRequired("association-id")
	fsxCmd.AddCommand(fsx_deleteDataRepositoryAssociationCmd)
}
