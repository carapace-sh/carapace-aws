package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_createLocationSmbCmd = &cobra.Command{
	Use:   "create-location-smb",
	Short: "Creates a transfer *location* for a Server Message Block (SMB) file server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_createLocationSmbCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasync_createLocationSmbCmd).Standalone()

		datasync_createLocationSmbCmd.Flags().String("agent-arns", "", "Specifies the DataSync agent (or agents) that can connect to your SMB file server.")
		datasync_createLocationSmbCmd.Flags().String("authentication-type", "", "Specifies the authentication protocol that DataSync uses to connect to your SMB file server.")
		datasync_createLocationSmbCmd.Flags().String("dns-ip-addresses", "", "Specifies the IPv4 or IPv6 addresses for the DNS servers that your SMB file server belongs to.")
		datasync_createLocationSmbCmd.Flags().String("domain", "", "Specifies the Windows domain name that your SMB file server belongs to.")
		datasync_createLocationSmbCmd.Flags().String("kerberos-keytab", "", "Specifies your Kerberos key table (keytab) file, which includes mappings between your Kerberos principal and encryption keys.")
		datasync_createLocationSmbCmd.Flags().String("kerberos-krb5-conf", "", "Specifies a Kerberos configuration file (`krb5.conf`) that defines your Kerberos realm configuration.")
		datasync_createLocationSmbCmd.Flags().String("kerberos-principal", "", "Specifies a Kerberos principal, which is an identity in your Kerberos realm that has permission to access the files, folders, and file metadata in your SMB file server.")
		datasync_createLocationSmbCmd.Flags().String("mount-options", "", "Specifies the version of the SMB protocol that DataSync uses to access your SMB file server.")
		datasync_createLocationSmbCmd.Flags().String("password", "", "Specifies the password of the user who can mount your SMB file server and has permission to access the files and folders involved in your transfer.")
		datasync_createLocationSmbCmd.Flags().String("server-hostname", "", "Specifies the domain name or IP address (IPv4 or IPv6) of the SMB file server that your DataSync agent connects to.")
		datasync_createLocationSmbCmd.Flags().String("subdirectory", "", "Specifies the name of the share exported by your SMB file server where DataSync will read or write data.")
		datasync_createLocationSmbCmd.Flags().String("tags", "", "Specifies labels that help you categorize, filter, and search for your Amazon Web Services resources.")
		datasync_createLocationSmbCmd.Flags().String("user", "", "Specifies the user that can mount and access the files, folders, and file metadata in your SMB file server.")
		datasync_createLocationSmbCmd.MarkFlagRequired("agent-arns")
		datasync_createLocationSmbCmd.MarkFlagRequired("server-hostname")
		datasync_createLocationSmbCmd.MarkFlagRequired("subdirectory")
	})
	datasyncCmd.AddCommand(datasync_createLocationSmbCmd)
}
