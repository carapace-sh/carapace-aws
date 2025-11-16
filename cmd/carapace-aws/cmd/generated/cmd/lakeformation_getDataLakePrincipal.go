package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_getDataLakePrincipalCmd = &cobra.Command{
	Use:   "get-data-lake-principal",
	Short: "Returns the identity of the invoking principal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_getDataLakePrincipalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_getDataLakePrincipalCmd).Standalone()

	})
	lakeformationCmd.AddCommand(lakeformation_getDataLakePrincipalCmd)
}
