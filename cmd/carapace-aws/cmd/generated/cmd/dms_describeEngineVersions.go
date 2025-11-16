package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeEngineVersionsCmd = &cobra.Command{
	Use:   "describe-engine-versions",
	Short: "Returns information about the replication instance versions used in the project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeEngineVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeEngineVersionsCmd).Standalone()

		dms_describeEngineVersionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		dms_describeEngineVersionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	dmsCmd.AddCommand(dms_describeEngineVersionsCmd)
}
