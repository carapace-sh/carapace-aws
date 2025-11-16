package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_describeOptionGroupOptionsCmd = &cobra.Command{
	Use:   "describe-option-group-options",
	Short: "Describes all available options for the specified engine.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_describeOptionGroupOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_describeOptionGroupOptionsCmd).Standalone()

		rds_describeOptionGroupOptionsCmd.Flags().String("engine-name", "", "The name of the engine to describe options for.")
		rds_describeOptionGroupOptionsCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
		rds_describeOptionGroupOptionsCmd.Flags().String("major-engine-version", "", "If specified, filters the results to include only options for the specified major engine version.")
		rds_describeOptionGroupOptionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		rds_describeOptionGroupOptionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		rds_describeOptionGroupOptionsCmd.MarkFlagRequired("engine-name")
	})
	rdsCmd.AddCommand(rds_describeOptionGroupOptionsCmd)
}
