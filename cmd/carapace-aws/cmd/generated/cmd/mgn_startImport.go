package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_startImportCmd = &cobra.Command{
	Use:   "start-import",
	Short: "Start import.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_startImportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_startImportCmd).Standalone()

		mgn_startImportCmd.Flags().String("client-token", "", "Start import request client token.")
		mgn_startImportCmd.Flags().String("s3-bucket-source", "", "Start import request s3 bucket source.")
		mgn_startImportCmd.MarkFlagRequired("s3-bucket-source")
	})
	mgnCmd.AddCommand(mgn_startImportCmd)
}
