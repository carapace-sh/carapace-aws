package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3tables_getNamespaceCmd = &cobra.Command{
	Use:   "get-namespace",
	Short: "Gets details about a namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3tables_getNamespaceCmd).Standalone()

	s3tables_getNamespaceCmd.Flags().String("namespace", "", "The name of the namespace.")
	s3tables_getNamespaceCmd.Flags().String("table-bucket-arn", "", "The Amazon Resource Name (ARN) of the table bucket.")
	s3tables_getNamespaceCmd.MarkFlagRequired("namespace")
	s3tables_getNamespaceCmd.MarkFlagRequired("table-bucket-arn")
	s3tablesCmd.AddCommand(s3tables_getNamespaceCmd)
}
