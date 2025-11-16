package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_cancelResizeCmd = &cobra.Command{
	Use:   "cancel-resize",
	Short: "Cancels a resize operation for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_cancelResizeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_cancelResizeCmd).Standalone()

		redshift_cancelResizeCmd.Flags().String("cluster-identifier", "", "The unique identifier for the cluster that you want to cancel a resize operation for.")
		redshift_cancelResizeCmd.MarkFlagRequired("cluster-identifier")
	})
	redshiftCmd.AddCommand(redshift_cancelResizeCmd)
}
