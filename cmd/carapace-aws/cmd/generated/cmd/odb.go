package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odbCmd = &cobra.Command{
	Use:   "odb",
	Short: "Oracle Database@Amazon Web Services is an offering that enables you to access Oracle Exadata infrastructure managed by Oracle Cloud Infrastructure (OCI) inside Amazon Web Services data centers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odbCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odbCmd).Standalone()

	})
	rootCmd.AddCommand(odbCmd)
}
