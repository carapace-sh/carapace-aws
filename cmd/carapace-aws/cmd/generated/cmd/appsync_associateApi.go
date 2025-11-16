package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appsync_associateApiCmd = &cobra.Command{
	Use:   "associate-api",
	Short: "Maps an endpoint to your custom domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appsync_associateApiCmd).Standalone()

	appsync_associateApiCmd.Flags().String("api-id", "", "The API ID.")
	appsync_associateApiCmd.Flags().String("domain-name", "", "The domain name.")
	appsync_associateApiCmd.MarkFlagRequired("api-id")
	appsync_associateApiCmd.MarkFlagRequired("domain-name")
	appsyncCmd.AddCommand(appsync_associateApiCmd)
}
