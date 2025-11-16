package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearchdomainCmd = &cobra.Command{
	Use:   "cloudsearchdomain",
	Short: "You use the AmazonCloudSearch2013 API to upload documents to a search domain and search those documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearchdomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearchdomainCmd).Standalone()

	})
	rootCmd.AddCommand(cloudsearchdomainCmd)
}
