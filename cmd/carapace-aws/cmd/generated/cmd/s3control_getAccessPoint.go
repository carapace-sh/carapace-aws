package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3control_getAccessPointCmd = &cobra.Command{
	Use:   "get-access-point",
	Short: "Returns configuration information about the specified access point.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3control_getAccessPointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3control_getAccessPointCmd).Standalone()

		s3control_getAccessPointCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the account that owns the specified access point.")
		s3control_getAccessPointCmd.Flags().String("name", "", "The name of the access point whose configuration information you want to retrieve.")
		s3control_getAccessPointCmd.MarkFlagRequired("account-id")
		s3control_getAccessPointCmd.MarkFlagRequired("name")
	})
	s3controlCmd.AddCommand(s3control_getAccessPointCmd)
}
