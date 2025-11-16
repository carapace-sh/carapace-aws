package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_describeResourceCmd = &cobra.Command{
	Use:   "describe-resource",
	Short: "Retrieves the current data access role for the given resource registered in Lake Formation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_describeResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_describeResourceCmd).Standalone()

		lakeformation_describeResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
		lakeformation_describeResourceCmd.MarkFlagRequired("resource-arn")
	})
	lakeformationCmd.AddCommand(lakeformation_describeResourceCmd)
}
