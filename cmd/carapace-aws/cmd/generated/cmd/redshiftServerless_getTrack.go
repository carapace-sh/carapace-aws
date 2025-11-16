package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_getTrackCmd = &cobra.Command{
	Use:   "get-track",
	Short: "Get the Redshift Serverless version for a specified track.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_getTrackCmd).Standalone()

	redshiftServerless_getTrackCmd.Flags().String("track-name", "", "The name of the track of which its version is fetched.")
	redshiftServerless_getTrackCmd.MarkFlagRequired("track-name")
	redshiftServerlessCmd.AddCommand(redshiftServerless_getTrackCmd)
}
