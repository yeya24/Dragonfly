package dfgetDownloader

import (
	"fmt"
	"strings"

	daemon_config "github.com/dragonflyoss/Dragonfly/dfdaemon/config"
	"github.com/dragonflyoss/Dragonfly/dfget/config"
	"github.com/dragonflyoss/Dragonfly/dfget/core"
	"github.com/dragonflyoss/Dragonfly/dfget/core/api"
	"github.com/dragonflyoss/Dragonfly/dfget/core/regist"
	"github.com/dragonflyoss/Dragonfly/pkg/errortypes"
)

// DFGetter implements Downloader to download file by dragonfly.
type DFGetter struct {
	cfg daemon_config.DFGetConfig
	supernodeAPI api.SupernodeAPI
	register     regist.SupernodeRegister
}

// NewGetter returns a dfget downloader from the given config.
func NewGetter(cfg *config.Config) *DFGetter {
	supernodeAPI := api.NewSupernodeAPI()
	register := regist.NewSupernodeRegister(cfg, supernodeAPI)
	return &DFGetter{
		cfg:          cfg,
		supernodeAPI: supernodeAPI,
		register:     register,
	}
}

// Download is the method of DFGetter to download by dragonfly.
func (dfGetter *DFGetter) Download(url string, header map[string][]string, name string) (string, error) {
	var err error
	var result *regist.RegisterResult
	if err := core.Prepare(); err != nil {
		return "", errortypes.New(config.CodePrepareError, err.Error())
	}

	if result, err = core.RegisterToSuperNode(dfGetter.cfg, dfGetter.register); err != nil {
		return "", errortypes.New(config.CodeRegisterError, err.Error())
	}

	if err = core.DownloadFile(dfGetter.cfg, dfGetter.supernodeAPI, dfGetter.register, result); err != nil {
		return "", errortypes.New(config.CodeDownloadError, err.Error())
	}
	return "", nil
}

func (dfGetter *DFGetter) newDFConfig(url string, header map[string][]string) *config.Config {
	c := &config.Config{}
	var headers []string
	for key, value := range header {
		// discard HTTP host header for backing to source successfully
		if strings.EqualFold(key, "host") {
			continue
		}
		if len(value) > 0 {
			for _, v := range value {
				headers = append(headers, fmt.Sprintf("%s:%s",key, v))
			}
		} else {
			headers = append(headers, fmt.Sprintf("%s:%s", key, ""))
		}
	}
	return &config.Config{
		URL: url,
		Header: header,
	}
}
