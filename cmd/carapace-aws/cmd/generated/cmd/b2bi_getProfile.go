package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_getProfileCmd = &cobra.Command{
	Use:   "get-profile",
	Short: "Retrieves the details for the profile specified by the profile ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_getProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(b2bi_getProfileCmd).Standalone()

		b2bi_getProfileCmd.Flags().String("profile-id", "", "Specifies the unique, system-generated identifier for the profile.")
		b2bi_getProfileCmd.MarkFlagRequired("profile-id")
	})
	b2biCmd.AddCommand(b2bi_getProfileCmd)
}
