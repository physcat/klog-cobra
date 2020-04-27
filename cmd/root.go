package cmd

import (
	"os"

	goflags "flag"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "klog-cobra",
	Short: "Example how of using klog with a cobra application",
	Long: `Example how of using klog with a cobra application
	
Klog uses golang flags that can sometimes confuse things when using
cobra. Cobra uses pflags, this shows one way of getting started 
using both klog and cobra.

Try these two commands:
go run . 
go run . -v 5
`,
	Run: func(cmd *cobra.Command, args []string) {
		klog.V(0).InfoS("Message printed at Info level 0")
		klog.V(2).InfoS("Message printed at Info level 2")
		klog.V(5).InfoS("Message printed at Info level 5")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		klog.ErrorS(err, "rootCmd.Execute()")
		os.Exit(1)
	}
}

func init() {
	klog.InitFlags(nil)
	rootCmd.Flags().AddGoFlagSet(goflags.CommandLine)

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.klog-cobra.yaml)")
}

func initConfig() {
	//map command line flags to viper variables.
	//viper will prefer flags from command line rather than file
	if err := viper.BindPFlags(rootCmd.Flags()); err != nil {
		klog.ErrorS(err, "Failed bind flags")
	}
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			klog.ErrorS(err, "")
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".klog-cobra")
	}

	viper.AutomaticEnv()

	verbosity := viper.GetString("v")
	klog.InfoS("Verbosity", "v", verbosity)

	if err := viper.ReadInConfig(); err == nil {
		klog.InfoS("Using config", "file", viper.ConfigFileUsed())
	}

	// Manually update v flag from viper config file
	if verbosity == "0" {
		if err := goflags.Set("v", viper.GetString("v")); err != nil {
			klog.Errorf("%+v", err)
			return
		}
	}
}
