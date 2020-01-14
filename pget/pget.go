package pget

import (
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Pget structs
type Pget struct {
	Trace		bool
	Utils
	TargetDir	string
	Procs		int
	URL			string
	TargetURLs	[]string
	args		[]string
	timeout		int
	useragent	string
	referer		string
	quiet		bool
}

type ignore struct {
	err error
}

type cause interface {
	Cause() error
}

// New for pget package
func New() *Pget {
	return &Pget{
		Trace:   false,
		Utils:   &Data{},
		Procs:   runtime.NumCPU(), // default
		timeout: 10,
	}
}

// ErrTop get important message from wrapped error message
func (pget Pget) ErrTop(err error) error {
	for e := err; e != nil; {
		switch e.(type) {
		case ignore:
			return nil
		case cause:
			e = e.(cause).Cause()
		default:
			return e
		}
	}

	return nil
}

// Run execute methods in pget package
func (pget *Pget) Run(url, targetDir, output string, procs, timeout int, userAgent, referer string, quiet bool) error {
	if err := pget.Ready(url, targetDir, output, procs, timeout, userAgent, referer, quiet); err != nil {
		return pget.ErrTop(err)
	}

	if err := pget.Checking(); err != nil {
		return errors.Wrap(err, "failed to check header")
	}

	if err := pget.Download(); err != nil {
		return err
	}
	time.Sleep(3)
	if err := pget.Utils.BindwithFiles(pget.Procs, quiet); err != nil {
		return err
	}

	return nil
}

// Ready method define the variables required to Download.
func (pget *Pget) Ready(url, targetDir, output string, procs, timeout int, userAgent, referer string, quiet bool) error {
	if procs := os.Getenv("GOMAXPROCS"); procs == "" {
		runtime.GOMAXPROCS(pget.Procs)
	}

	if procs > 2 {
		pget.Procs = procs
	}

	if timeout > 0 {
		pget.timeout = timeout
	}

	if url != ""{
		pget.URL = url
	}else{
		return errors.New("URL MUST BE NOT BLANK")
	}

	if targetDir != "" {
		info, err := os.Stat(targetDir)
		if err != nil {
			if !os.IsNotExist(err) {
				return errors.Wrap(err, "target dir is invalid")
			}

			if err := os.MkdirAll(targetDir, 0755); err != nil {
				return errors.Wrapf(err, "failed to create diretory at %s", targetDir)
			}

		} else if !info.IsDir() {
			return errors.New("target dir is not a valid directory")
		}
		targetDir = strings.TrimSuffix(targetDir, "/")
	}
	pget.TargetDir = targetDir

	if output != "" {
		pget.Utils.SetFileName(output)
	}

	if userAgent != "" {
		pget.useragent = userAgent
	}

	if referer != "" {
		pget.referer = referer
	}

	pget.quiet = quiet

	return nil
}

func (pget Pget) makeIgnoreErr() ignore {
	return ignore{
		err: errors.New("this is ignore message"),
	}
}

// Error for options: version, usage
func (i ignore) Error() string {
	return i.err.Error()
}

func (i ignore) Cause() error {
	return i.err
}
