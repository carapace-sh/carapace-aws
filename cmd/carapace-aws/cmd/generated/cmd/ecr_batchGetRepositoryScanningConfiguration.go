package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_batchGetRepositoryScanningConfigurationCmd = &cobra.Command{
	Use:   "batch-get-repository-scanning-configuration",
	Short: "Gets the scanning configuration for one or more repositories.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_batchGetRepositoryScanningConfigurationCmd).Standalone()

	ecr_batchGetRepositoryScanningConfigurationCmd.Flags().String("repository-names", "", "One or more repository names to get the scanning configuration for.")
	ecr_batchGetRepositoryScanningConfigurationCmd.MarkFlagRequired("repository-names")
	ecrCmd.AddCommand(ecr_batchGetRepositoryScanningConfigurationCmd)
}
