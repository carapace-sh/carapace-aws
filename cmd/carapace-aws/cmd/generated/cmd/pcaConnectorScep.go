package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcaConnectorScepCmd = &cobra.Command{
	Use:   "pca-connector-scep",
	Short: "Connector for SCEP creates a connector between Amazon Web Services Private CA and your SCEP-enabled clients and devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcaConnectorScepCmd).Standalone()

	rootCmd.AddCommand(pcaConnectorScepCmd)
}
