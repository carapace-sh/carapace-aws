package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_deleteNamespaceCmd = &cobra.Command{
	Use:   "delete-namespace",
	Short: "Deletes a namespace from Amazon Redshift Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_deleteNamespaceCmd).Standalone()

	redshiftServerless_deleteNamespaceCmd.Flags().String("final-snapshot-name", "", "The name of the snapshot to be created before the namespace is deleted.")
	redshiftServerless_deleteNamespaceCmd.Flags().String("final-snapshot-retention-period", "", "How long to retain the final snapshot.")
	redshiftServerless_deleteNamespaceCmd.Flags().String("namespace-name", "", "The name of the namespace to delete.")
	redshiftServerless_deleteNamespaceCmd.MarkFlagRequired("namespace-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_deleteNamespaceCmd)
}
