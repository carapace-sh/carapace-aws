package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_deleteLakeFormationOptInCmd = &cobra.Command{
	Use:   "delete-lake-formation-opt-in",
	Short: "Remove the Lake Formation permissions enforcement of the given databases, tables, and principals.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_deleteLakeFormationOptInCmd).Standalone()

	lakeformation_deleteLakeFormationOptInCmd.Flags().String("condition", "", "")
	lakeformation_deleteLakeFormationOptInCmd.Flags().String("principal", "", "")
	lakeformation_deleteLakeFormationOptInCmd.Flags().String("resource", "", "")
	lakeformation_deleteLakeFormationOptInCmd.MarkFlagRequired("principal")
	lakeformation_deleteLakeFormationOptInCmd.MarkFlagRequired("resource")
	lakeformationCmd.AddCommand(lakeformation_deleteLakeFormationOptInCmd)
}
