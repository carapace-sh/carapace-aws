package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_disableAwsorganizationsAccessCmd = &cobra.Command{
	Use:   "disable-awsorganizations-access",
	Short: "Disable portfolio sharing through the Organizations service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_disableAwsorganizationsAccessCmd).Standalone()

	servicecatalogCmd.AddCommand(servicecatalog_disableAwsorganizationsAccessCmd)
}
