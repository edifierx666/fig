package fviper

import (
  "fmt"
  "os"

  "github.com/fsnotify/fsnotify"
  "github.com/pkg/errors"
  "github.com/spf13/viper"
)

const ConfigName = "config"

type ViperConfiger struct {
  ConfigName   string
  ConfigType   string
  ConfigFile   string
  ConfigPaths  []string
  EnvPrefix    string
  Result       any
  ErrorHandler func(err error)
  ReadEnv      bool
  v            *viper.Viper
  maunalRead   bool
  pure         bool
  nowatch      bool
}
type ViperOption func(configer *ViperConfiger)

func WithConfigName(name string) ViperOption {
  return func(configer *ViperConfiger) {
    configer.ConfigName = name
  }
}
func WithConfigType(name string) ViperOption {
  return func(configer *ViperConfiger) {
    configer.ConfigType = name
  }
}
func WithConfigFile(name string) ViperOption {
  return func(configer *ViperConfiger) {
    configer.ConfigFile = name
  }
}
func WithConfigPaths(paths []string) ViperOption {
  return func(configer *ViperConfiger) {
    configer.ConfigPaths = paths
  }
}
func WithEnvPrefix(name string) ViperOption {
  return func(configer *ViperConfiger) {
    configer.EnvPrefix = name
  }
}

// 序列化对象
func WithResult(result any) ViperOption {
  return func(configer *ViperConfiger) {
    configer.Result = result
  }
}

// 手动调用readConfig
func WithMaunalRead(maunalRead bool) ViperOption {
  return func(configer *ViperConfiger) {
    configer.maunalRead = maunalRead
  }
}

// WithPureViper 只生产viper + 路径
func WithPureViper(pure bool) ViperOption {
  return func(configer *ViperConfiger) {
    configer.pure = pure
  }
}

// WithReadEnv 读取环境变量
func WithReadEnv(read bool) ViperOption {
  return func(configer *ViperConfiger) {
    configer.ReadEnv = read
  }
}

// WithNoWatch 不监听文件修改
func WithNoWatch(nowatch bool) ViperOption {
  return func(configer *ViperConfiger) {
    configer.nowatch = nowatch
  }
} // 出现错误回调函数
func WithErrorHandler(f func(err error)) ViperOption {
  return func(configer *ViperConfiger) {
    configer.ErrorHandler = f
  }
}

func (vc *ViperConfiger) GetViper() *viper.Viper {
  return vc.v
}

func (vc *ViperConfiger) GetResult() any {
  return vc.Result
}

func (vc *ViperConfiger) unmarshall() {
  if err := errors.WithStack(vc.GetViper().Unmarshal(&vc.Result)); err != nil {
    fmt.Println(err)
  }
}

func New(options ...ViperOption) *ViperConfiger {
  v := viper.New()
  var result map[string]any
  executable, _ := os.Executable()
  dir, _ := os.Getwd()
  vc := &ViperConfiger{
    ConfigName:  ConfigName,
    ConfigType:  "",
    ConfigFile:  "",
    ConfigPaths: []string{".", "../", "./config", executable, dir},
    EnvPrefix:   "",
    Result:      &result,
    ErrorHandler: func(err error) {
      if err := errors.WithStack(errors.WithMessage(err, "文件配置解析错误")); err != nil {
        fmt.Println(err)
      }
    },
    maunalRead: false,
    pure:       false,
    nowatch:    false,
    v:          v,
  }

  for _, option := range options {
    option(vc)
  }

  for _, path := range vc.ConfigPaths {
    v.AddConfigPath(path)
  }
  if vc.ReadEnv {
    v.AutomaticEnv()
  }
  if vc.ConfigType != "" {
    v.SetConfigType(vc.ConfigType)
  }

  if vc.ConfigFile != "" {
    v.SetConfigFile(vc.ConfigFile)
  }

  if vc.EnvPrefix != "" {
    v.SetEnvPrefix(vc.EnvPrefix)
    v.AutomaticEnv()
  }

  if vc.ConfigName != "" {
    v.SetConfigName(vc.ConfigName)
  }

  if vc.pure {
    return vc
  }

  if !vc.nowatch {
    v.WatchConfig()
    v.OnConfigChange(
      func(in fsnotify.Event) {
        vc.unmarshall()
      },
    )
  }

  if !vc.maunalRead {
    if err := v.ReadInConfig(); err != nil {
      vc.ErrorHandler(err)
    }
    vc.unmarshall()
  }

  return vc
}
