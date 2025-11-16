package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_getAwsorganizationsAccessStatusCmd = &cobra.Command{
	Use:   "get-awsorganizations-access-status",
	Short: "Get the Access Status for Organizations portfolio share feature.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_getAwsorganizationsAccessStatusCmd).Standalone()

	servicecatalogCmd.AddCommand(servicecatalog_getAwsorganizationsAccessStatusCmd)
}
