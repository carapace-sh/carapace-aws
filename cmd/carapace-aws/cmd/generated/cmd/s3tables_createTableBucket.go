package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_createTableBucketCmd = &cobra.Command{
	Use:   "create-table-bucket",
	Short: "Creates a table bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_createTableBucketCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_createTableBucketCmd).Standalone()

		s3tables_createTableBucketCmd.Flags().String("encryption-configuration", "", "The encryption configuration to use for the table bucket.")
		s3tables_createTableBucketCmd.Flags().String("name", "", "The name for the table bucket.")
		s3tables_createTableBucketCmd.Flags().String("tags", "", "A map of user-defined tags that you would like to apply to the table bucket that you are creating.")
		s3tables_createTableBucketCmd.MarkFlagRequired("name")
	})
	s3tablesCmd.AddCommand(s3tables_createTableBucketCmd)
}
