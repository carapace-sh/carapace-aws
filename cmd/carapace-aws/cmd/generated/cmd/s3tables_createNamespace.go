package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_createNamespaceCmd = &cobra.Command{
	Use:   "create-namespace",
	Short: "Creates a namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_createNamespaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3tables_createNamespaceCmd).Standalone()

		s3tables_createNamespaceCmd.Flags().String("namespace", "", "A name for the namespace.")
		s3tables_createNamespaceCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket to create the namespace in.")
		s3tables_createNamespaceCmd.MarkFlagRequired("namespace")
		s3tables_createNamespaceCmd.MarkFlagRequired("table-bucket-arn")
	})
	s3tablesCmd.AddCommand(s3tables_createNamespaceCmd)
}
