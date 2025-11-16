package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_createPersistentAppUiCmd = &cobra.Command{
	Use:   "create-persistent-app-ui",
	Short: "Creates a persistent application user interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_createPersistentAppUiCmd).Standalone()

	emr_createPersistentAppUiCmd.Flags().String("emrcontainers-config", "", "The EMR containers configuration.")
	emr_createPersistentAppUiCmd.Flags().String("profiler-type", "", "The profiler type for the persistent application user interface.")
	emr_createPersistentAppUiCmd.Flags().String("tags", "", "Tags for the persistent application user interface.")
	emr_createPersistentAppUiCmd.Flags().String("target-resource-arn", "", "The unique Amazon Resource Name (ARN) of the target resource.")
	emr_createPersistentAppUiCmd.Flags().String("xreferer", "", "The cross reference for the persistent application user interface.")
	emr_createPersistentAppUiCmd.MarkFlagRequired("target-resource-arn")
	emrCmd.AddCommand(emr_createPersistentAppUiCmd)
}
