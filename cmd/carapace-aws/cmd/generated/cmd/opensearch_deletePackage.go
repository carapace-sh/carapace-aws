package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_deletePackageCmd = &cobra.Command{
	Use:   "delete-package",
	Short: "Deletes an Amazon OpenSearch Service package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_deletePackageCmd).Standalone()

	opensearch_deletePackageCmd.Flags().String("package-id", "", "The internal ID of the package you want to delete.")
	opensearch_deletePackageCmd.MarkFlagRequired("package-id")
	opensearchCmd.AddCommand(opensearch_deletePackageCmd)
}
