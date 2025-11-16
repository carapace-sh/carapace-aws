package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeEngineVersionsCmd = &cobra.Command{
	Use:   "describe-engine-versions",
	Short: "Returns a list of the available Redis OSS engine versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeEngineVersionsCmd).Standalone()

	memorydb_describeEngineVersionsCmd.Flags().Bool("default-only", false, "If true, specifies that only the default version of the specified engine or engine and major version combination is to be returned.")
	memorydb_describeEngineVersionsCmd.Flags().String("engine", "", "The name of the engine for which to list available versions.")
	memorydb_describeEngineVersionsCmd.Flags().String("engine-version", "", "The Redis OSS engine version")
	memorydb_describeEngineVersionsCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
	memorydb_describeEngineVersionsCmd.Flags().String("next-token", "", "An optional argument to pass in case the total number of records exceeds the value of MaxResults.")
	memorydb_describeEngineVersionsCmd.Flags().Bool("no-default-only", false, "If true, specifies that only the default version of the specified engine or engine and major version combination is to be returned.")
	memorydb_describeEngineVersionsCmd.Flags().String("parameter-group-family", "", "The name of a specific parameter group family to return details for.")
	memorydb_describeEngineVersionsCmd.Flag("no-default-only").Hidden = true
	memorydbCmd.AddCommand(memorydb_describeEngineVersionsCmd)
}
