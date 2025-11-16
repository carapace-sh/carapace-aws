package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissionsCmd = &cobra.Command{
	Use:   "verifiedpermissions",
	Short: "Amazon Verified Permissions is a permissions management service from Amazon Web Services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissionsCmd).Standalone()

	rootCmd.AddCommand(verifiedpermissionsCmd)
}
