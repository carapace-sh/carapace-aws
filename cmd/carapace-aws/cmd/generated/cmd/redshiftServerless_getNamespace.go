package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getNamespaceCmd = &cobra.Command{
	Use:   "get-namespace",
	Short: "Returns information about a namespace in Amazon Redshift Serverless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getNamespaceCmd).Standalone()

	redshiftServerless_getNamespaceCmd.Flags().String("namespace-name", "", "The name of the namespace to retrieve information for.")
	redshiftServerless_getNamespaceCmd.MarkFlagRequired("namespace-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_getNamespaceCmd)
}
