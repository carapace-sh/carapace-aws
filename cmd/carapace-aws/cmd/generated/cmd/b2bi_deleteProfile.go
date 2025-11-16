package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var b2bi_deleteProfileCmd = &cobra.Command{
	Use:   "delete-profile",
	Short: "Deletes the specified profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(b2bi_deleteProfileCmd).Standalone()

	b2bi_deleteProfileCmd.Flags().String("profile-id", "", "Specifies the unique, system-generated identifier for the profile.")
	b2bi_deleteProfileCmd.MarkFlagRequired("profile-id")
	b2biCmd.AddCommand(b2bi_deleteProfileCmd)
}
