package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_createLakeFormationOptInCmd = &cobra.Command{
	Use:   "create-lake-formation-opt-in",
	Short: "Enforce Lake Formation permissions for the given databases, tables, and principals.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_createLakeFormationOptInCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_createLakeFormationOptInCmd).Standalone()

		lakeformation_createLakeFormationOptInCmd.Flags().String("condition", "", "")
		lakeformation_createLakeFormationOptInCmd.Flags().String("principal", "", "")
		lakeformation_createLakeFormationOptInCmd.Flags().String("resource", "", "")
		lakeformation_createLakeFormationOptInCmd.MarkFlagRequired("principal")
		lakeformation_createLakeFormationOptInCmd.MarkFlagRequired("resource")
	})
	lakeformationCmd.AddCommand(lakeformation_createLakeFormationOptInCmd)
}
