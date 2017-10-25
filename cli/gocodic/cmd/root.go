package cmd

import (
	"fmt"
	"io"
	"os"
	"runtime"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/gocli"
)

const (
	//Name is applicatin name
	Name = "gpgpdump"
	//Version is version for applicatin
	Version = "v0.2.0dev"
)

//ExitCode is OS exit code enumeration class
type ExitCode int

const (
	//Normal is OS exit code "normal"
	Normal ExitCode = iota
	//Abnormal is OS exit code "abnormal"
	Abnormal
)

//Int convert integer value
func (c ExitCode) Int() int {
	return int(c)
}

var (
	cfgFile   string
	reader    io.Reader //input reader (maybe os.Stdin)
	result    string
	resultErr string
)

// RootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocodic [flags] [command]",
	Short: "APIs for codic.jp",
	Long:  "APIs for codic.jp",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(cui *gocli.UI) (exit ExitCode) {
	defer func() {
		if r := recover(); r != nil {
			cui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, _, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				cui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ": line", line)
			}
			exit = 1
		}
	}()

	exit = Normal
	reader = cui.Reader()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		exit = Abnormal
		return
	}
	if len(resultErr) > 0 {
		cui.OutputErrln(resultErr)
		exit = Abnormal
		return
	}
	cui.Outputln(result)
	return
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gocodic.yaml)")
	rootCmd.PersistentFlags().BoolP("json", "j", false, "output by JSON format (raw data)")
	rootCmd.PersistentFlags().StringP("token", "t", "", "access token of codic.jp")
	viper.BindPFlag("json", rootCmd.PersistentFlags().Lookup("json"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gocodic" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gocodic")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
