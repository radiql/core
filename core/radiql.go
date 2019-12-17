package core

import (
	"github.com/radiql/core/fs"
)

// Runtime provides access to core RADiQL services
type Runtime struct {
	fs          fs.Fs
	fsProviders []fs.Provider
}

var rt = Runtime{}

// RADiQL returns an instance of the RADiQL runtime.
func RADiQL() Runtime {
	return rt
}

// RegisterFsProvider registers a file system provider. This is intended to be
// used by RADiQL platform implementations such as AppStratum. However, it can also
// be used to register a custom file system for testing or production.
func (r *Runtime) RegisterFsProvider(provider fs.Provider) {
	for _, p := range r.fsProviders {
		if p == provider {
			return
		}
	}
	r.fsProviders = append(r.fsProviders, provider)
}

// SelectFs selects the file system
func (r *Runtime) SelectFs(fsType string) (fs.Fs, error) {
	// If existing file system then clean up

	// Then switch to newly selected file system.
	for _, p := range r.fsProviders {
		if p.Supports(fsType) {
			r.fs = p.Get(fsType)
			return r.fs, nil
		}
	}
	return nil, fs.ErrUnknownFileSystem
}

// Fs returns the current file system and mounts a default file system implementation if
// none has been set explicitly. By default this is the file system provided by the os package.
func (r *Runtime) Fs() (fs.Fs, error) {
	if r.fs == nil {
		return r.SelectFs(fs.OsFileSystem)
	}
	return r.fs, nil
}
