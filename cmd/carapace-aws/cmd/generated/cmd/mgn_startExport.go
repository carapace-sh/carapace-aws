package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_startExportCmd = &cobra.Command{
	Use:   "start-export",
	Short: "Start export.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_startExportCmd).Standalone()

	mgn_startExportCmd.Flags().String("s3-bucket", "", "Start export request s3 bucket.")
	mgn_startExportCmd.Flags().String("s3-bucket-owner", "", "Start export request s3 bucket owner.")
	mgn_startExportCmd.Flags().String("s3-key", "", "Start export request s3key.")
	mgn_startExportCmd.MarkFlagRequired("s3-bucket")
	mgn_startExportCmd.MarkFlagRequired("s3-key")
	mgnCmd.AddCommand(mgn_startExportCmd)
}
