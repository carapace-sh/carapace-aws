package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_enableAwsorganizationsAccessCmd = &cobra.Command{
	Use:   "enable-awsorganizations-access",
	Short: "Enable portfolio sharing feature through Organizations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_enableAwsorganizationsAccessCmd).Standalone()

	servicecatalogCmd.AddCommand(servicecatalog_enableAwsorganizationsAccessCmd)
}
