package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_putSourceServerActionCmd = &cobra.Command{
	Use:   "put-source-server-action",
	Short: "Put source server post migration custom action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_putSourceServerActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_putSourceServerActionCmd).Standalone()

		mgn_putSourceServerActionCmd.Flags().String("account-id", "", "Source server post migration custom account ID.")
		mgn_putSourceServerActionCmd.Flags().String("action-id", "", "Source server post migration custom action ID.")
		mgn_putSourceServerActionCmd.Flags().String("action-name", "", "Source server post migration custom action name.")
		mgn_putSourceServerActionCmd.Flags().Bool("active", false, "Source server post migration custom action active status.")
		mgn_putSourceServerActionCmd.Flags().String("category", "", "Source server post migration custom action category.")
		mgn_putSourceServerActionCmd.Flags().String("description", "", "Source server post migration custom action description.")
		mgn_putSourceServerActionCmd.Flags().String("document-identifier", "", "Source server post migration custom action document identifier.")
		mgn_putSourceServerActionCmd.Flags().String("document-version", "", "Source server post migration custom action document version.")
		mgn_putSourceServerActionCmd.Flags().String("external-parameters", "", "Source server post migration custom action external parameters.")
		mgn_putSourceServerActionCmd.Flags().Bool("must-succeed-for-cutover", false, "Source server post migration custom action must succeed for cutover.")
		mgn_putSourceServerActionCmd.Flags().Bool("no-active", false, "Source server post migration custom action active status.")
		mgn_putSourceServerActionCmd.Flags().Bool("no-must-succeed-for-cutover", false, "Source server post migration custom action must succeed for cutover.")
		mgn_putSourceServerActionCmd.Flags().String("order", "", "Source server post migration custom action order.")
		mgn_putSourceServerActionCmd.Flags().String("parameters", "", "Source server post migration custom action parameters.")
		mgn_putSourceServerActionCmd.Flags().String("source-server-id", "", "Source server ID.")
		mgn_putSourceServerActionCmd.Flags().String("timeout-seconds", "", "Source server post migration custom action timeout in seconds.")
		mgn_putSourceServerActionCmd.MarkFlagRequired("action-id")
		mgn_putSourceServerActionCmd.MarkFlagRequired("action-name")
		mgn_putSourceServerActionCmd.MarkFlagRequired("document-identifier")
		mgn_putSourceServerActionCmd.Flag("no-active").Hidden = true
		mgn_putSourceServerActionCmd.Flag("no-must-succeed-for-cutover").Hidden = true
		mgn_putSourceServerActionCmd.MarkFlagRequired("order")
		mgn_putSourceServerActionCmd.MarkFlagRequired("source-server-id")
	})
	mgnCmd.AddCommand(mgn_putSourceServerActionCmd)
}
