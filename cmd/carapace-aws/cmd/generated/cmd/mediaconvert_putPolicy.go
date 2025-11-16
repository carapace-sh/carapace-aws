package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_putPolicyCmd = &cobra.Command{
	Use:   "put-policy",
	Short: "Create or change your policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_putPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_putPolicyCmd).Standalone()

		mediaconvert_putPolicyCmd.Flags().String("policy", "", "A policy configures behavior that you allow or disallow for your account.")
		mediaconvert_putPolicyCmd.MarkFlagRequired("policy")
	})
	mediaconvertCmd.AddCommand(mediaconvert_putPolicyCmd)
}
