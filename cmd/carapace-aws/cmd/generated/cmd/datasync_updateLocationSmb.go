package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_updateLocationSmbCmd = &cobra.Command{
	Use:   "update-location-smb",
	Short: "Modifies the following configuration parameters of the Server Message Block (SMB) transfer location that you're using with DataSync.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_updateLocationSmbCmd).Standalone()

	datasync_updateLocationSmbCmd.Flags().String("agent-arns", "", "Specifies the DataSync agent (or agents) that can connect to your SMB file server.")
	datasync_updateLocationSmbCmd.Flags().String("authentication-type", "", "Specifies the authentication protocol that DataSync uses to connect to your SMB file server.")
	datasync_updateLocationSmbCmd.Flags().String("dns-ip-addresses", "", "Specifies the IP addresses (IPv4 or IPv6) for the DNS servers that your SMB file server belongs to.")
	datasync_updateLocationSmbCmd.Flags().String("domain", "", "Specifies the Windows domain name that your SMB file server belongs to.")
	datasync_updateLocationSmbCmd.Flags().String("kerberos-keytab", "", "Specifies your Kerberos key table (keytab) file, which includes mappings between your Kerberos principal and encryption keys.")
	datasync_updateLocationSmbCmd.Flags().String("kerberos-krb5-conf", "", "Specifies a Kerberos configuration file (`krb5.conf`) that defines your Kerberos realm configuration.")
	datasync_updateLocationSmbCmd.Flags().String("kerberos-principal", "", "Specifies a Kerberos prinicpal, which is an identity in your Kerberos realm that has permission to access the files, folders, and file metadata in your SMB file server.")
	datasync_updateLocationSmbCmd.Flags().String("location-arn", "", "Specifies the ARN of the SMB location that you want to update.")
	datasync_updateLocationSmbCmd.Flags().String("mount-options", "", "")
	datasync_updateLocationSmbCmd.Flags().String("password", "", "Specifies the password of the user who can mount your SMB file server and has permission to access the files and folders involved in your transfer.")
	datasync_updateLocationSmbCmd.Flags().String("server-hostname", "", "Specifies the domain name or IP address (IPv4 or IPv6) of the SMB file server that your DataSync agent connects to.")
	datasync_updateLocationSmbCmd.Flags().String("subdirectory", "", "Specifies the name of the share exported by your SMB file server where DataSync will read or write data.")
	datasync_updateLocationSmbCmd.Flags().String("user", "", "Specifies the user name that can mount your SMB file server and has permission to access the files and folders involved in your transfer.")
	datasync_updateLocationSmbCmd.MarkFlagRequired("location-arn")
	datasyncCmd.AddCommand(datasync_updateLocationSmbCmd)
}
